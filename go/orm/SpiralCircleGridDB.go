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
var dummy_SpiralCircleGrid_sql sql.NullBool
var dummy_SpiralCircleGrid_time time.Duration
var dummy_SpiralCircleGrid_sort sort.Float64Slice

// SpiralCircleGridAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spiralcirclegridAPI
type SpiralCircleGridAPI struct {
	gorm.Model

	models.SpiralCircleGrid_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SpiralCircleGridPointersEncoding SpiralCircleGridPointersEncoding
}

// SpiralCircleGridPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SpiralCircleGridPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64

	// field SpiralRhombusGrid is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	SpiralRhombusGridID sql.NullInt64

	// field SpiralCircles is a slice of pointers to another Struct (optional or 0..1)
	SpiralCircles IntSlice `gorm:"type:TEXT"`
}

// SpiralCircleGridDB describes a spiralcirclegrid in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spiralcirclegridDB
type SpiralCircleGridDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spiralcirclegridDB.Name
	Name_Data sql.NullString

	// Declation for basic field spiralcirclegridDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SpiralCircleGridPointersEncoding
}

// SpiralCircleGridDBs arrays spiralcirclegridDBs
// swagger:response spiralcirclegridDBsResponse
type SpiralCircleGridDBs []SpiralCircleGridDB

// SpiralCircleGridDBResponse provides response
// swagger:response spiralcirclegridDBResponse
type SpiralCircleGridDBResponse struct {
	SpiralCircleGridDB
}

// SpiralCircleGridWOP is a SpiralCircleGrid without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SpiralCircleGridWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var SpiralCircleGrid_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
}

type BackRepoSpiralCircleGridStruct struct {
	// stores SpiralCircleGridDB according to their gorm ID
	Map_SpiralCircleGridDBID_SpiralCircleGridDB map[uint]*SpiralCircleGridDB

	// stores SpiralCircleGridDB ID according to SpiralCircleGrid address
	Map_SpiralCircleGridPtr_SpiralCircleGridDBID map[*models.SpiralCircleGrid]uint

	// stores SpiralCircleGrid according to their gorm ID
	Map_SpiralCircleGridDBID_SpiralCircleGridPtr map[uint]*models.SpiralCircleGrid

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSpiralCircleGrid.stage
	return
}

func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) GetDB() *gorm.DB {
	return backRepoSpiralCircleGrid.db
}

// GetSpiralCircleGridDBFromSpiralCircleGridPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) GetSpiralCircleGridDBFromSpiralCircleGridPtr(spiralcirclegrid *models.SpiralCircleGrid) (spiralcirclegridDB *SpiralCircleGridDB) {
	id := backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]
	spiralcirclegridDB = backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[id]
	return
}

// BackRepoSpiralCircleGrid.CommitPhaseOne commits all staged instances of SpiralCircleGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for spiralcirclegrid := range stage.SpiralCircleGrids {
		backRepoSpiralCircleGrid.CommitPhaseOneInstance(spiralcirclegrid)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spiralcirclegrid := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr {
		if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; !ok {
			backRepoSpiralCircleGrid.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSpiralCircleGrid.CommitDeleteInstance commits deletion of SpiralCircleGrid to the BackRepo
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CommitDeleteInstance(id uint) (Error error) {

	spiralcirclegrid := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[id]

	// spiralcirclegrid is not staged anymore, remove spiralcirclegridDB
	spiralcirclegridDB := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[id]
	query := backRepoSpiralCircleGrid.db.Unscoped().Delete(&spiralcirclegridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID, spiralcirclegrid)
	delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr, id)
	delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB, id)

	return
}

// BackRepoSpiralCircleGrid.CommitPhaseOneInstance commits spiralcirclegrid staged instances of SpiralCircleGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CommitPhaseOneInstance(spiralcirclegrid *models.SpiralCircleGrid) (Error error) {

	// check if the spiralcirclegrid is not commited yet
	if _, ok := backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]; ok {
		return
	}

	// initiate spiralcirclegrid
	var spiralcirclegridDB SpiralCircleGridDB
	spiralcirclegridDB.CopyBasicFieldsFromSpiralCircleGrid(spiralcirclegrid)

	query := backRepoSpiralCircleGrid.db.Create(&spiralcirclegridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid] = spiralcirclegridDB.ID
	backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID] = spiralcirclegrid
	backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[spiralcirclegridDB.ID] = &spiralcirclegridDB

	return
}

