// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Axis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AngleDegree float64
	Length float64
	CenterX float64
	CenterY float64
	EndX float64
	EndY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Axis) CopyBasicFields(to *Axis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AngleDegree = from.AngleDegree
	to.Length = from.Length
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type AxisGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *AxisGrid) CopyBasicFields(to *AxisGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type Bezier_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	StartX float64
	StartY float64
	ControlPointStartX float64
	ControlPointStartY float64
	EndX float64
	EndY float64
	ControlPointEndX float64
	ControlPointEndY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Bezier) CopyBasicFields(to *Bezier) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.ControlPointStartX = from.ControlPointStartX
	to.ControlPointStartY = from.ControlPointStartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.ControlPointEndX = from.ControlPointEndX
	to.ControlPointEndY = from.ControlPointEndY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type BezierGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *BezierGrid) CopyBasicFields(to *BezierGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type BezierGridStack_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *BezierGridStack) CopyBasicFields(to *BezierGridStack) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type Circle_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	HasBespokeRadius bool
	BespopkeRadius float64
	Pitch int
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	ShowName bool
}

func (from *Circle) CopyBasicFields(to *Circle) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.HasBespokeRadius = from.HasBespokeRadius
	to.BespopkeRadius = from.BespopkeRadius
	to.Pitch = from.Pitch
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.ShowName = from.ShowName
}

type CircleGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *CircleGrid) CopyBasicFields(to *CircleGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type FrontCurve_WOP struct {
	// insertion point
	Name string
	Path string
}

func (from *FrontCurve) CopyBasicFields(to *FrontCurve) {
	// insertion point
	to.Name = from.Name
	to.Path = from.Path
}

type FrontCurveStack_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *FrontCurveStack) CopyBasicFields(to *FrontCurveStack) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type HorizontalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *HorizontalAxis) CopyBasicFields(to *HorizontalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Key_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	Path string
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Key) CopyBasicFields(to *Key) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.Path = from.Path
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type NoteInfo_WOP struct {
	// insertion point
	Name string
	IsKept bool
}

func (from *NoteInfo) CopyBasicFields(to *NoteInfo) {
	// insertion point
	to.Name = from.Name
	to.IsKept = from.IsKept
}

type Parameter_WOP struct {
	// insertion point
	Name string
	BackendColor string
	MinuteColor string
	HourColor string
	N int
	M int
	Z int
	ShiftToNearestCircle int
	InsideAngle float64
	SideLength float64
	StackWidth int
	NbShitRight int
	StackHeight int
	BezierControlLengthRatio float64
	SpiralBezierStrength float64
	NbInterpolationPoints int
	HourHandleRotationAngle float64
	HourHandleDiskDistance float64
	HourHandleRadius float64
	MinuteHandleRotationAngle float64
	MinuteHandleDiskDistance float64
	MinuteHandleRadius float64
	MinuteOffset float64
	BackendHandleRotationAngle float64
	BackendHandleDiskDistance float64
	BackendHandleRadius float64
	BackendOffset float64
	FkeySizeRatio float64
	FkeyOriginRelativeX float64
	FkeyOriginRelativeY float64
	PitchHeight float64
	NbPitchLines int
	MeasureLinesHeightRatio float64
	NbMeasureLines int
	NbMeasureLinesPerCurve int
	FirstVoiceShiftX float64
	FirstVoiceShiftY float64
	PitchDifference int
	Speed float64
	Level float64
	IsMinor bool
	OriginX float64
	OriginY float64
	SpiralOriginX float64
	SpiralOriginY float64
	OriginCrossWidth float64
	SpiralRadiusRatio float64
	ShowSpiralBezierConstruct bool
	ShowInterpolationPoints bool
}

