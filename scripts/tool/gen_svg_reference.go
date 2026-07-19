package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"hash/crc32"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

const (
	resourceRoot     = "etc/resources/aws/svg"
	serviceCatalog   = "etc/resources/aws/service-catalog.csv"
	outputDir        = "docs/src/reference/icons"
	sampleDir        = "docs/src/reference/samples/icons"
	previewDir       = "docs/src/reference/previews/icons"
	referenceRoot    = "docs/src/reference"
	groupIconRelPath = "etc/resources/aws/svg/Architecture-Group-Icons/"

	sampleMargin      = 24
	groupSampleHeight = 112
	groupSampleMinW   = 156
	iconPageSize      = 90
)

type svgAsset struct {
	Path     string
	Name     string
	Group    string
	Slug     string
	Usage    string
	Sample   string
	SizeByte int
}

type catalogEntry struct {
	ID       string
	Category string
	Service  string
}

func main() {
	catalog, err := loadCatalog(serviceCatalog)
	if err != nil {
		fatal(err)
	}
	assets, err := collectSVGAssets(resourceRoot, catalog)
	if err != nil {
		fatal(err)
	}
	prepareGeneratedDirs()
	renderer := newUsecase()
	groups := groupAssets(assets)
	writeIndex(groups, len(assets))
	for _, group := range groups {
		writeGroupSamples(renderer, group)
		writeGroupPage(group)
	}
	writeReferenceSamples(renderer)
}

type assetGroup struct {
	Name   string
	Slug   string
	Assets []svgAsset
}

type assetPage struct {
	Number int
	File   string
	Start  int
	End    int
	Assets []svgAsset
}

func groupAssets(assets []svgAsset) []assetGroup {
	groupByName := make(map[string][]svgAsset)
	for _, asset := range assets {
		groupByName[asset.Group] = append(groupByName[asset.Group], asset)
	}
	names := make([]string, 0, len(groupByName))
	for name := range groupByName {
		names = append(names, name)
	}
	sort.Strings(names)
	groups := make([]assetGroup, 0, len(names))
	for _, name := range names {
		groups = append(groups, assetGroup{
			Name:   name,
			Slug:   slug(name),
			Assets: groupByName[name],
		})
	}
	return groups
}

func writeIndex(groups []assetGroup, total int) {
	var out bytes.Buffer
	fmt.Fprintf(&out, `# SVG Icon Reference

This reference uses the SVG files published from `+"`etc/resources/aws/svg`"+`.
Large sections are split into paginated pages to keep local preview responsive.
Total SVG files: %d.

`, total)
	for _, group := range groups {
		pages := groupPages(group)
		if len(pages) > 1 {
			fmt.Fprintf(&out, "- [%s](%s.md) (%d, %d pages)\n", group.Name, group.Slug, len(group.Assets), len(pages))
			continue
		}
		fmt.Fprintf(&out, "- [%s](%s.md) (%d)\n", group.Name, group.Slug, len(group.Assets))
	}
	writeFile(filepath.Join(outputDir, "index.md"), out.Bytes())
}

func writeGroupPage(group assetGroup) {
	pages := groupPages(group)
	if len(pages) > 1 {
		writeGroupIndexPage(group, pages)
		for _, page := range pages {
			writeGroupAssetPage(group, page)
		}
		return
	}
	writeGroupAssetPage(group, pages[0])
}

func groupPages(group assetGroup) []assetPage {
	if len(group.Assets) == 0 {
		return []assetPage{{Number: 1, File: group.Slug + ".md"}}
	}
	count := (len(group.Assets) + iconPageSize - 1) / iconPageSize
	pages := make([]assetPage, 0, count)
	for i := 0; i < count; i++ {
		start := i * iconPageSize
		end := start + iconPageSize
		if end > len(group.Assets) {
			end = len(group.Assets)
		}
		file := group.Slug + ".md"
		if count > 1 {
			file = fmt.Sprintf("%s-%03d.md", group.Slug, i+1)
		}
		pages = append(pages, assetPage{
			Number: i + 1,
			File:   file,
			Start:  start,
			End:    end,
			Assets: group.Assets[start:end],
		})
	}
	return pages
}

