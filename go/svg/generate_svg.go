package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func GenerateSvg(gongsvgStage *gongsvg_models.StageStruct, parameter *phylotaxymusic_models.Parameter) {

	gongsvgStage.Reset()

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvgStage)
	layer := (&gongsvg_models.Layer{Name: "Layer 1"}).Stage(gongsvgStage)
	layer.Display = true
	svg.Layers = append(svg.Layers, layer)

	if parameter.HorizontalAxis.IsDisplayed {
		drawHorizontalAxis(gongsvgStage, layer, parameter, parameter.HorizontalAxis)
	}
	if parameter.VerticalAxis.IsDisplayed {
		drawVerticalAxis(gongsvgStage, layer, parameter, parameter.VerticalAxis)
	}
	if parameter.InitialRhombus.IsDisplayed {
		drawRhombus(gongsvgStage, layer, parameter, parameter.InitialRhombus)
	}
	if parameter.InitialCircle.IsDisplayed {
		drawCircle(gongsvgStage, layer, parameter, parameter.InitialCircle)
	}
	if parameter.InitialRhombusGrid.IsDisplayed {
		drawRhombusGrid(gongsvgStage, layer, parameter, parameter.InitialRhombusGrid)
	}
	if parameter.InitialCircleGrid.IsDisplayed {
		drawCircleGrid(gongsvgStage, layer, parameter, parameter.InitialCircleGrid)
	}
	if parameter.InitialAxis.IsDisplayed {
		drawAxis(gongsvgStage, layer, parameter, parameter.InitialAxis)
	}
	if parameter.RotatedAxis.IsDisplayed {
		drawAxis(gongsvgStage, layer, parameter, parameter.RotatedAxis)
	}
	if parameter.RotatedRhombus.IsDisplayed {
		drawRhombus(gongsvgStage, layer, parameter, parameter.RotatedRhombus)
	}
	if parameter.RotatedRhombusGrid.IsDisplayed {
		drawRhombusGrid(gongsvgStage, layer, parameter, parameter.RotatedRhombusGrid)
	}
	if parameter.RotatedCircleGrid.IsDisplayed {
		drawCircleGrid(gongsvgStage, layer, parameter, parameter.RotatedCircleGrid)
	}
	if parameter.NextRhombus.IsDisplayed {
		drawRhombus(gongsvgStage, layer, parameter, parameter.NextRhombus)
	}
	if parameter.NextCircle.IsDisplayed {
		drawCircle(gongsvgStage, layer, parameter, parameter.NextCircle)
	}
	if parameter.GrowingRhombusGrid.IsDisplayed {
		drawRhombusGrid(gongsvgStage, layer, parameter, parameter.GrowingRhombusGrid)
	}
	gongsvgStage.Commit()

}