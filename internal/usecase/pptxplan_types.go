package usecase

import "github.com/ryo-arima/xaligo/internal/entity"

// Type aliases keep the calculation package focused on behavior while all
// cross-layer data structures live in the entity layer.
type Scene = entity.PptxScene
type AppState = entity.AppState
type SceneFile = entity.SceneFile
type Binding = entity.Binding
type Element = entity.Element
type CustomData = entity.CustomData
type Options = entity.PptxOptions
type Plan = entity.Plan
type LegendEntry = entity.LegendEntry
type ConnectorLegendEntry = entity.ConnectorLegendEntry
type PlanSlide = entity.PlanSlide
type DrawOp = entity.DrawOp
type LineStyle = entity.LineStyle
type FillStyle = entity.FillStyle
type PtIn = entity.PtIn
