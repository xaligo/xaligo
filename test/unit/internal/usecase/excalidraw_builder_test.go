package usecase_test

import (
	"testing"

	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestMakeTextAndImageBuildExpectedElements(t *testing.T) {
	text := usecase.MakeText("label", 1.2, 2.6, 30.4, 10.2, "Hello", 14, "#111111", true, "center", 42)
	if text["type"] != "text" || text["id"] != "label" || text["x"] != float64(1) || text["y"] != float64(3) || text["fontStyle"] != "bold" || text["textAlign"] != "center" {
		t.Fatalf("text = %#v", text)
	}
	image := usecase.MakeImage("icon", 10, 20, 32, 32, "file-id", "#ffffff", 77)
	if image["type"] != "image" || image["fileId"] != "file-id" || image["status"] != "saved" || image["backgroundColor"] != "#ffffff" {
		t.Fatalf("image = %#v", image)
	}
}