package svg

import (
	"math"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawRhombus(
	gongsvgStage *gongsvg_models.StageStruct,
	axisLayer *gongsvg_models.Layer,
	p *phylotaxymusic_models.Parameter,
	r *phylotaxymusic_models.Rhombus,
) {

	// Convert angle from degrees to radians for computation
	angleRad := r.Angle * math.Pi / 180
	insideAngleRad := r.InsideAngle * math.Pi / 180

	sinInsideAngle := math.Sin(insideAngleRad / 2)
	cosInsideAngle := math.Cos(insideAngleRad / 2)

	// Calculate half of the inside angles to position the vertices
	halfVerticalDiagonal := r.SideLength * sinInsideAngle
	halfHorizontalDiagonal := r.SideLength * cosInsideAngle

	var x_s [4]float64
	var y_s [4]float64

	// Coordinates of the first vertex (using the angle + half of the inside angle)
	x_s[0] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad)
	y_s[0] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad)

	// Coordinates of the second vertex (using the angle - half of the inside angle)
	x_s[1] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad+math.Pi/2.0)
	y_s[1] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad+math.Pi/2.0)

	// Coordinates of the third vertex (opposite of the first vertex)
	x_s[2] = r.CenterX + halfVerticalDiagonal*math.Cos(angleRad+math.Pi)
	y_s[2] = r.CenterY + halfVerticalDiagonal*math.Sin(angleRad+math.Pi)

	// Coordinates of the fourth vertex (opposite of the second vertex)
	x_s[3] = r.CenterX + halfHorizontalDiagonal*math.Cos(angleRad+math.Pi*1.5)
	y_s[3] = r.CenterY + halfHorizontalDiagonal*math.Sin(angleRad+math.Pi*1.5)

	for i := range 4 {
		line := (&gongsvg_models.Line{
			X1: x_s[i%4] + p.OriginX,
			Y1: p.OriginY - y_s[i%4],
			X2: x_s[(i+1)%4] + p.OriginX,
			Y2: p.OriginY - y_s[(i+1)%4],
			Presentation: gongsvg_models.Presentation{
				Stroke:        gongsvg_models.Black.ToString(),
				StrokeWidth:   r.StrokeWidth,
				StrokeOpacity: 1,
			},
		}).Stage(gongsvgStage)

		axisLayer.Lines = append(axisLayer.Lines, line)
	}

}