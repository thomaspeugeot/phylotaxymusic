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
var dummy_RhombusGrid_sql sql.NullBool
var dummy_RhombusGrid_time time.Duration
var dummy_RhombusGrid_sort sort.Float64Slice

// RhombusGridAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model rhombusgridAPI
type RhombusGridAPI struct {
	gorm.Model

	models.RhombusGrid_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	RhombusGridPointersEncoding RhombusGridPointersEncoding
}

// RhombusGridPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type RhombusGridPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Reference is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ReferenceID sql.NullInt64
}

// RhombusGridDB describes a rhombusgrid in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model rhombusgridDB
type RhombusGridDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field rhombusgridDB.Name
	Name_Data sql.NullString

	// Declation for basic field rhombusgridDB.N
	N_Data sql.NullInt64

	// Declation for basic field rhombusgridDB.M
	M_Data sql.NullInt64
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	RhombusGridPointersEncoding
}

// RhombusGridDBs arrays rhombusgridDBs
// swagger:response rhombusgridDBsResponse
type RhombusGridDBs []RhombusGridDB

// RhombusGridDBResponse provides response
// swagger:response rhombusgridDBResponse
type RhombusGridDBResponse struct {
	RhombusGridDB
}

// RhombusGridWOP is a RhombusGrid without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type RhombusGridWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	N int `xlsx:"2"`

	M int `xlsx:"3"`
	// insertion for WOP pointer fields
}

var RhombusGrid_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"N",
	"M",
}

type BackRepoRhombusGridStruct struct {
	// stores RhombusGridDB according to their gorm ID
	Map_RhombusGridDBID_RhombusGridDB map[uint]*RhombusGridDB

	// stores RhombusGridDB ID according to RhombusGrid address
	Map_RhombusGridPtr_RhombusGridDBID map[*models.RhombusGrid]uint

	// stores RhombusGrid according to their gorm ID
	Map_RhombusGridDBID_RhombusGridPtr map[uint]*models.RhombusGrid

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoRhombusGrid *BackRepoRhombusGridStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoRhombusGrid.stage
	return
}

func (backRepoRhombusGrid *BackRepoRhombusGridStruct) GetDB() *gorm.DB {
	return backRepoRhombusGrid.db
}

// GetRhombusGridDBFromRhombusGridPtr is a handy function to access the back repo instance from the stage instance
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) GetRhombusGridDBFromRhombusGridPtr(rhombusgrid *models.RhombusGrid) (rhombusgridDB *RhombusGridDB) {
	id := backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]
	rhombusgridDB = backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[id]
	return
}

// BackRepoRhombusGrid.CommitPhaseOne commits all staged instances of RhombusGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for rhombusgrid := range stage.RhombusGrids {
		backRepoRhombusGrid.CommitPhaseOneInstance(rhombusgrid)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, rhombusgrid := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr {
		if _, ok := stage.RhombusGrids[rhombusgrid]; !ok {
			backRepoRhombusGrid.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoRhombusGrid.CommitDeleteInstance commits deletion of RhombusGrid to the BackRepo
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CommitDeleteInstance(id uint) (Error error) {

	rhombusgrid := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[id]

	// rhombusgrid is not staged anymore, remove rhombusgridDB
	rhombusgridDB := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[id]
	query := backRepoRhombusGrid.db.Unscoped().Delete(&rhombusgridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID, rhombusgrid)
	delete(backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr, id)
	delete(backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB, id)

	return
}

// BackRepoRhombusGrid.CommitPhaseOneInstance commits rhombusgrid staged instances of RhombusGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CommitPhaseOneInstance(rhombusgrid *models.RhombusGrid) (Error error) {

	// check if the rhombusgrid is not commited yet
	if _, ok := backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]; ok {
		return
	}

	// initiate rhombusgrid
	var rhombusgridDB RhombusGridDB
	rhombusgridDB.CopyBasicFieldsFromRhombusGrid(rhombusgrid)

	query := backRepoRhombusGrid.db.Create(&rhombusgridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid] = rhombusgridDB.ID
	backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID] = rhombusgrid
	backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[rhombusgridDB.ID] = &rhombusgridDB

	return
}

// BackRepoRhombusGrid.CommitPhaseTwo commits all staged instances of RhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, rhombusgrid := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr {
		backRepoRhombusGrid.CommitPhaseTwoInstance(backRepo, idx, rhombusgrid)
	}

	return
}

