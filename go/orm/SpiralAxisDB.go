// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SpiralAxis_sql sql.NullBool
var dummy_SpiralAxis_time time.Duration
var dummy_SpiralAxis_sort sort.Float64Slice

// SpiralAxisAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spiralaxisAPI
type SpiralAxisAPI struct {
	gorm.Model

	models.SpiralAxis_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SpiralAxisPointersEncoding SpiralAxisPointersEncoding
}

// SpiralAxisPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SpiralAxisPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64
}

// SpiralAxisDB describes a spiralaxis in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spiralaxisDB
type SpiralAxisDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spiralaxisDB.Name
	Name_Data sql.NullString

	// Declation for basic field spiralaxisDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool

	// Declation for basic field spiralaxisDB.Angle
	Angle_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.Length
	Length_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.CenterX
	CenterX_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.CenterY
	CenterY_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.Color
	Color_Data sql.NullString

	// Declation for basic field spiralaxisDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field spiralaxisDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field spiralaxisDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field spiralaxisDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field spiralaxisDB.Transform
	Transform_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SpiralAxisPointersEncoding
}

// SpiralAxisDBs arrays spiralaxisDBs
// swagger:response spiralaxisDBsResponse
type SpiralAxisDBs []SpiralAxisDB

// SpiralAxisDBResponse provides response
// swagger:response spiralaxisDBResponse
type SpiralAxisDBResponse struct {
	SpiralAxisDB
}

// SpiralAxisWOP is a SpiralAxis without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SpiralAxisWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`

	Angle float64 `xlsx:"3"`

	Length float64 `xlsx:"4"`

	CenterX float64 `xlsx:"5"`

	CenterY float64 `xlsx:"6"`

	Color string `xlsx:"7"`

	FillOpacity float64 `xlsx:"8"`

	Stroke string `xlsx:"9"`

	StrokeOpacity float64 `xlsx:"10"`

	StrokeWidth float64 `xlsx:"11"`

	StrokeDashArray string `xlsx:"12"`

	StrokeDashArrayWhenSelected string `xlsx:"13"`

	Transform string `xlsx:"14"`
	// insertion for WOP pointer fields
}

var SpiralAxis_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
	"Angle",
	"Length",
	"CenterX",
	"CenterY",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoSpiralAxisStruct struct {
	// stores SpiralAxisDB according to their gorm ID
	Map_SpiralAxisDBID_SpiralAxisDB map[uint]*SpiralAxisDB

	// stores SpiralAxisDB ID according to SpiralAxis address
	Map_SpiralAxisPtr_SpiralAxisDBID map[*models.SpiralAxis]uint

	// stores SpiralAxis according to their gorm ID
	Map_SpiralAxisDBID_SpiralAxisPtr map[uint]*models.SpiralAxis

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSpiralAxis.stage
	return
}

func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) GetDB() *gorm.DB {
	return backRepoSpiralAxis.db
}

// GetSpiralAxisDBFromSpiralAxisPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) GetSpiralAxisDBFromSpiralAxisPtr(spiralaxis *models.SpiralAxis) (spiralaxisDB *SpiralAxisDB) {
	id := backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]
	spiralaxisDB = backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[id]
	return
}

// BackRepoSpiralAxis.CommitPhaseOne commits all staged instances of SpiralAxis to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for spiralaxis := range stage.SpiralAxiss {
		backRepoSpiralAxis.CommitPhaseOneInstance(spiralaxis)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spiralaxis := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr {
		if _, ok := stage.SpiralAxiss[spiralaxis]; !ok {
			backRepoSpiralAxis.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSpiralAxis.CommitDeleteInstance commits deletion of SpiralAxis to the BackRepo
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CommitDeleteInstance(id uint) (Error error) {

	spiralaxis := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[id]

	// spiralaxis is not staged anymore, remove spiralaxisDB
	spiralaxisDB := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[id]
	query := backRepoSpiralAxis.db.Unscoped().Delete(&spiralaxisDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID, spiralaxis)
	delete(backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr, id)
	delete(backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB, id)

	return
}

// BackRepoSpiralAxis.CommitPhaseOneInstance commits spiralaxis staged instances of SpiralAxis to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CommitPhaseOneInstance(spiralaxis *models.SpiralAxis) (Error error) {

	// check if the spiralaxis is not commited yet
	if _, ok := backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]; ok {
		return
	}

	// initiate spiralaxis
	var spiralaxisDB SpiralAxisDB
	spiralaxisDB.CopyBasicFieldsFromSpiralAxis(spiralaxis)

	query := backRepoSpiralAxis.db.Create(&spiralaxisDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis] = spiralaxisDB.ID
	backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID] = spiralaxis
	backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[spiralaxisDB.ID] = &spiralaxisDB

	return
}

