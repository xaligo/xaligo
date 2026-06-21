package controller

import (
	"os"

	"github.com/ryo-arima/xaligo/internal/usecase"
)

type RenderOptions struct {
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

type PptxGenerateOptions struct {
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

func defaultUseCase(uc usecase.API) usecase.API {
	if uc != nil {
		return uc
	}
	return usecase.New(usecase.Dependencies{})
}