// BackRepoSpiralCircleGrid.CommitPhaseTwo commits all staged instances of SpiralCircleGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralcirclegrid := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr {
		backRepoSpiralCircleGrid.CommitPhaseTwoInstance(backRepo, idx, spiralcirclegrid)
	}

	return
}

// BackRepoSpiralCircleGrid.CommitPhaseTwoInstance commits {{structname }} of models.SpiralCircleGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spiralcirclegrid *models.SpiralCircleGrid) (Error error) {

	// fetch matching spiralcirclegridDB
	if spiralcirclegridDB, ok := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[idx]; ok {

		spiralcirclegridDB.CopyBasicFieldsFromSpiralCircleGrid(spiralcirclegrid)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value spiralcirclegrid.ShapeCategory translates to updating the spiralcirclegrid.ShapeCategoryID
		spiralcirclegridDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if spiralcirclegrid.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[spiralcirclegrid.ShapeCategory]; ok {
				spiralcirclegridDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				spiralcirclegridDB.ShapeCategoryID.Valid = true
			}
		} else {
			spiralcirclegridDB.ShapeCategoryID.Int64 = 0
			spiralcirclegridDB.ShapeCategoryID.Valid = true
		}

		// commit pointer value spiralcirclegrid.SpiralRhombusGrid translates to updating the spiralcirclegrid.SpiralRhombusGridID
		spiralcirclegridDB.SpiralRhombusGridID.Valid = true // allow for a 0 value (nil association)
		if spiralcirclegrid.SpiralRhombusGrid != nil {
			if SpiralRhombusGridId, ok := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralcirclegrid.SpiralRhombusGrid]; ok {
				spiralcirclegridDB.SpiralRhombusGridID.Int64 = int64(SpiralRhombusGridId)
				spiralcirclegridDB.SpiralRhombusGridID.Valid = true
			}
		} else {
			spiralcirclegridDB.SpiralRhombusGridID.Int64 = 0
			spiralcirclegridDB.SpiralRhombusGridID.Valid = true
		}

		// 1. reset
		spiralcirclegridDB.SpiralCircleGridPointersEncoding.SpiralCircles = make([]int, 0)
		// 2. encode
		for _, spiralcircleAssocEnd := range spiralcirclegrid.SpiralCircles {
			spiralcircleAssocEnd_DB :=
				backRepo.BackRepoSpiralCircle.GetSpiralCircleDBFromSpiralCirclePtr(spiralcircleAssocEnd)
			
			// the stage might be inconsistant, meaning that the spiralcircleAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if spiralcircleAssocEnd_DB == nil {
				continue
			}
			
			spiralcirclegridDB.SpiralCircleGridPointersEncoding.SpiralCircles =
				append(spiralcirclegridDB.SpiralCircleGridPointersEncoding.SpiralCircles, int(spiralcircleAssocEnd_DB.ID))
		}

		query := backRepoSpiralCircleGrid.db.Save(&spiralcirclegridDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SpiralCircleGrid intance %s", spiralcirclegrid.Name))
		return err
	}

	return
}