func writeGroupIndexPage(group assetGroup, pages []assetPage) {
	var out bytes.Buffer
	fmt.Fprintf(&out, `# %s

This section contains %d SVG files from `+"`etc/resources/aws/svg/%s`"+`.
Open a page below to load a smaller set of previews.

`, group.Name, len(group.Assets), group.Name)
	for _, page := range pages {
		fmt.Fprintf(&out, "- [Page %d](%s) (%d-%d)\n", page.Number, page.File, page.Start+1, page.End)
	}
	writeFile(filepath.Join(outputDir, group.Slug+".md"), out.Bytes())
}

func writeGroupAssetPage(group assetGroup, page assetPage) {
	var out bytes.Buffer
	writeHeader(&out, group, page)
	currentGroup := ""
	for i, asset := range page.Assets {
		if asset.Group != currentGroup {
			if currentGroup != "" {
				out.WriteString("</div>\n\n")
			}
			currentGroup = asset.Group
			out.WriteString("<div class=\"xal-ref-grid\">\n")
		}
		writeCard(&out, page.Start+i, asset)
	}
	if currentGroup != "" {
		out.WriteString("</div>\n")
	}
	writeFile(filepath.Join(outputDir, page.File), out.Bytes())
}

func writeGroupSamples(renderer usecase.RenderUsecase, group assetGroup) {
	for _, asset := range group.Assets {
		xalPath := iconSamplePath(asset)
		svgPath := iconPreviewPath(asset)
		writeFile(xalPath, []byte(asset.Sample))
		renderSVG(renderer, asset.Sample, svgPath)
	}
}

func writeFile(path string, data []byte) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		fatal(err)
	}
}

func loadCatalog(path string) (map[string]catalogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	entries := make(map[string]catalogEntry, len(records))
	for i, record := range records {
		if i == 0 || len(record) < 5 {
			continue
		}
		entries[filepath.ToSlash(record[4])] = catalogEntry{
			ID:       record[0],
			Category: record[1],
			Service:  record[2],
		}
	}
	return entries, nil
}

func collectSVGAssets(root string, catalog map[string]catalogEntry) ([]svgAsset, error) {
	var paths []string
	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || strings.ToLower(filepath.Ext(path)) != ".svg" {
			return nil
		}
		paths = append(paths, filepath.ToSlash(path))
		return nil
	}); err != nil {
		return nil, err
	}
	sort.Strings(paths)
	assets := make([]svgAsset, 0, len(paths))
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		entry, ok := catalog[path]
		usage := fallbackUsage(path)
		if ok {
			usage = usageSnippet(path, entry)
		}
		assets = append(assets, svgAsset{
			Path:     path,
			Name:     strings.TrimSuffix(filepath.Base(path), ".svg"),
			Group:    groupName(path),
			Slug:     assetSlug(path),
			Usage:    usage,
			Sample:   sampleXAL(path, usage),
			SizeByte: int(info.Size()),
		})
	}
	return assets, nil
}

func prepareGeneratedDirs() {
	dirs := []string{
		outputDir,
		sampleDir,
		previewDir,
		filepath.Join(referenceRoot, "samples", "arrows"),
		filepath.Join(referenceRoot, "samples", "frames"),
		filepath.Join(referenceRoot, "previews", "arrows"),
		filepath.Join(referenceRoot, "previews", "frames"),
	}
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			fatal(err)
		}
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fatal(err)
		}
	}
}

func groupName(path string) string {
	const prefix = "etc/resources/aws/svg/"
	rel := strings.TrimPrefix(filepath.ToSlash(path), prefix)
	parts := strings.Split(rel, "/")
	if len(parts) == 0 || parts[0] == "" {
		return "SVG Assets"
	}
	return parts[0]
}

