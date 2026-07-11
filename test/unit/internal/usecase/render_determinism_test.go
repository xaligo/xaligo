package usecase_test

import (
	"bytes"
	"context"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestRenderExcalidrawIsExactlyRepeatableV1(t *testing.T) {
	source := []byte(`<frame width="640" height="240" layout="horizontal" gap="20" class="pa-2">
  <generic-group id="left" title="Left" col="1">
    <item id="27" name="web" />
  </generic-group>
  <generic-group id="right" title="Right" col="1">
    <item id="117" name="db" />
  </generic-group>
  <connection src="web" dst="db" />
</frame>`)
	opts := entity.RenderOptions{
		Format: usecase.FormatExcalidraw,
		Theme:  "light",
		Assets: &entity.AssetSource{
			FS:               awsassets.Assets,
			CatalogCSV:       awsassets.CatalogCSV,
			GroupIconsDir:    awsassets.GroupIconsDir,
			IsoflowIconsJSON: awsassets.IsoflowIconsJSON,
			ItemIconSize:     32,
		},
	}
	renderer := newUsecase()
	first, err := renderer.RenderExcalidraw(context.Background(), source, opts)
	if err != nil {
		t.Fatalf("first render: %v", err)
	}
	second, err := renderer.RenderExcalidraw(context.Background(), source, opts)
	if err != nil {
		t.Fatalf("second render: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("identical V1 source, options, and assets produced different Excalidraw bytes")
	}

	nativeOpts := opts
	nativeOpts.Assets = nil
	native, err := renderer.RenderExcalidraw(context.Background(), source, nativeOpts)
	if err != nil {
		t.Fatalf("native-default render: %v", err)
	}
	if !bytes.Equal(first, native) {
		t.Fatal("embedded V1 assets at 32px produced output different from native V1 defaults")
	}
}

func TestRenderExplicitV1VersionMatchesUnversionedV1(t *testing.T) {
	renderer := newUsecase()
	opts := entity.RenderOptions{Format: usecase.FormatExcalidraw, Theme: "light"}
	unversioned := []byte(`<frame width="240" height="120"><rectangle id="service" title="Service" /></frame>`)
	explicitV1 := []byte(`<frame version="1" width="240" height="120"><rectangle id="service" title="Service" /></frame>`)

	want, err := renderer.RenderExcalidraw(context.Background(), unversioned, opts)
	if err != nil {
		t.Fatalf("unversioned V1 render: %v", err)
	}
	got, err := renderer.RenderExcalidraw(context.Background(), explicitV1, opts)
	if err != nil {
		t.Fatalf("explicit V1 render: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal(`version="1" changed the rendered V1 output`)
	}
}

func TestRenderV1CompatibilityModesAreEquivalentV1(t *testing.T) {
	renderer := newUsecase()
	source := []byte(`<frame width="320" height="120" layout="horizontal"><rectangle id="left" title="Left" /><rectangle id="right" title="Right" /><connection src="left" dst="right" /></frame>`)
	modes := []entity.Mode{usecase.ModeStandard, usecase.ModeNetwork, usecase.ModeAWS}
	var want []byte
	for _, mode := range modes {
		got, err := renderer.RenderExcalidraw(context.Background(), source, entity.RenderOptions{
			Format: usecase.FormatExcalidraw,
			Theme:  "light",
			Mode:   mode,
		})
		if err != nil {
			t.Fatalf("mode %q render: %v", mode, err)
		}
		if want == nil {
			want = got
			continue
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("V1 compatibility mode %q changed the rendered output", mode)
		}
	}
}