// BackRepoRhombusGrid.CommitPhaseTwoInstance commits {{structname }} of models.RhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, rhombusgrid *models.RhombusGrid) (Error error) {

	// fetch matching rhombusgridDB
	if rhombusgridDB, ok := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[idx]; ok {

		rhombusgridDB.CopyBasicFieldsFromRhombusGrid(rhombusgrid)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value rhombusgrid.Reference translates to updating the rhombusgrid.ReferenceID
		rhombusgridDB.ReferenceID.Valid = true // allow for a 0 value (nil association)
		if rhombusgrid.Reference != nil {
			if ReferenceId, ok := backRepo.BackRepoRhombus.Map_RhombusPtr_RhombusDBID[rhombusgrid.Reference]; ok {
				rhombusgridDB.ReferenceID.Int64 = int64(ReferenceId)
				rhombusgridDB.ReferenceID.Valid = true
			}
		} else {
			rhombusgridDB.ReferenceID.Int64 = 0
			rhombusgridDB.ReferenceID.Valid = true
		}

		query := backRepoRhombusGrid.db.Save(&rhombusgridDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown RhombusGrid intance %s", rhombusgrid.Name))
		return err
	}

	return
}

// BackRepoRhombusGrid.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CheckoutPhaseOne() (Error error) {

	rhombusgridDBArray := make([]RhombusGridDB, 0)
	query := backRepoRhombusGrid.db.Find(&rhombusgridDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	rhombusgridInstancesToBeRemovedFromTheStage := make(map[*models.RhombusGrid]any)
	for key, value := range backRepoRhombusGrid.stage.RhombusGrids {
		rhombusgridInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, rhombusgridDB := range rhombusgridDBArray {
		backRepoRhombusGrid.CheckoutPhaseOneInstance(&rhombusgridDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		rhombusgrid, ok := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]
		if ok {
			delete(rhombusgridInstancesToBeRemovedFromTheStage, rhombusgrid)
		}
	}

	// remove from stage and back repo's 3 maps all rhombusgrids that are not in the checkout
	for rhombusgrid := range rhombusgridInstancesToBeRemovedFromTheStage {
		rhombusgrid.Unstage(backRepoRhombusGrid.GetStage())

		// remove instance from the back repo 3 maps
		rhombusgridID := backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]
		delete(backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID, rhombusgrid)
		delete(backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB, rhombusgridID)
		delete(backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr, rhombusgridID)
	}

	return
}

// CheckoutPhaseOneInstance takes a rhombusgridDB that has been found in the DB, updates the backRepo and stages the
// models version of the rhombusgridDB
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CheckoutPhaseOneInstance(rhombusgridDB *RhombusGridDB) (Error error) {

	rhombusgrid, ok := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]
	if !ok {
		rhombusgrid = new(models.RhombusGrid)

		backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID] = rhombusgrid
		backRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid] = rhombusgridDB.ID

		// append model store with the new element
		rhombusgrid.Name = rhombusgridDB.Name_Data.String
		rhombusgrid.Stage(backRepoRhombusGrid.GetStage())
	}
	rhombusgridDB.CopyBasicFieldsToRhombusGrid(rhombusgrid)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	rhombusgrid.Stage(backRepoRhombusGrid.GetStage())

	// preserve pointer to rhombusgridDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_RhombusGridDBID_RhombusGridDB)[rhombusgridDB hold variable pointers
	rhombusgridDB_Data := *rhombusgridDB
	preservedPtrToRhombusGrid := &rhombusgridDB_Data
	backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[rhombusgridDB.ID] = preservedPtrToRhombusGrid

	return
}

// BackRepoRhombusGrid.CheckoutPhaseTwo Checkouts all staged instances of RhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, rhombusgridDB := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB {
		backRepoRhombusGrid.CheckoutPhaseTwoInstance(backRepo, rhombusgridDB)
	}
	return
}

// BackRepoRhombusGrid.CheckoutPhaseTwoInstance Checkouts staged instances of RhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, rhombusgridDB *RhombusGridDB) (Error error) {

	rhombusgrid := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[rhombusgridDB.ID]

	rhombusgridDB.DecodePointers(backRepo, rhombusgrid)

	return
}