func writeHeader(out *bytes.Buffer, group assetGroup, page assetPage) {
	title := group.Name
	if page.File != group.Slug+".md" {
		title = fmt.Sprintf("%s - Page %d", group.Name, page.Number)
	}
	fmt.Fprintf(out, `# %s

This page references SVG files under `+"`etc/resources/aws/svg/%s`"+`.
Each card shows the SVG preview first and the XAL tag syntax in the Code tab.
Showing %d-%d of %d SVG files.

<style>
.xal-ref-grid{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:1rem;margin:1rem 0 1.5rem}.xal-ref-card{border:1px solid var(--table-border-color);border-radius:8px;overflow:hidden;background:var(--bg);padding:.75rem}.xal-ref-title{font-size:.9rem;font-weight:600;line-height:1.25;margin:0 0 .5rem;overflow-wrap:anywhere}.xal-ref-preview{display:flex;align-items:center;justify-content:center}.xal-ref-preview img{display:block;width:100%%;height:auto}.xal-ref-card pre{margin:0;max-height:18rem;overflow:auto;white-space:pre}.xal-ref-card code{font-size:.78em}.xal-ref-tag{margin:.5rem 0 0;color:var(--fg);opacity:.72;overflow:auto;white-space:pre}.xal-ref-tag code{font-size:.72em}@media(max-width:900px){.xal-ref-grid{grid-template-columns:repeat(2,minmax(0,1fr))}}@media(max-width:560px){.xal-ref-grid{grid-template-columns:1fr}}
</style>

`, title, group.Name, page.Start+1, page.End, len(group.Assets))
	if page.File != group.Slug+".md" {
		out.WriteString("[Back to section index](" + group.Slug + ".md)\n\n")
	}
}

func writeCard(out *bytes.Buffer, i int, asset svgAsset) {
	tabID := fmt.Sprintf("svg-ref-%04d", i)
	preview := "../previews/icons/" + asset.Group + "/" + asset.Slug + ".svg"
	sample := "../samples/icons/" + asset.Group + "/" + asset.Slug + ".xal"
	fmt.Fprintf(out, `<div class="xal-ref-card">
<div class="xal-ref-title">%s</div>

{{#tabs name="%s"}}
{{#tab name="Preview"}}
<div class="xal-ref-preview">
<div>
<img src="%s" alt="%s">
</div>
</div>

{{#endtab}}
{{#tab name="Code"}}

`+"```xml"+`
{{#include %s}}
`+"```"+`

{{#endtab}}
{{#endtabs}}

<div class="xal-ref-tag"><code>%s</code></div>

</div>
`, escapeHTML(asset.Name), tabID, escapeHTML(preview), escapeHTML(asset.Name), sample, escapeHTML(asset.Usage))
}

func usageSnippet(path string, entry catalogEntry) string {
	if strings.HasPrefix(path, groupIconRelPath) {
		return fmt.Sprintf(`<generic-group id="%s" title="%s" icon-id="%s" />`, slug(entry.Service), entry.Service, entry.ID)
	}
	return fmt.Sprintf(`<item id="%s" />`, entry.ID)
}

func fallbackUsage(path string) string {
	base := strings.TrimPrefix(strings.TrimSuffix(filepath.Base(path), ".svg"), "isoflow-")
	var id strings.Builder
	for _, r := range base {
		if r < '0' || r > '9' {
			break
		}
		id.WriteRune(r)
	}
	if id.Len() > 0 {
		return fmt.Sprintf(`<item id="%s" />`, id.String())
	}
	return fmt.Sprintf(`<!-- No service-catalog entry found for %s. -->`, path)
}

func sampleXAL(path, usage string) string {
	width := 220
	height := 160
	if strings.HasPrefix(path, groupIconRelPath) {
		width, height = groupFrameSize(extractTitle(usage), true)
		return fmt.Sprintf(`<frame version="1" width="%d" height="%d" margin="%d">
  %s
</frame>
`, width, height, sampleMargin, indentUsage(usage, "  "))
	}
	return fmt.Sprintf(`<frame version="1" width="%d" height="%d" margin="%d" item-size="64">
  %s
</frame>
`, width, height, sampleMargin, indentUsage(usage, "  "))
}

type referenceSample struct {
	Section string
	Slug    string
	Title   string
	Alt     string
	XAL     string
}

type referencePage struct {
	Section string
	File    string
	Title   string
	Samples []referenceSample
}