// BackRepoSpiralAxis.CommitPhaseTwo commits all staged instances of SpiralAxis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralaxis := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr {
		backRepoSpiralAxis.CommitPhaseTwoInstance(backRepo, idx, spiralaxis)
	}

	return
}

// BackRepoSpiralAxis.CommitPhaseTwoInstance commits {{structname }} of models.SpiralAxis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spiralaxis *models.SpiralAxis) (Error error) {

	// fetch matching spiralaxisDB
	if spiralaxisDB, ok := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[idx]; ok {

		spiralaxisDB.CopyBasicFieldsFromSpiralAxis(spiralaxis)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value spiralaxis.ShapeCategory translates to updating the spiralaxis.ShapeCategoryID
		spiralaxisDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if spiralaxis.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[spiralaxis.ShapeCategory]; ok {
				spiralaxisDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				spiralaxisDB.ShapeCategoryID.Valid = true
			}
		} else {
			spiralaxisDB.ShapeCategoryID.Int64 = 0
			spiralaxisDB.ShapeCategoryID.Valid = true
		}

		query := backRepoSpiralAxis.db.Save(&spiralaxisDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SpiralAxis intance %s", spiralaxis.Name))
		return err
	}

	return
}

// BackRepoSpiralAxis.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CheckoutPhaseOne() (Error error) {

	spiralaxisDBArray := make([]SpiralAxisDB, 0)
	query := backRepoSpiralAxis.db.Find(&spiralaxisDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spiralaxisInstancesToBeRemovedFromTheStage := make(map[*models.SpiralAxis]any)
	for key, value := range backRepoSpiralAxis.stage.SpiralAxiss {
		spiralaxisInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spiralaxisDB := range spiralaxisDBArray {
		backRepoSpiralAxis.CheckoutPhaseOneInstance(&spiralaxisDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spiralaxis, ok := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]
		if ok {
			delete(spiralaxisInstancesToBeRemovedFromTheStage, spiralaxis)
		}
	}

	// remove from stage and back repo's 3 maps all spiralaxiss that are not in the checkout
	for spiralaxis := range spiralaxisInstancesToBeRemovedFromTheStage {
		spiralaxis.Unstage(backRepoSpiralAxis.GetStage())

		// remove instance from the back repo 3 maps
		spiralaxisID := backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]
		delete(backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID, spiralaxis)
		delete(backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB, spiralaxisID)
		delete(backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr, spiralaxisID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spiralaxisDB that has been found in the DB, updates the backRepo and stages the
// models version of the spiralaxisDB
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CheckoutPhaseOneInstance(spiralaxisDB *SpiralAxisDB) (Error error) {

	spiralaxis, ok := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]
	if !ok {
		spiralaxis = new(models.SpiralAxis)

		backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID] = spiralaxis
		backRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis] = spiralaxisDB.ID

		// append model store with the new element
		spiralaxis.Name = spiralaxisDB.Name_Data.String
		spiralaxis.Stage(backRepoSpiralAxis.GetStage())
	}
	spiralaxisDB.CopyBasicFieldsToSpiralAxis(spiralaxis)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spiralaxis.Stage(backRepoSpiralAxis.GetStage())

	// preserve pointer to spiralaxisDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SpiralAxisDBID_SpiralAxisDB)[spiralaxisDB hold variable pointers
	spiralaxisDB_Data := *spiralaxisDB
	preservedPtrToSpiralAxis := &spiralaxisDB_Data
	backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[spiralaxisDB.ID] = preservedPtrToSpiralAxis

	return
}

// BackRepoSpiralAxis.CheckoutPhaseTwo Checkouts all staged instances of SpiralAxis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spiralaxisDB := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB {
		backRepoSpiralAxis.CheckoutPhaseTwoInstance(backRepo, spiralaxisDB)
	}
	return
}