// BackRepoSpiralCircleGrid.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CheckoutPhaseOne() (Error error) {

	spiralcirclegridDBArray := make([]SpiralCircleGridDB, 0)
	query := backRepoSpiralCircleGrid.db.Find(&spiralcirclegridDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spiralcirclegridInstancesToBeRemovedFromTheStage := make(map[*models.SpiralCircleGrid]any)
	for key, value := range backRepoSpiralCircleGrid.stage.SpiralCircleGrids {
		spiralcirclegridInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spiralcirclegridDB := range spiralcirclegridDBArray {
		backRepoSpiralCircleGrid.CheckoutPhaseOneInstance(&spiralcirclegridDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spiralcirclegrid, ok := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]
		if ok {
			delete(spiralcirclegridInstancesToBeRemovedFromTheStage, spiralcirclegrid)
		}
	}

	// remove from stage and back repo's 3 maps all spiralcirclegrids that are not in the checkout
	for spiralcirclegrid := range spiralcirclegridInstancesToBeRemovedFromTheStage {
		spiralcirclegrid.Unstage(backRepoSpiralCircleGrid.GetStage())

		// remove instance from the back repo 3 maps
		spiralcirclegridID := backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]
		delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID, spiralcirclegrid)
		delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB, spiralcirclegridID)
		delete(backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr, spiralcirclegridID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spiralcirclegridDB that has been found in the DB, updates the backRepo and stages the
// models version of the spiralcirclegridDB
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CheckoutPhaseOneInstance(spiralcirclegridDB *SpiralCircleGridDB) (Error error) {

	spiralcirclegrid, ok := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]
	if !ok {
		spiralcirclegrid = new(models.SpiralCircleGrid)

		backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID] = spiralcirclegrid
		backRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid] = spiralcirclegridDB.ID

		// append model store with the new element
		spiralcirclegrid.Name = spiralcirclegridDB.Name_Data.String
		spiralcirclegrid.Stage(backRepoSpiralCircleGrid.GetStage())
	}
	spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid(spiralcirclegrid)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spiralcirclegrid.Stage(backRepoSpiralCircleGrid.GetStage())

	// preserve pointer to spiralcirclegridDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SpiralCircleGridDBID_SpiralCircleGridDB)[spiralcirclegridDB hold variable pointers
	spiralcirclegridDB_Data := *spiralcirclegridDB
	preservedPtrToSpiralCircleGrid := &spiralcirclegridDB_Data
	backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[spiralcirclegridDB.ID] = preservedPtrToSpiralCircleGrid

	return
}

// BackRepoSpiralCircleGrid.CheckoutPhaseTwo Checkouts all staged instances of SpiralCircleGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spiralcirclegridDB := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB {
		backRepoSpiralCircleGrid.CheckoutPhaseTwoInstance(backRepo, spiralcirclegridDB)
	}
	return
}

// BackRepoSpiralCircleGrid.CheckoutPhaseTwoInstance Checkouts staged instances of SpiralCircleGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spiralcirclegridDB *SpiralCircleGridDB) (Error error) {

	spiralcirclegrid := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr[spiralcirclegridDB.ID]

	spiralcirclegridDB.DecodePointers(backRepo, spiralcirclegrid)

	return
}

func (spiralcirclegridDB *SpiralCircleGridDB) DecodePointers(backRepo *BackRepoStruct, spiralcirclegrid *models.SpiralCircleGrid) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field
	spiralcirclegrid.ShapeCategory = nil
	if spiralcirclegridDB.ShapeCategoryID.Int64 != 0 {
		spiralcirclegrid.ShapeCategory = backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(spiralcirclegridDB.ShapeCategoryID.Int64)]
	}
	// SpiralRhombusGrid field
	spiralcirclegrid.SpiralRhombusGrid = nil
	if spiralcirclegridDB.SpiralRhombusGridID.Int64 != 0 {
		spiralcirclegrid.SpiralRhombusGrid = backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[uint(spiralcirclegridDB.SpiralRhombusGridID.Int64)]
	}
	// This loop redeem spiralcirclegrid.SpiralCircles in the stage from the encode in the back repo
	// It parses all SpiralCircleDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	spiralcirclegrid.SpiralCircles = spiralcirclegrid.SpiralCircles[:0]
	for _, _SpiralCircleid := range spiralcirclegridDB.SpiralCircleGridPointersEncoding.SpiralCircles {
		spiralcirclegrid.SpiralCircles = append(spiralcirclegrid.SpiralCircles, backRepo.BackRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[uint(_SpiralCircleid)])
	}

	return
}