func writeReferenceSamples(renderer usecase.RenderUsecase) {
	pages := []referencePage{
		{
			Section: "arrows",
			File:    "connector-types.md",
			Title:   "Connector Types",
			Samples: []referenceSample{
				arrowSample("normal", "Normal Connection", "Normal connection preview", `<connection src="web" dst="app" />`),
				arrowSample("route", "Route", "Route connection preview", `<connection src="web" dst="app" kind="route" />`),
				arrowSample("traffic", "Traffic", "Traffic connection preview", `<connection src="web" dst="app" kind="route" />
  <connection src="web" dst="app"
              kind="traffic"
              stroke-style="dashed" />`),
			},
		},
		{
			Section: "arrows",
			File:    "anchors-bends.md",
			Title:   "Anchors and Bends",
			Samples: []referenceSample{
				arrowSample("anchor", "Explicit Anchors", "Anchor connection preview", `<connection src="web" dst="app"
              src-anchor="right-3"
              dst-anchor="left-3" />`),
				manualBendSample(),
			},
		},
		{
			Section: "frames",
			File:    "frame-containers.md",
			Title:   "Frame and Container Shapes",
			Samples: []referenceSample{
				{
					Section: "frames",
					Slug:    "root",
					Title:   "Root Frame",
					Alt:     "Root frame preview",
					XAL: `<frame version="1" width="360" height="220" margin="24">
  <container layout="horizontal" gap="16">
    <card title="Web" />
    <card title="App" />
    <card title="Data" />
  </container>
</frame>
`,
				},
				{
					Section: "frames",
					Slug:    "leaf",
					Title:   "Generic Leaf Box",
					Alt:     "Generic leaf box preview",
					XAL: `<frame version="1" width="240" height="160" margin="24">
  <card title="Dashboard" />
</frame>
`,
				},
			},
		},
		{
			Section: "frames",
			File:    "aws-tags.md",
			Title:   "AWS Tags",
			Samples: awsTagSamples(),
		},
		{
			Section: "frames",
			File:    "rectangles-ports.md",
			Title:   "Rectangles and Ports",
			Samples: []referenceSample{
				{
					Section: "frames",
					Slug:    "rectangle-ports",
					Title:   "Rectangle With Ports",
					Alt:     "Rectangle and ports preview",
					XAL: `<frame version="1" width="300" height="180" margin="32">
  <rectangle id="service" title="Service" width="180" height="100">
    <port id="service-in" side="left" title="in" />
    <port id="service-out" side="right" title="out" />
  </rectangle>
</frame>
`,
				},
				{
					Section: "frames",
					Slug:    "borderless",
					Title:   "Borderless Leaf",
					Alt:     "Borderless leaf preview",
					XAL: `<frame version="1" width="260" height="160" margin="32">
  <panel title="No Border" border="none" />
</frame>
`,
				},
			},
		},
	}
	for _, page := range pages {
		for _, sample := range page.Samples {
			writeFile(referenceSamplePath(sample), []byte(sample.XAL))
			renderSVG(renderer, sample.XAL, referencePreviewPath(sample))
		}
		writeReferencePage(page)
	}
}

func arrowSample(slug, title, alt, connections string) referenceSample {
	return referenceSample{
		Section: "arrows",
		Slug:    slug,
		Title:   title,
		Alt:     alt,
		XAL: fmt.Sprintf(`<frame version="1" width="360" height="160" margin="24" layout="horizontal" gap="120" align="middle-center">
  <rectangle id="web" title="web" width="80" height="48" />
  <rectangle id="app" title="app" width="80" height="48" />
  %s
</frame>
`, connections),
	}
}

func manualBendSample() referenceSample {
	return referenceSample{
		Section: "arrows",
		Slug:    "bend",
		Title:   "Manual Bends",
		Alt:     "Manual bend preview",
		XAL: `<frame version="1" width="360" height="200" margin="24" layout="horizontal" gap="120" align="middle-center">
  <rectangle id="web" title="web" width="80" height="48" />
  <rectangle id="db" title="db" width="80" height="48" />
  <connection src="web" dst="db" grid="8">
    <bend x="180" y="64" />
    <bend x="180" y="136" />
  </connection>
</frame>
`,
	}
}