// BackRepoSpiralAxis.CheckoutPhaseTwoInstance Checkouts staged instances of SpiralAxis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spiralaxisDB *SpiralAxisDB) (Error error) {

	spiralaxis := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr[spiralaxisDB.ID]

	spiralaxisDB.DecodePointers(backRepo, spiralaxis)

	return
}

func (spiralaxisDB *SpiralAxisDB) DecodePointers(backRepo *BackRepoStruct, spiralaxis *models.SpiralAxis) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field
	spiralaxis.ShapeCategory = nil
	if spiralaxisDB.ShapeCategoryID.Int64 != 0 {
		spiralaxis.ShapeCategory = backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(spiralaxisDB.ShapeCategoryID.Int64)]
	}
	return
}

// CommitSpiralAxis allows commit of a single spiralaxis (if already staged)
func (backRepo *BackRepoStruct) CommitSpiralAxis(spiralaxis *models.SpiralAxis) {
	backRepo.BackRepoSpiralAxis.CommitPhaseOneInstance(spiralaxis)
	if id, ok := backRepo.BackRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]; ok {
		backRepo.BackRepoSpiralAxis.CommitPhaseTwoInstance(backRepo, id, spiralaxis)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSpiralAxis allows checkout of a single spiralaxis (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSpiralAxis(spiralaxis *models.SpiralAxis) {
	// check if the spiralaxis is staged
	if _, ok := backRepo.BackRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]; ok {

		if id, ok := backRepo.BackRepoSpiralAxis.Map_SpiralAxisPtr_SpiralAxisDBID[spiralaxis]; ok {
			var spiralaxisDB SpiralAxisDB
			spiralaxisDB.ID = id

			if err := backRepo.BackRepoSpiralAxis.db.First(&spiralaxisDB, id).Error; err != nil {
				log.Fatalln("CheckoutSpiralAxis : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSpiralAxis.CheckoutPhaseOneInstance(&spiralaxisDB)
			backRepo.BackRepoSpiralAxis.CheckoutPhaseTwoInstance(backRepo, &spiralaxisDB)
		}
	}
}

// CopyBasicFieldsFromSpiralAxis
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsFromSpiralAxis(spiralaxis *models.SpiralAxis) {
	// insertion point for fields commit

	spiralaxisDB.Name_Data.String = spiralaxis.Name
	spiralaxisDB.Name_Data.Valid = true

	spiralaxisDB.IsDisplayed_Data.Bool = spiralaxis.IsDisplayed
	spiralaxisDB.IsDisplayed_Data.Valid = true

	spiralaxisDB.Angle_Data.Float64 = spiralaxis.Angle
	spiralaxisDB.Angle_Data.Valid = true

	spiralaxisDB.Length_Data.Float64 = spiralaxis.Length
	spiralaxisDB.Length_Data.Valid = true

	spiralaxisDB.CenterX_Data.Float64 = spiralaxis.CenterX
	spiralaxisDB.CenterX_Data.Valid = true

	spiralaxisDB.CenterY_Data.Float64 = spiralaxis.CenterY
	spiralaxisDB.CenterY_Data.Valid = true

	spiralaxisDB.Color_Data.String = spiralaxis.Color
	spiralaxisDB.Color_Data.Valid = true

	spiralaxisDB.FillOpacity_Data.Float64 = spiralaxis.FillOpacity
	spiralaxisDB.FillOpacity_Data.Valid = true

	spiralaxisDB.Stroke_Data.String = spiralaxis.Stroke
	spiralaxisDB.Stroke_Data.Valid = true

	spiralaxisDB.StrokeOpacity_Data.Float64 = spiralaxis.StrokeOpacity
	spiralaxisDB.StrokeOpacity_Data.Valid = true

	spiralaxisDB.StrokeWidth_Data.Float64 = spiralaxis.StrokeWidth
	spiralaxisDB.StrokeWidth_Data.Valid = true

	spiralaxisDB.StrokeDashArray_Data.String = spiralaxis.StrokeDashArray
	spiralaxisDB.StrokeDashArray_Data.Valid = true

	spiralaxisDB.StrokeDashArrayWhenSelected_Data.String = spiralaxis.StrokeDashArrayWhenSelected
	spiralaxisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralaxisDB.Transform_Data.String = spiralaxis.Transform
	spiralaxisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromSpiralAxis_WOP
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsFromSpiralAxis_WOP(spiralaxis *models.SpiralAxis_WOP) {
	// insertion point for fields commit

	spiralaxisDB.Name_Data.String = spiralaxis.Name
	spiralaxisDB.Name_Data.Valid = true

	spiralaxisDB.IsDisplayed_Data.Bool = spiralaxis.IsDisplayed
	spiralaxisDB.IsDisplayed_Data.Valid = true

	spiralaxisDB.Angle_Data.Float64 = spiralaxis.Angle
	spiralaxisDB.Angle_Data.Valid = true

	spiralaxisDB.Length_Data.Float64 = spiralaxis.Length
	spiralaxisDB.Length_Data.Valid = true

	spiralaxisDB.CenterX_Data.Float64 = spiralaxis.CenterX
	spiralaxisDB.CenterX_Data.Valid = true

	spiralaxisDB.CenterY_Data.Float64 = spiralaxis.CenterY
	spiralaxisDB.CenterY_Data.Valid = true

	spiralaxisDB.Color_Data.String = spiralaxis.Color
	spiralaxisDB.Color_Data.Valid = true

	spiralaxisDB.FillOpacity_Data.Float64 = spiralaxis.FillOpacity
	spiralaxisDB.FillOpacity_Data.Valid = true

	spiralaxisDB.Stroke_Data.String = spiralaxis.Stroke
	spiralaxisDB.Stroke_Data.Valid = true

	spiralaxisDB.StrokeOpacity_Data.Float64 = spiralaxis.StrokeOpacity
	spiralaxisDB.StrokeOpacity_Data.Valid = true

	spiralaxisDB.StrokeWidth_Data.Float64 = spiralaxis.StrokeWidth
	spiralaxisDB.StrokeWidth_Data.Valid = true

	spiralaxisDB.StrokeDashArray_Data.String = spiralaxis.StrokeDashArray
	spiralaxisDB.StrokeDashArray_Data.Valid = true

	spiralaxisDB.StrokeDashArrayWhenSelected_Data.String = spiralaxis.StrokeDashArrayWhenSelected
	spiralaxisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralaxisDB.Transform_Data.String = spiralaxis.Transform
	spiralaxisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromSpiralAxisWOP
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsFromSpiralAxisWOP(spiralaxis *SpiralAxisWOP) {
	// insertion point for fields commit

	spiralaxisDB.Name_Data.String = spiralaxis.Name
	spiralaxisDB.Name_Data.Valid = true

	spiralaxisDB.IsDisplayed_Data.Bool = spiralaxis.IsDisplayed
	spiralaxisDB.IsDisplayed_Data.Valid = true

	spiralaxisDB.Angle_Data.Float64 = spiralaxis.Angle
	spiralaxisDB.Angle_Data.Valid = true

	spiralaxisDB.Length_Data.Float64 = spiralaxis.Length
	spiralaxisDB.Length_Data.Valid = true

	spiralaxisDB.CenterX_Data.Float64 = spiralaxis.CenterX
	spiralaxisDB.CenterX_Data.Valid = true

	spiralaxisDB.CenterY_Data.Float64 = spiralaxis.CenterY
	spiralaxisDB.CenterY_Data.Valid = true

	spiralaxisDB.Color_Data.String = spiralaxis.Color
	spiralaxisDB.Color_Data.Valid = true

	spiralaxisDB.FillOpacity_Data.Float64 = spiralaxis.FillOpacity
	spiralaxisDB.FillOpacity_Data.Valid = true

	spiralaxisDB.Stroke_Data.String = spiralaxis.Stroke
	spiralaxisDB.Stroke_Data.Valid = true

	spiralaxisDB.StrokeOpacity_Data.Float64 = spiralaxis.StrokeOpacity
	spiralaxisDB.StrokeOpacity_Data.Valid = true

	spiralaxisDB.StrokeWidth_Data.Float64 = spiralaxis.StrokeWidth
	spiralaxisDB.StrokeWidth_Data.Valid = true

	spiralaxisDB.StrokeDashArray_Data.String = spiralaxis.StrokeDashArray
	spiralaxisDB.StrokeDashArray_Data.Valid = true

	spiralaxisDB.StrokeDashArrayWhenSelected_Data.String = spiralaxis.StrokeDashArrayWhenSelected
	spiralaxisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralaxisDB.Transform_Data.String = spiralaxis.Transform
	spiralaxisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToSpiralAxis
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsToSpiralAxis(spiralaxis *models.SpiralAxis) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralaxis.Name = spiralaxisDB.Name_Data.String
	spiralaxis.IsDisplayed = spiralaxisDB.IsDisplayed_Data.Bool
	spiralaxis.Angle = spiralaxisDB.Angle_Data.Float64
	spiralaxis.Length = spiralaxisDB.Length_Data.Float64
	spiralaxis.CenterX = spiralaxisDB.CenterX_Data.Float64
	spiralaxis.CenterY = spiralaxisDB.CenterY_Data.Float64
	spiralaxis.Color = spiralaxisDB.Color_Data.String
	spiralaxis.FillOpacity = spiralaxisDB.FillOpacity_Data.Float64
	spiralaxis.Stroke = spiralaxisDB.Stroke_Data.String
	spiralaxis.StrokeOpacity = spiralaxisDB.StrokeOpacity_Data.Float64
	spiralaxis.StrokeWidth = spiralaxisDB.StrokeWidth_Data.Float64
	spiralaxis.StrokeDashArray = spiralaxisDB.StrokeDashArray_Data.String
	spiralaxis.StrokeDashArrayWhenSelected = spiralaxisDB.StrokeDashArrayWhenSelected_Data.String
	spiralaxis.Transform = spiralaxisDB.Transform_Data.String
}

// CopyBasicFieldsToSpiralAxis_WOP
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsToSpiralAxis_WOP(spiralaxis *models.SpiralAxis_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralaxis.Name = spiralaxisDB.Name_Data.String
	spiralaxis.IsDisplayed = spiralaxisDB.IsDisplayed_Data.Bool
	spiralaxis.Angle = spiralaxisDB.Angle_Data.Float64
	spiralaxis.Length = spiralaxisDB.Length_Data.Float64
	spiralaxis.CenterX = spiralaxisDB.CenterX_Data.Float64
	spiralaxis.CenterY = spiralaxisDB.CenterY_Data.Float64
	spiralaxis.Color = spiralaxisDB.Color_Data.String
	spiralaxis.FillOpacity = spiralaxisDB.FillOpacity_Data.Float64
	spiralaxis.Stroke = spiralaxisDB.Stroke_Data.String
	spiralaxis.StrokeOpacity = spiralaxisDB.StrokeOpacity_Data.Float64
	spiralaxis.StrokeWidth = spiralaxisDB.StrokeWidth_Data.Float64
	spiralaxis.StrokeDashArray = spiralaxisDB.StrokeDashArray_Data.String
	spiralaxis.StrokeDashArrayWhenSelected = spiralaxisDB.StrokeDashArrayWhenSelected_Data.String
	spiralaxis.Transform = spiralaxisDB.Transform_Data.String
}

// CopyBasicFieldsToSpiralAxisWOP
func (spiralaxisDB *SpiralAxisDB) CopyBasicFieldsToSpiralAxisWOP(spiralaxis *SpiralAxisWOP) {
	spiralaxis.ID = int(spiralaxisDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spiralaxis.Name = spiralaxisDB.Name_Data.String
	spiralaxis.IsDisplayed = spiralaxisDB.IsDisplayed_Data.Bool
	spiralaxis.Angle = spiralaxisDB.Angle_Data.Float64
	spiralaxis.Length = spiralaxisDB.Length_Data.Float64
	spiralaxis.CenterX = spiralaxisDB.CenterX_Data.Float64
	spiralaxis.CenterY = spiralaxisDB.CenterY_Data.Float64
	spiralaxis.Color = spiralaxisDB.Color_Data.String
	spiralaxis.FillOpacity = spiralaxisDB.FillOpacity_Data.Float64
	spiralaxis.Stroke = spiralaxisDB.Stroke_Data.String
	spiralaxis.StrokeOpacity = spiralaxisDB.StrokeOpacity_Data.Float64
	spiralaxis.StrokeWidth = spiralaxisDB.StrokeWidth_Data.Float64
	spiralaxis.StrokeDashArray = spiralaxisDB.StrokeDashArray_Data.String
	spiralaxis.StrokeDashArrayWhenSelected = spiralaxisDB.StrokeDashArrayWhenSelected_Data.String
	spiralaxis.Transform = spiralaxisDB.Transform_Data.String
}

// Backup generates a json file from a slice of all SpiralAxisDB instances in the backrepo
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SpiralAxisDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralAxisDB, 0)
	for _, spiralaxisDB := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB {
		forBackup = append(forBackup, spiralaxisDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SpiralAxis ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SpiralAxis file", err.Error())
	}
}

// Backup generates a json file from a slice of all SpiralAxisDB instances in the backrepo
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralAxisDB, 0)
	for _, spiralaxisDB := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB {
		forBackup = append(forBackup, spiralaxisDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SpiralAxis")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SpiralAxis_Fields, -1)
	for _, spiralaxisDB := range forBackup {

		var spiralaxisWOP SpiralAxisWOP
		spiralaxisDB.CopyBasicFieldsToSpiralAxisWOP(&spiralaxisWOP)

		row := sh.AddRow()
		row.WriteStruct(&spiralaxisWOP, -1)
	}
}

// RestoreXL from the "SpiralAxis" sheet all SpiralAxisDB instances
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSpiralAxisid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SpiralAxis"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSpiralAxis.rowVisitorSpiralAxis)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) rowVisitorSpiralAxis(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spiralaxisWOP SpiralAxisWOP
		row.ReadStruct(&spiralaxisWOP)

		// add the unmarshalled struct to the stage
		spiralaxisDB := new(SpiralAxisDB)
		spiralaxisDB.CopyBasicFieldsFromSpiralAxisWOP(&spiralaxisWOP)

		spiralaxisDB_ID_atBackupTime := spiralaxisDB.ID
		spiralaxisDB.ID = 0
		query := backRepoSpiralAxis.db.Create(spiralaxisDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[spiralaxisDB.ID] = spiralaxisDB
		BackRepoSpiralAxisid_atBckpTime_newID[spiralaxisDB_ID_atBackupTime] = spiralaxisDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SpiralAxisDB.json" in dirPath that stores an array
// of SpiralAxisDB and stores it in the database
// the map BackRepoSpiralAxisid_atBckpTime_newID is updated accordingly
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSpiralAxisid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SpiralAxisDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SpiralAxis file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SpiralAxisDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SpiralAxisDBID_SpiralAxisDB
	for _, spiralaxisDB := range forRestore {

		spiralaxisDB_ID_atBackupTime := spiralaxisDB.ID
		spiralaxisDB.ID = 0
		query := backRepoSpiralAxis.db.Create(spiralaxisDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[spiralaxisDB.ID] = spiralaxisDB
		BackRepoSpiralAxisid_atBckpTime_newID[spiralaxisDB_ID_atBackupTime] = spiralaxisDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SpiralAxis file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SpiralAxis>id_atBckpTime_newID
// to compute new index
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) RestorePhaseTwo() {

	for _, spiralaxisDB := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB {

		// next line of code is to avert unused variable compilation error
		_ = spiralaxisDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if spiralaxisDB.ShapeCategoryID.Int64 != 0 {
			spiralaxisDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(spiralaxisDB.ShapeCategoryID.Int64)])
			spiralaxisDB.ShapeCategoryID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSpiralAxis.db.Model(spiralaxisDB).Updates(*spiralaxisDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSpiralAxis.ResetReversePointers commits all staged instances of SpiralAxis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralaxis := range backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisPtr {
		backRepoSpiralAxis.ResetReversePointersInstance(backRepo, idx, spiralaxis)
	}

	return
}

func (backRepoSpiralAxis *BackRepoSpiralAxisStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spiralaxis *models.SpiralAxis) (Error error) {

	// fetch matching spiralaxisDB
	if spiralaxisDB, ok := backRepoSpiralAxis.Map_SpiralAxisDBID_SpiralAxisDB[idx]; ok {
		_ = spiralaxisDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSpiralAxisid_atBckpTime_newID map[uint]uint