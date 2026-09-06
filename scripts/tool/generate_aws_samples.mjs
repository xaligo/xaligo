// Bootstrap editable examples; regeneration never overwrites an existing .xal.
// --render refreshes SVGs from current sources, including hand-edited examples.
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { spawn } from 'node:child_process';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const base = path.join(root, 'docs/src/examples/samples/aws');
const definitions = JSON.parse(fs.readFileSync(path.join(base, 'catalog.json'), 'utf8'));
const xml = value => String(value).replaceAll('&', '&amp;').replaceAll('"', '&quot;').replaceAll('<', '&lt;').replaceAll('>', '&gt;');
const selected = process.argv.find(arg => arg.startsWith('--tag='))?.slice(6);
const tags = selected ? definitions.filter(d => d.tag === selected) : definitions;
if (!tags.length) throw new Error('Unknown AWS tag: ' + selected);
const check = process.argv.includes('--check');
const source = definition => {
  const d = definition;
  const annotation = d.parameters[0] ? ` ${d.parameters[0].name}="${xml(d.parameters[0].example)}"` : '';
  let content;
  if (d.group) {
    content = `    <${d.tag} id="component" title="${xml(d.name)}"${annotation} width="840" height="340" margin="24">\n      <rectangle id="member" title="Member / 子要素" width="180" height="72" />\n    </${d.tag}>`;
  } else if (d.boundary) {
    content = `    <vpc id="network" title="VPC / ネットワーク" width="800" height="340" margin="48">\n      <private-subnet id="workload" title="Private subnet" width="520" height="200">\n        <aws-ec2 id="app" label="Application" />\n      </private-subnet>\n      <${d.tag} id="component" side="${d.boundary.defaultSide}" anchor="0.55"${annotation} />\n    </vpc>\n    <connections><connection src="app" dst="component" kind="traffic" /></connections>`;
  } else {
    const owner = d.scope === 'subnet' ? (d.tag === 'aws-vpc-nat-gateway' ? 'public-subnet' : 'private-subnet') : d.scope === 'region' ? 'region' : d.scope === 'global' ? 'aws-cloud' : 'generic-group';
    const title = d.scope === 'subnet' ? 'Subnet-scoped resource / サブネット内' : d.scope === 'region' ? 'Regional service / リージョン' : d.scope === 'global' ? 'Global service' : `${d.category} / logical view`;
    content = `    <${owner} id="scope" title="${xml(title)}" width="840" height="340" margin="24" layout="horizontal" gap="100">\n      <${d.tag} id="default" label-width="240" />\n      <${d.tag} id="configured" label="Example / サンプル" label-width="240"${annotation} />\n    </${owner}>`;
  }
  return `<!-- ${d.name.replaceAll('--', '—')} | ${d.tag} -->\n<!-- Editable source. See README.md for parameters and diagram scope. -->\n<xaligo version="2">\n  <frames>\n    <frame id="sample" width="960" height="500" title="${xml(d.tag)}">\n${content}\n    </frame>\n  </frames>\n</xaligo>\n`;
};
const spec = d => `# \`${d.tag}\` — ${d.name}\n\n[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)\n\n![${d.name}](sample.svg)\n\n${d.description}\n\n- Kind: \`${d.kind}\`; category: ${d.category}.\n- Diagram scope: \`${d.scope}\` (recommendation, not AWS deployment validation).\n- Default catalog ID: ${d.catalogID || 'none (text-only group)'}. Covered catalog IDs: ${d.catalogIDs.join(', ') || 'none'}.\n- Implementation: V1 and V2; ${d.group ? 'container with AWS border/header styling and connectable children' : d.boundary ? 'icon-only boundary port, excluded from normal layout flow' : 'fixed AWS icon with a wrapped label and explicit functional annotations'}.\n\n## Parameters\n\n${d.group ? 'Existing container parameters (`id`, `title`, `width`, `height`, `gap`, `class`, `layout`, `visible`) remain available.' : d.boundary ? '`id` is required. `side` = top/right/bottom/left (default '+d.boundary.defaultSide+'); `anchor` = 0..1; `offset` = finite tangent displacement; `size` > 0 (default 48). The tag must be an empty direct child of `<vpc>`. Functional annotations are preserved in source/project metadata; the boundary icon has no text label.' : '`id` is a required, unique connection ID, not a catalog number. `label`/`title`/`name` override the label; an empty label hides it. `size` > 0 defaults to 48 px. `label-width` > 0 defaults to 160 px (default box width, at least icon size + 12 px). Explicit `width`/`height` must contain the icon and label. `visible="false"` hides it. Children and icon overrides are not supported; use a group for containment.'}\n\n\`detail\` adds a free-form diagram annotation. \`show-details="false"\` hides annotation text. Only supplied values are shown; none are sent to AWS. ${d.group ? 'Group annotations are appended to the single-line header.' : 'Service/resource annotations appear on separate wrapped lines.'}\n\n| Parameter | Type | Meaning | Example |\n|---|---|---|---|\n${d.parameters.map(p => `| \`${p.name}\` | ${p.type === 'enum' ? p.values.map(v=>'\`'+v+'\`').join(' / ') : p.type} | ${p.description} | \`${p.example}\` |`).join('\n')}\n\n## Review notes\n\nThe catalog provides a baseline for per-component development, not a simulation of the AWS control plane. This component's current functional parameters are the ones listed above. Additional service-specific visual behavior can be developed here without replacing catalog IDs in diagrams. Edit \`sample.xal\`, then run:\n\n\`\`\`sh\nnpm run generate:aws-samples -- --render --tag=${d.tag}\n\`\`\`\n`;
for (const d of tags) {
  const directory = path.join(base, d.tag);
  if (check) {
    for (const file of ['sample.xal', 'sample.svg', 'README.md']) if (!fs.existsSync(path.join(directory, file))) throw new Error('Missing sample: ' + d.tag + '/' + file);
    for (const file of fs.readdirSync(directory).filter(file => file.endsWith('.xal'))) if (!fs.existsSync(path.join(directory, file.replace(/\.xal$/, '.svg')))) throw new Error('Missing SVG pair: ' + d.tag + '/' + file);
    continue;
  }
  fs.mkdirSync(directory, { recursive: true });
  for (const [file, content] of [['sample.xal', source(d)], ['README.md', spec(d)]]) {
    const target = path.join(directory, file);
    if (!fs.existsSync(target)) fs.writeFileSync(target, content);
  }
}
const index = `# AWS component catalog\n\n${definitions.length} dedicated tags cover all 1,875 bundled AWS catalog entries, including size variants, category/resource icons, and 21 group tags. Tabler/Yamaha icons are outside this AWS catalog.\n\nEach directory contains an editable \`sample.xal\`, its rendered \`sample.svg\`, and parameter/design notes. Numeric \`<item id="…">\` remains supported. New resource tags use \`aws-…\`; existing group tags and \`vpc-endpoint\` retain their names. Same-name service/resource icons are distinct (\`-resource\` suffix); catalog size variants share one tag.\n\nStart with [VPC endpoint](vpc-endpoint/README.md), [Internet Gateway](aws-vpc-internet-gateway/README.md), [NAT Gateway](aws-vpc-nat-gateway/README.md), [EC2](aws-ec2/README.md), [S3 bucket](aws-s3-bucket/README.md), and [VPC](vpc/README.md).\n\nFunctional parameters are diagram annotations, not provisioning commands or a complete AWS API schema. Scope recommendations are deliberately non-enforcing for logical service/feature icons. VPC endpoint border placement is an authoring convention: [interface endpoints create network interfaces in subnets](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html), whereas [gateway endpoints are associated with route tables](https://docs.aws.amazon.com/vpc/latest/privatelink/gateway-endpoints.html). [Internet gateways attach to VPCs](https://docs.aws.amazon.com/vpc/latest/userguide/working-with-igw.html); the [subnet-scoped NAT example](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat.html) stays inside a subnet.\n\nRegenerate the registry with \`npm run generate:aws-tags\`. Bootstrap missing examples with \`npm run generate:aws-samples\`. Refresh SVGs from current sources with \`npm run generate:aws-samples -- --render\` (optionally \`--tag=vpc-endpoint\`). Existing XAL/README files are never overwritten. \`catalog.json\` is the generated tag/parameter/asset manifest.\n\n| Tag | Component | Scope | Preview |\n|---|---|---|---|\n${definitions.map(d=>`| [\`${d.tag}\`](${d.tag}/README.md) | ${d.name} | ${d.scope} | [SVG](${d.tag}/sample.svg) |`).join('\n')}\n`;
const functionalIndex = '\n## Functional design review\n\n全877タグの機能レビューシート、型付きパラメータ表、関連リソース、API一覧、出典を各READMEへ追加しました。[調査範囲とソース](research/README.md)を参照してください。サービス文脈の設定と、アイコン自身の独立したリソース仕様を明示的に区別しています。\n\n- [ALB: listeners / rules / target groups / mTLS trust store](aws-elastic-load-balancing-application-load-balancer/README.md) · [mTLS接続図](aws-elastic-load-balancing-application-load-balancer/verify.svg)\n- [NLB: listeners / addressing / TLS / targets](aws-elastic-load-balancing-network-load-balancer/README.md) · [TLS終端](aws-elastic-load-balancing-network-load-balancer/termination.svg) · [TCPパススルー](aws-elastic-load-balancing-network-load-balancer/passthrough.svg)\n\n通常の描画は編集内容を保持します。全デザインを意図的に更新する場合だけ `npm run generate:aws-designs -- --update` を使用してください（上書き前のソースは一時ディレクトリへ退避）。詳細カードは編集可能な既存XAL部品であり、掲載したAWS APIフィールドをすべて新しいXAL属性として実装したという意味ではありません。\n\n';
if (!check) fs.writeFileSync(path.join(base, 'README.md'), index.replace('| Tag | Component | Scope | Preview |', functionalIndex + '| Tag | Component | Scope | Preview |'));
if (process.argv.includes('--render') && !check) {
  const samples = tags.flatMap(d => fs.readdirSync(path.join(base, d.tag)).filter(file => file.endsWith('.xal')).sort().map(file => ({...d, file})));
  let next = 0, completed = 0;
  const failures = [];
  const worker = async () => {
    while (next < samples.length) {
      const d = samples[next++];
      const directory = path.join(base, d.tag);
      await new Promise(resolve => {
        const child = spawn(path.join(root, '.bin/xaligo'), ['render', path.join(directory, d.file), '--format', 'svg', '--combine-frames', '-o', path.join(directory, d.file.replace(/\.xal$/, '.svg'))], { cwd: root });
        let output = '';
        child.stdout.on('data', data => { output += data; });
        child.stderr.on('data', data => { output += data; });
        child.on('error', error => { failures.push(`${d.tag}: ${error.message}`); resolve(); });
        child.on('close', code => { if (code !== 0) failures.push(`${d.tag}/${d.file}: ${output.split('\nUsage:')[0]}`); if (++completed % 100 === 0) console.log(`Rendered ${completed}/${samples.length}`); resolve(); });
      });
    }
  };
  await Promise.all(Array.from({ length: 4 }, worker));
  if (failures.length) throw new Error(failures.join('\n'));
  console.log(`Rendered ${completed} AWS sample SVGs.`);
}
console.log(`${check ? 'Checked' : 'Prepared'} ${tags.length} AWS sample directories.`);