func awsTagSamples() []referenceSample {
	tags := []struct {
		slug  string
		title string
		alt   string
		tag   string
		id    string
		label string
		extra string
		icon  bool
	}{
		{"aws-cloud", "AWS Cloud", "AWS Cloud tag preview", "aws-cloud", "aws", "AWS Cloud", "", true},
		{"aws-cloud-alt", "AWS Cloud Alt", "AWS Cloud alternate tag preview", "aws-cloud-alt", "aws-alt", "AWS Cloud", "", true},
		{"aws-account", "AWS Account", "AWS account tag preview", "aws-account", "account", "AWS Account", "", true},
		{"region", "Region", "Region tag preview", "region", "apne1", "ap-northeast-1", "", true},
		{"availability-zone", "Availability Zone", "Availability Zone tag preview", "availability-zone", "az-a", "ap-northeast-1a", "", false},
		{"vpc", "VPC", "VPC tag preview", "vpc", "prod-vpc", "Production VPC", "", true},
		{"public-subnet", "Public Subnet", "Public subnet tag preview", "public-subnet", "public", "Public Subnet", "", true},
		{"private-subnet", "Private Subnet", "Private subnet tag preview", "private-subnet", "private", "Private Subnet", "", true},
		{"security-group", "Security Group", "Security group tag preview", "security-group", "sg-web", "Security Group", "", false},
		{"auto-scaling-group", "Auto Scaling Group", "Auto Scaling group tag preview", "auto-scaling-group", "asg", "Auto Scaling group", "", true},
		{"ec2-instance-contents", "EC2 Instance Contents", "EC2 instance contents tag preview", "ec2-instance-contents", "ec2-box", "EC2 instance contents", "", true},
		{"server-contents", "Server Contents", "Server contents tag preview", "server-contents", "server-box", "Server contents", "", true},
		{"corporate-data-center", "Corporate Data Center", "Corporate data center tag preview", "corporate-data-center", "dc", "Corporate data center", "", true},
		{"spot-fleet", "Spot Fleet", "Spot Fleet tag preview", "spot-fleet", "spot", "Spot Fleet", "", true},
		{"aws-iot-greengrass-deployment", "AWS IoT Greengrass Deployment", "AWS IoT Greengrass Deployment tag preview", "aws-iot-greengrass-deployment", "greengrass-deploy", "Greengrass Deployment", "", true},
		{"aws-iot-greengrass", "AWS IoT Greengrass", "AWS IoT Greengrass tag preview", "aws-iot-greengrass", "greengrass", "AWS IoT Greengrass", "", false},
		{"elastic-beanstalk-container", "Elastic Beanstalk Container", "Elastic Beanstalk container tag preview", "elastic-beanstalk-container", "beanstalk", "Elastic Beanstalk", "", false},
		{"aws-step-functions-workflow", "AWS Step Functions Workflow", "AWS Step Functions workflow tag preview", "aws-step-functions-workflow", "workflow", "Workflow", "", false},
		{"generic-group", "Generic Group", "Generic group tag preview", "generic-group", "custom", "Custom Group", ` icon-id="1178"`, true},
	}
	samples := make([]referenceSample, 0, len(tags))
	for _, tag := range tags {
		width, height := groupFrameSize(tag.label, tag.icon)
		samples = append(samples, referenceSample{
			Section: "frames",
			Slug:    tag.slug,
			Title:   tag.title,
			Alt:     tag.alt,
			XAL: fmt.Sprintf(`<frame version="1" width="%d" height="%d" margin="%d">
  <%s id="%s" title="%s"%s />
</frame>
`, width, height, sampleMargin, tag.tag, tag.id, tag.label, tag.extra),
		})
	}
	return samples
}

func groupFrameSize(title string, hasIcon bool) (int, int) {
	groupW := groupSampleMinW
	if want := int(math.Ceil(groupHeaderWidth(title, hasIcon))) + 12; want > groupW {
		groupW = want
	}
	return groupW + sampleMargin*2, groupSampleHeight + sampleMargin*2
}