func (from *Parameter) CopyBasicFields(to *Parameter) {
	// insertion point
	to.Name = from.Name
	to.BackendColor = from.BackendColor
	to.MinuteColor = from.MinuteColor
	to.HourColor = from.HourColor
	to.N = from.N
	to.M = from.M
	to.Z = from.Z
	to.ShiftToNearestCircle = from.ShiftToNearestCircle
	to.InsideAngle = from.InsideAngle
	to.SideLength = from.SideLength
	to.StackWidth = from.StackWidth
	to.NbShitRight = from.NbShitRight
	to.StackHeight = from.StackHeight
	to.BezierControlLengthRatio = from.BezierControlLengthRatio
	to.SpiralBezierStrength = from.SpiralBezierStrength
	to.NbInterpolationPoints = from.NbInterpolationPoints
	to.HourHandleRotationAngle = from.HourHandleRotationAngle
	to.HourHandleDiskDistance = from.HourHandleDiskDistance
	to.HourHandleRadius = from.HourHandleRadius
	to.MinuteHandleRotationAngle = from.MinuteHandleRotationAngle
	to.MinuteHandleDiskDistance = from.MinuteHandleDiskDistance
	to.MinuteHandleRadius = from.MinuteHandleRadius
	to.MinuteOffset = from.MinuteOffset
	to.BackendHandleRotationAngle = from.BackendHandleRotationAngle
	to.BackendHandleDiskDistance = from.BackendHandleDiskDistance
	to.BackendHandleRadius = from.BackendHandleRadius
	to.BackendOffset = from.BackendOffset
	to.FkeySizeRatio = from.FkeySizeRatio
	to.FkeyOriginRelativeX = from.FkeyOriginRelativeX
	to.FkeyOriginRelativeY = from.FkeyOriginRelativeY
	to.PitchHeight = from.PitchHeight
	to.NbPitchLines = from.NbPitchLines
	to.MeasureLinesHeightRatio = from.MeasureLinesHeightRatio
	to.NbMeasureLines = from.NbMeasureLines
	to.NbMeasureLinesPerCurve = from.NbMeasureLinesPerCurve
	to.FirstVoiceShiftX = from.FirstVoiceShiftX
	to.FirstVoiceShiftY = from.FirstVoiceShiftY
	to.PitchDifference = from.PitchDifference
	to.Speed = from.Speed
	to.Level = from.Level
	to.IsMinor = from.IsMinor
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.SpiralOriginX = from.SpiralOriginX
	to.SpiralOriginY = from.SpiralOriginY
	to.OriginCrossWidth = from.OriginCrossWidth
	to.SpiralRadiusRatio = from.SpiralRadiusRatio
	to.ShowSpiralBezierConstruct = from.ShowSpiralBezierConstruct
	to.ShowInterpolationPoints = from.ShowInterpolationPoints
}

type Rhombus_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	SideLength float64
	AngleDegree float64
	InsideAngle float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Rhombus) CopyBasicFields(to *Rhombus) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.SideLength = from.SideLength
	to.AngleDegree = from.AngleDegree
	to.InsideAngle = from.InsideAngle
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type RhombusGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *RhombusGrid) CopyBasicFields(to *RhombusGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type ShapeCategory_WOP struct {
	// insertion point
	Name string
	IsExpanded bool
}

func (from *ShapeCategory) CopyBasicFields(to *ShapeCategory) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
}

type SpiralBezier_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	StartX float64
	StartY float64
	ControlPointStartX float64
	ControlPointStartY float64
	EndX float64
	EndY float64
	ControlPointEndX float64
	ControlPointEndY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *SpiralBezier) CopyBasicFields(to *SpiralBezier) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.ControlPointStartX = from.ControlPointStartX
	to.ControlPointStartY = from.ControlPointStartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.ControlPointEndX = from.ControlPointEndX
	to.ControlPointEndY = from.ControlPointEndY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type SpiralBezierGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *SpiralBezierGrid) CopyBasicFields(to *SpiralBezierGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type SpiralCircle_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	HasBespokeRadius bool
	BespopkeRadius float64
	Pitch int
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	ShowName bool
	Path string
}

func (from *SpiralCircle) CopyBasicFields(to *SpiralCircle) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.HasBespokeRadius = from.HasBespokeRadius
	to.BespopkeRadius = from.BespopkeRadius
	to.Pitch = from.Pitch
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.ShowName = from.ShowName
	to.Path = from.Path
}

type SpiralCircleGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *SpiralCircleGrid) CopyBasicFields(to *SpiralCircleGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type SpiralLine_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	StartX float64
	EndX float64
	StartY float64
	EndY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *SpiralLine) CopyBasicFields(to *SpiralLine) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.StartX = from.StartX
	to.EndX = from.EndX
	to.StartY = from.StartY
	to.EndY = from.EndY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type SpiralLineGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *SpiralLineGrid) CopyBasicFields(to *SpiralLineGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type SpiralOrigin_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *SpiralOrigin) CopyBasicFields(to *SpiralOrigin) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type SpiralRhombus_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	X_r0 float64
	Y_r0 float64
	X_r1 float64
	Y_r1 float64
	X_r2 float64
	Y_r2 float64
	X_r3 float64
	Y_r3 float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *SpiralRhombus) CopyBasicFields(to *SpiralRhombus) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.X_r0 = from.X_r0
	to.Y_r0 = from.Y_r0
	to.X_r1 = from.X_r1
	to.Y_r1 = from.Y_r1
	to.X_r2 = from.X_r2
	to.Y_r2 = from.Y_r2
	to.X_r3 = from.X_r3
	to.Y_r3 = from.Y_r3
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type SpiralRhombusGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *SpiralRhombusGrid) CopyBasicFields(to *SpiralRhombusGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type VerticalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *VerticalAxis) CopyBasicFields(to *VerticalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