func (rhombusgridDB *RhombusGridDB) DecodePointers(backRepo *BackRepoStruct, rhombusgrid *models.RhombusGrid) {

	// insertion point for checkout of pointer encoding
	// Reference field
	rhombusgrid.Reference = nil
	if rhombusgridDB.ReferenceID.Int64 != 0 {
		rhombusgrid.Reference = backRepo.BackRepoRhombus.Map_RhombusDBID_RhombusPtr[uint(rhombusgridDB.ReferenceID.Int64)]
	}
	return
}

// CommitRhombusGrid allows commit of a single rhombusgrid (if already staged)
func (backRepo *BackRepoStruct) CommitRhombusGrid(rhombusgrid *models.RhombusGrid) {
	backRepo.BackRepoRhombusGrid.CommitPhaseOneInstance(rhombusgrid)
	if id, ok := backRepo.BackRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]; ok {
		backRepo.BackRepoRhombusGrid.CommitPhaseTwoInstance(backRepo, id, rhombusgrid)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitRhombusGrid allows checkout of a single rhombusgrid (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutRhombusGrid(rhombusgrid *models.RhombusGrid) {
	// check if the rhombusgrid is staged
	if _, ok := backRepo.BackRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]; ok {

		if id, ok := backRepo.BackRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[rhombusgrid]; ok {
			var rhombusgridDB RhombusGridDB
			rhombusgridDB.ID = id

			if err := backRepo.BackRepoRhombusGrid.db.First(&rhombusgridDB, id).Error; err != nil {
				log.Fatalln("CheckoutRhombusGrid : Problem with getting object with id:", id)
			}
			backRepo.BackRepoRhombusGrid.CheckoutPhaseOneInstance(&rhombusgridDB)
			backRepo.BackRepoRhombusGrid.CheckoutPhaseTwoInstance(backRepo, &rhombusgridDB)
		}
	}
}

// CopyBasicFieldsFromRhombusGrid
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsFromRhombusGrid(rhombusgrid *models.RhombusGrid) {
	// insertion point for fields commit

	rhombusgridDB.Name_Data.String = rhombusgrid.Name
	rhombusgridDB.Name_Data.Valid = true

	rhombusgridDB.N_Data.Int64 = int64(rhombusgrid.N)
	rhombusgridDB.N_Data.Valid = true

	rhombusgridDB.M_Data.Int64 = int64(rhombusgrid.M)
	rhombusgridDB.M_Data.Valid = true
}

// CopyBasicFieldsFromRhombusGrid_WOP
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsFromRhombusGrid_WOP(rhombusgrid *models.RhombusGrid_WOP) {
	// insertion point for fields commit

	rhombusgridDB.Name_Data.String = rhombusgrid.Name
	rhombusgridDB.Name_Data.Valid = true

	rhombusgridDB.N_Data.Int64 = int64(rhombusgrid.N)
	rhombusgridDB.N_Data.Valid = true

	rhombusgridDB.M_Data.Int64 = int64(rhombusgrid.M)
	rhombusgridDB.M_Data.Valid = true
}

// CopyBasicFieldsFromRhombusGridWOP
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsFromRhombusGridWOP(rhombusgrid *RhombusGridWOP) {
	// insertion point for fields commit

	rhombusgridDB.Name_Data.String = rhombusgrid.Name
	rhombusgridDB.Name_Data.Valid = true

	rhombusgridDB.N_Data.Int64 = int64(rhombusgrid.N)
	rhombusgridDB.N_Data.Valid = true

	rhombusgridDB.M_Data.Int64 = int64(rhombusgrid.M)
	rhombusgridDB.M_Data.Valid = true
}

// CopyBasicFieldsToRhombusGrid
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsToRhombusGrid(rhombusgrid *models.RhombusGrid) {
	// insertion point for checkout of basic fields (back repo to stage)
	rhombusgrid.Name = rhombusgridDB.Name_Data.String
	rhombusgrid.N = int(rhombusgridDB.N_Data.Int64)
	rhombusgrid.M = int(rhombusgridDB.M_Data.Int64)
}

// CopyBasicFieldsToRhombusGrid_WOP
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsToRhombusGrid_WOP(rhombusgrid *models.RhombusGrid_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	rhombusgrid.Name = rhombusgridDB.Name_Data.String
	rhombusgrid.N = int(rhombusgridDB.N_Data.Int64)
	rhombusgrid.M = int(rhombusgridDB.M_Data.Int64)
}

// CopyBasicFieldsToRhombusGridWOP
func (rhombusgridDB *RhombusGridDB) CopyBasicFieldsToRhombusGridWOP(rhombusgrid *RhombusGridWOP) {
	rhombusgrid.ID = int(rhombusgridDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	rhombusgrid.Name = rhombusgridDB.Name_Data.String
	rhombusgrid.N = int(rhombusgridDB.N_Data.Int64)
	rhombusgrid.M = int(rhombusgridDB.M_Data.Int64)
}

// Backup generates a json file from a slice of all RhombusGridDB instances in the backrepo
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "RhombusGridDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RhombusGridDB, 0)
	for _, rhombusgridDB := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB {
		forBackup = append(forBackup, rhombusgridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json RhombusGrid ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json RhombusGrid file", err.Error())
	}
}

// Backup generates a json file from a slice of all RhombusGridDB instances in the backrepo
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*RhombusGridDB, 0)
	for _, rhombusgridDB := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB {
		forBackup = append(forBackup, rhombusgridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("RhombusGrid")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&RhombusGrid_Fields, -1)
	for _, rhombusgridDB := range forBackup {

		var rhombusgridWOP RhombusGridWOP
		rhombusgridDB.CopyBasicFieldsToRhombusGridWOP(&rhombusgridWOP)

		row := sh.AddRow()
		row.WriteStruct(&rhombusgridWOP, -1)
	}
}

// RestoreXL from the "RhombusGrid" sheet all RhombusGridDB instances
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoRhombusGridid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["RhombusGrid"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoRhombusGrid.rowVisitorRhombusGrid)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoRhombusGrid *BackRepoRhombusGridStruct) rowVisitorRhombusGrid(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var rhombusgridWOP RhombusGridWOP
		row.ReadStruct(&rhombusgridWOP)

		// add the unmarshalled struct to the stage
		rhombusgridDB := new(RhombusGridDB)
		rhombusgridDB.CopyBasicFieldsFromRhombusGridWOP(&rhombusgridWOP)

		rhombusgridDB_ID_atBackupTime := rhombusgridDB.ID
		rhombusgridDB.ID = 0
		query := backRepoRhombusGrid.db.Create(rhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[rhombusgridDB.ID] = rhombusgridDB
		BackRepoRhombusGridid_atBckpTime_newID[rhombusgridDB_ID_atBackupTime] = rhombusgridDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "RhombusGridDB.json" in dirPath that stores an array
// of RhombusGridDB and stores it in the database
// the map BackRepoRhombusGridid_atBckpTime_newID is updated accordingly
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoRhombusGridid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "RhombusGridDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json RhombusGrid file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*RhombusGridDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_RhombusGridDBID_RhombusGridDB
	for _, rhombusgridDB := range forRestore {

		rhombusgridDB_ID_atBackupTime := rhombusgridDB.ID
		rhombusgridDB.ID = 0
		query := backRepoRhombusGrid.db.Create(rhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[rhombusgridDB.ID] = rhombusgridDB
		BackRepoRhombusGridid_atBckpTime_newID[rhombusgridDB_ID_atBackupTime] = rhombusgridDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json RhombusGrid file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<RhombusGrid>id_atBckpTime_newID
// to compute new index
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) RestorePhaseTwo() {

	for _, rhombusgridDB := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB {

		// next line of code is to avert unused variable compilation error
		_ = rhombusgridDB

		// insertion point for reindexing pointers encoding
		// reindexing Reference field
		if rhombusgridDB.ReferenceID.Int64 != 0 {
			rhombusgridDB.ReferenceID.Int64 = int64(BackRepoRhombusid_atBckpTime_newID[uint(rhombusgridDB.ReferenceID.Int64)])
			rhombusgridDB.ReferenceID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoRhombusGrid.db.Model(rhombusgridDB).Updates(*rhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoRhombusGrid.ResetReversePointers commits all staged instances of RhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoRhombusGrid *BackRepoRhombusGridStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, rhombusgrid := range backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr {
		backRepoRhombusGrid.ResetReversePointersInstance(backRepo, idx, rhombusgrid)
	}

	return
}

func (backRepoRhombusGrid *BackRepoRhombusGridStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, rhombusgrid *models.RhombusGrid) (Error error) {

	// fetch matching rhombusgridDB
	if rhombusgridDB, ok := backRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB[idx]; ok {
		_ = rhombusgridDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoRhombusGridid_atBckpTime_newID map[uint]uint