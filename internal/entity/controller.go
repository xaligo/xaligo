package entity

import (
	"os"
	"time"
)

type ControllerRenderOptions struct {
	InputPath         string
	OutputPath        string
	Format            string
	ServicesFile      string
	Title             string
	Author            string
	Company           string
	Subject           string
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
	ExporterWASM      string
	Theme             string
	Mode              string
	Stdout            *os.File
	Stderr            *os.File
}

type ControllerPptxGenerateOptions struct {
	XalPath           string
	Output            string
	ServicesFile      string
	Title             string
	Author            string
	Company           string
	Subject           string
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
	ExporterWASM      string
	Theme             string
	Mode              string
	Stdout            *os.File
	Stderr            *os.File
}

type ControllerServeOptions struct {
	InputPath    string
	Address      string
	Mode         string
	Theme        string
	PollInterval time.Duration
}