func groupHeaderWidth(title string, hasIcon bool) float64 {
	headerH := float64(20)
	textX := float64(4)
	if hasIcon {
		headerH = 32
		textX = 32 + 4
	}
	headerTip := math.Min(14, headerH/2)
	return textX + float64(visualColumns(title))*9.6 + 18 + headerTip
}

func extractTitle(source string) string {
	const key = `title="`
	start := strings.Index(source, key)
	if start < 0 {
		return ""
	}
	start += len(key)
	end := strings.Index(source[start:], `"`)
	if end < 0 {
		return ""
	}
	return source[start : start+end]
}

func visualColumns(s string) int {
	cols := 0
	for _, r := range s {
		if r > 0xFF {
			cols += 2
			continue
		}
		cols++
	}
	return cols
}

func writeReferencePage(page referencePage) {
	var out bytes.Buffer
	fmt.Fprintf(&out, "# %s\n\n", page.Title)
	for _, sample := range page.Samples {
		writeReferenceCard(&out, sample)
	}
	writeFile(filepath.Join(referenceRoot, page.Section, page.File), out.Bytes())
}

func writeReferenceCard(out *bytes.Buffer, sample referenceSample) {
	fmt.Fprintf(out, `## %s

{{#tabs name="%s-%s"}}
{{#tab name="Preview"}}

![%s](../previews/%s/%s.svg)

{{#endtab}}
{{#tab name="Code"}}

`+"```xml"+`
{{#include ../samples/%s/%s.xal}}
`+"```"+`

{{#endtab}}
{{#endtabs}}

`, sample.Title, sample.Section, sample.Slug, sample.Alt, sample.Section, sample.Slug, sample.Section, sample.Slug)
}

func referenceSamplePath(sample referenceSample) string {
	return filepath.Join(referenceRoot, "samples", sample.Section, sample.Slug+".xal")
}

func referencePreviewPath(sample referenceSample) string {
	return filepath.Join(referenceRoot, "previews", sample.Section, sample.Slug+".svg")
}

func indentUsage(usage, prefix string) string {
	lines := strings.Split(strings.TrimSpace(usage), "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func iconSamplePath(asset svgAsset) string {
	return filepath.Join(sampleDir, asset.Group, asset.Slug+".xal")
}

func iconPreviewPath(asset svgAsset) string {
	return filepath.Join(previewDir, asset.Group, asset.Slug+".svg")
}

func renderSVG(renderer usecase.RenderUsecase, source, outputPath string) {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		fatal(err)
	}
	out, err := renderer.RenderSVG(context.Background(), []byte(source), entity.RenderOptions{
		Format: usecase.FormatSVG,
		Mode:   usecase.ModeStandard,
		Theme:  "light",
		Assets: &entity.AssetSource{
			FS:               awsassets.Assets,
			CatalogCSV:       awsassets.CatalogCSV,
			GroupIconsDir:    awsassets.GroupIconsDir,
			IsoflowIconsJSON: awsassets.IsoflowIconsJSON,
			ItemIconSize:     32,
		},
		PxPerInch: 96,
	})
	if err != nil {
		fatal(fmt.Errorf("render %s: %w", outputPath, err))
	}
	writeFile(outputPath, out)
}

func newUsecase() usecase.RenderUsecase {
	return usecase.NewRenderUsecase(
		repository.NewExcalidrawRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewIsoflowRepository(),
		repository.NewSVGRepository(),
		repository.NewXYFlowRepository(),
		repository.NewPDFRepository(),
		repository.NewSpreadsheetRepository(),
	)
}

func slug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	lastDash := false
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			lastDash = false
			continue
		}
		if !lastDash {
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func assetSlug(path string) string {
	base := slug(strings.TrimSuffix(filepath.Base(path), ".svg"))
	sum := crc32.ChecksumIEEE([]byte(filepath.ToSlash(path)))
	return fmt.Sprintf("%s-%08x", base, sum)
}

func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&#34;",
		"'", "&#39;",
	)
	return replacer.Replace(s)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
