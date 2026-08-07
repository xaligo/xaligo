package entity

import (
	"io"
	"time"
)

type ControllerDiffOptions struct {
	BeforePath   string
	AfterPath    string
	OutputPrefix string
	Theme        string
	Mode         string
	PxPerInch    float64
	Stdout       io.Writer
}

type ControllerRenderOptions struct {
	InputPath         string
	OutputPath        string
	Format            string
	ServicesFile      string
	Title             string
	Author            string
	Company           string
	Subject           string
	CombineFrames     bool
	Compression       *bool
	PxPerInch         float64
	ArrowStyle        string
	ArrowStub         float64
	ArrowMargin       float64
	Paper             string
	Orientation       string
	PaperMargin       float64
	PaperMarginTop    float64
	PaperMarginRight  float64
	PaperMarginBottom float64
	PaperMarginLeft   float64
	SVGLegendPosition string
	Theme             string
	Mode              string
}

type ControllerServeOptions struct {
	InputPath    string
	Address      string
	Port         int
	PortSet      bool
	Mode         string
	Theme        string
	Paper        string
	Orientation  string
	PollInterval time.Duration
}

type ControllerRenderMarkdownOptions struct {
	InputPath    string
	OutputPath   string
	SVGDir       string
	ServicesFile string
	PxPerInch    float64
	Theme        string
	Mode         string
	Paper        string
	Orientation  string
}