// CommitSpiralCircleGrid allows commit of a single spiralcirclegrid (if already staged)
func (backRepo *BackRepoStruct) CommitSpiralCircleGrid(spiralcirclegrid *models.SpiralCircleGrid) {
	backRepo.BackRepoSpiralCircleGrid.CommitPhaseOneInstance(spiralcirclegrid)
	if id, ok := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]; ok {
		backRepo.BackRepoSpiralCircleGrid.CommitPhaseTwoInstance(backRepo, id, spiralcirclegrid)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSpiralCircleGrid allows checkout of a single spiralcirclegrid (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSpiralCircleGrid(spiralcirclegrid *models.SpiralCircleGrid) {
	// check if the spiralcirclegrid is staged
	if _, ok := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]; ok {

		if id, ok := backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridPtr_SpiralCircleGridDBID[spiralcirclegrid]; ok {
			var spiralcirclegridDB SpiralCircleGridDB
			spiralcirclegridDB.ID = id

			if err := backRepo.BackRepoSpiralCircleGrid.db.First(&spiralcirclegridDB, id).Error; err != nil {
				log.Fatalln("CheckoutSpiralCircleGrid : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSpiralCircleGrid.CheckoutPhaseOneInstance(&spiralcirclegridDB)
			backRepo.BackRepoSpiralCircleGrid.CheckoutPhaseTwoInstance(backRepo, &spiralcirclegridDB)
		}
	}
}

// CopyBasicFieldsFromSpiralCircleGrid
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsFromSpiralCircleGrid(spiralcirclegrid *models.SpiralCircleGrid) {
	// insertion point for fields commit

	spiralcirclegridDB.Name_Data.String = spiralcirclegrid.Name
	spiralcirclegridDB.Name_Data.Valid = true

	spiralcirclegridDB.IsDisplayed_Data.Bool = spiralcirclegrid.IsDisplayed
	spiralcirclegridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsFromSpiralCircleGrid_WOP
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsFromSpiralCircleGrid_WOP(spiralcirclegrid *models.SpiralCircleGrid_WOP) {
	// insertion point for fields commit

	spiralcirclegridDB.Name_Data.String = spiralcirclegrid.Name
	spiralcirclegridDB.Name_Data.Valid = true

	spiralcirclegridDB.IsDisplayed_Data.Bool = spiralcirclegrid.IsDisplayed
	spiralcirclegridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsFromSpiralCircleGridWOP
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsFromSpiralCircleGridWOP(spiralcirclegrid *SpiralCircleGridWOP) {
	// insertion point for fields commit

	spiralcirclegridDB.Name_Data.String = spiralcirclegrid.Name
	spiralcirclegridDB.Name_Data.Valid = true

	spiralcirclegridDB.IsDisplayed_Data.Bool = spiralcirclegrid.IsDisplayed
	spiralcirclegridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsToSpiralCircleGrid
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsToSpiralCircleGrid(spiralcirclegrid *models.SpiralCircleGrid) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcirclegrid.Name = spiralcirclegridDB.Name_Data.String
	spiralcirclegrid.IsDisplayed = spiralcirclegridDB.IsDisplayed_Data.Bool
}

// CopyBasicFieldsToSpiralCircleGrid_WOP
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsToSpiralCircleGrid_WOP(spiralcirclegrid *models.SpiralCircleGrid_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcirclegrid.Name = spiralcirclegridDB.Name_Data.String
	spiralcirclegrid.IsDisplayed = spiralcirclegridDB.IsDisplayed_Data.Bool
}

// CopyBasicFieldsToSpiralCircleGridWOP
func (spiralcirclegridDB *SpiralCircleGridDB) CopyBasicFieldsToSpiralCircleGridWOP(spiralcirclegrid *SpiralCircleGridWOP) {
	spiralcirclegrid.ID = int(spiralcirclegridDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcirclegrid.Name = spiralcirclegridDB.Name_Data.String
	spiralcirclegrid.IsDisplayed = spiralcirclegridDB.IsDisplayed_Data.Bool
}

// Backup generates a json file from a slice of all SpiralCircleGridDB instances in the backrepo
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SpiralCircleGridDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralCircleGridDB, 0)
	for _, spiralcirclegridDB := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB {
		forBackup = append(forBackup, spiralcirclegridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SpiralCircleGrid ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SpiralCircleGrid file", err.Error())
	}
}

// Backup generates a json file from a slice of all SpiralCircleGridDB instances in the backrepo
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralCircleGridDB, 0)
	for _, spiralcirclegridDB := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB {
		forBackup = append(forBackup, spiralcirclegridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SpiralCircleGrid")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SpiralCircleGrid_Fields, -1)
	for _, spiralcirclegridDB := range forBackup {

		var spiralcirclegridWOP SpiralCircleGridWOP
		spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGridWOP(&spiralcirclegridWOP)

		row := sh.AddRow()
		row.WriteStruct(&spiralcirclegridWOP, -1)
	}
}

// RestoreXL from the "SpiralCircleGrid" sheet all SpiralCircleGridDB instances
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSpiralCircleGridid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SpiralCircleGrid"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSpiralCircleGrid.rowVisitorSpiralCircleGrid)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) rowVisitorSpiralCircleGrid(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spiralcirclegridWOP SpiralCircleGridWOP
		row.ReadStruct(&spiralcirclegridWOP)

		// add the unmarshalled struct to the stage
		spiralcirclegridDB := new(SpiralCircleGridDB)
		spiralcirclegridDB.CopyBasicFieldsFromSpiralCircleGridWOP(&spiralcirclegridWOP)

		spiralcirclegridDB_ID_atBackupTime := spiralcirclegridDB.ID
		spiralcirclegridDB.ID = 0
		query := backRepoSpiralCircleGrid.db.Create(spiralcirclegridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[spiralcirclegridDB.ID] = spiralcirclegridDB
		BackRepoSpiralCircleGridid_atBckpTime_newID[spiralcirclegridDB_ID_atBackupTime] = spiralcirclegridDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SpiralCircleGridDB.json" in dirPath that stores an array
// of SpiralCircleGridDB and stores it in the database
// the map BackRepoSpiralCircleGridid_atBckpTime_newID is updated accordingly
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSpiralCircleGridid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SpiralCircleGridDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SpiralCircleGrid file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SpiralCircleGridDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SpiralCircleGridDBID_SpiralCircleGridDB
	for _, spiralcirclegridDB := range forRestore {

		spiralcirclegridDB_ID_atBackupTime := spiralcirclegridDB.ID
		spiralcirclegridDB.ID = 0
		query := backRepoSpiralCircleGrid.db.Create(spiralcirclegridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[spiralcirclegridDB.ID] = spiralcirclegridDB
		BackRepoSpiralCircleGridid_atBckpTime_newID[spiralcirclegridDB_ID_atBackupTime] = spiralcirclegridDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SpiralCircleGrid file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SpiralCircleGrid>id_atBckpTime_newID
// to compute new index
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) RestorePhaseTwo() {

	for _, spiralcirclegridDB := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB {

		// next line of code is to avert unused variable compilation error
		_ = spiralcirclegridDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if spiralcirclegridDB.ShapeCategoryID.Int64 != 0 {
			spiralcirclegridDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(spiralcirclegridDB.ShapeCategoryID.Int64)])
			spiralcirclegridDB.ShapeCategoryID.Valid = true
		}

		// reindexing SpiralRhombusGrid field
		if spiralcirclegridDB.SpiralRhombusGridID.Int64 != 0 {
			spiralcirclegridDB.SpiralRhombusGridID.Int64 = int64(BackRepoSpiralRhombusGridid_atBckpTime_newID[uint(spiralcirclegridDB.SpiralRhombusGridID.Int64)])
			spiralcirclegridDB.SpiralRhombusGridID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSpiralCircleGrid.db.Model(spiralcirclegridDB).Updates(*spiralcirclegridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSpiralCircleGrid.ResetReversePointers commits all staged instances of SpiralCircleGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralcirclegrid := range backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridPtr {
		backRepoSpiralCircleGrid.ResetReversePointersInstance(backRepo, idx, spiralcirclegrid)
	}

	return
}

func (backRepoSpiralCircleGrid *BackRepoSpiralCircleGridStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spiralcirclegrid *models.SpiralCircleGrid) (Error error) {

	// fetch matching spiralcirclegridDB
	if spiralcirclegridDB, ok := backRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB[idx]; ok {
		_ = spiralcirclegridDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSpiralCircleGridid_atBckpTime_newID map[uint]uint