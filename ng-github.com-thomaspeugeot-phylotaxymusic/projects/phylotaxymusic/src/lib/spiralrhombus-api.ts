// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SpiralRhombusAPI {

	static GONGSTRUCT_NAME = "SpiralRhombus"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false
	X_r0: number = 0
	Y_r0: number = 0
	X_r1: number = 0
	Y_r1: number = 0
	X_r2: number = 0
	Y_r2: number = 0
	X_r3: number = 0
	Y_r3: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	SpiralRhombusPointersEncoding: SpiralRhombusPointersEncoding = new SpiralRhombusPointersEncoding
}

export class SpiralRhombusPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ShapeCategoryID: NullInt64 = new NullInt64 // if pointer is null, ShapeCategory.ID = 0

}
