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

	"github.com/fullstack-lang/gongtone/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Freqency_sql sql.NullBool
var dummy_Freqency_time time.Duration
var dummy_Freqency_sort sort.Float64Slice

// FreqencyAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model freqencyAPI
type FreqencyAPI struct {
	gorm.Model

	models.Freqency_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	FreqencyPointersEncoding FreqencyPointersEncoding
}

// FreqencyPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FreqencyPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FreqencyDB describes a freqency in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model freqencyDB
type FreqencyDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field freqencyDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	FreqencyPointersEncoding
}

// FreqencyDBs arrays freqencyDBs
// swagger:response freqencyDBsResponse
type FreqencyDBs []FreqencyDB

// FreqencyDBResponse provides response
// swagger:response freqencyDBResponse
type FreqencyDBResponse struct {
	FreqencyDB
}

// FreqencyWOP is a Freqency without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FreqencyWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Freqency_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoFreqencyStruct struct {
	// stores FreqencyDB according to their gorm ID
	Map_FreqencyDBID_FreqencyDB map[uint]*FreqencyDB

	// stores FreqencyDB ID according to Freqency address
	Map_FreqencyPtr_FreqencyDBID map[*models.Freqency]uint

	// stores Freqency according to their gorm ID
	Map_FreqencyDBID_FreqencyPtr map[uint]*models.Freqency

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFreqency *BackRepoFreqencyStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFreqency.stage
	return
}

func (backRepoFreqency *BackRepoFreqencyStruct) GetDB() *gorm.DB {
	return backRepoFreqency.db
}

// GetFreqencyDBFromFreqencyPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFreqency *BackRepoFreqencyStruct) GetFreqencyDBFromFreqencyPtr(freqency *models.Freqency) (freqencyDB *FreqencyDB) {
	id := backRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]
	freqencyDB = backRepoFreqency.Map_FreqencyDBID_FreqencyDB[id]
	return
}

// BackRepoFreqency.CommitPhaseOne commits all staged instances of Freqency to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFreqency *BackRepoFreqencyStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for freqency := range stage.Freqencys {
		backRepoFreqency.CommitPhaseOneInstance(freqency)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, freqency := range backRepoFreqency.Map_FreqencyDBID_FreqencyPtr {
		if _, ok := stage.Freqencys[freqency]; !ok {
			backRepoFreqency.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFreqency.CommitDeleteInstance commits deletion of Freqency to the BackRepo
func (backRepoFreqency *BackRepoFreqencyStruct) CommitDeleteInstance(id uint) (Error error) {

	freqency := backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[id]

	// freqency is not staged anymore, remove freqencyDB
	freqencyDB := backRepoFreqency.Map_FreqencyDBID_FreqencyDB[id]
	query := backRepoFreqency.db.Unscoped().Delete(&freqencyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFreqency.Map_FreqencyPtr_FreqencyDBID, freqency)
	delete(backRepoFreqency.Map_FreqencyDBID_FreqencyPtr, id)
	delete(backRepoFreqency.Map_FreqencyDBID_FreqencyDB, id)

	return
}

// BackRepoFreqency.CommitPhaseOneInstance commits freqency staged instances of Freqency to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFreqency *BackRepoFreqencyStruct) CommitPhaseOneInstance(freqency *models.Freqency) (Error error) {

	// check if the freqency is not commited yet
	if _, ok := backRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]; ok {
		return
	}

	// initiate freqency
	var freqencyDB FreqencyDB
	freqencyDB.CopyBasicFieldsFromFreqency(freqency)

	query := backRepoFreqency.db.Create(&freqencyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency] = freqencyDB.ID
	backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID] = freqency
	backRepoFreqency.Map_FreqencyDBID_FreqencyDB[freqencyDB.ID] = &freqencyDB

	return
}

// BackRepoFreqency.CommitPhaseTwo commits all staged instances of Freqency to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFreqency *BackRepoFreqencyStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, freqency := range backRepoFreqency.Map_FreqencyDBID_FreqencyPtr {
		backRepoFreqency.CommitPhaseTwoInstance(backRepo, idx, freqency)
	}

	return
}

// BackRepoFreqency.CommitPhaseTwoInstance commits {{structname }} of models.Freqency to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFreqency *BackRepoFreqencyStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, freqency *models.Freqency) (Error error) {

	// fetch matching freqencyDB
	if freqencyDB, ok := backRepoFreqency.Map_FreqencyDBID_FreqencyDB[idx]; ok {

		freqencyDB.CopyBasicFieldsFromFreqency(freqency)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFreqency.db.Save(&freqencyDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Freqency intance %s", freqency.Name))
		return err
	}

	return
}

// BackRepoFreqency.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFreqency *BackRepoFreqencyStruct) CheckoutPhaseOne() (Error error) {

	freqencyDBArray := make([]FreqencyDB, 0)
	query := backRepoFreqency.db.Find(&freqencyDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	freqencyInstancesToBeRemovedFromTheStage := make(map[*models.Freqency]any)
	for key, value := range backRepoFreqency.stage.Freqencys {
		freqencyInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, freqencyDB := range freqencyDBArray {
		backRepoFreqency.CheckoutPhaseOneInstance(&freqencyDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		freqency, ok := backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
		if ok {
			delete(freqencyInstancesToBeRemovedFromTheStage, freqency)
		}
	}

	// remove from stage and back repo's 3 maps all freqencys that are not in the checkout
	for freqency := range freqencyInstancesToBeRemovedFromTheStage {
		freqency.Unstage(backRepoFreqency.GetStage())

		// remove instance from the back repo 3 maps
		freqencyID := backRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]
		delete(backRepoFreqency.Map_FreqencyPtr_FreqencyDBID, freqency)
		delete(backRepoFreqency.Map_FreqencyDBID_FreqencyDB, freqencyID)
		delete(backRepoFreqency.Map_FreqencyDBID_FreqencyPtr, freqencyID)
	}

	return
}

// CheckoutPhaseOneInstance takes a freqencyDB that has been found in the DB, updates the backRepo and stages the
// models version of the freqencyDB
func (backRepoFreqency *BackRepoFreqencyStruct) CheckoutPhaseOneInstance(freqencyDB *FreqencyDB) (Error error) {

	freqency, ok := backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]
	if !ok {
		freqency = new(models.Freqency)

		backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID] = freqency
		backRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency] = freqencyDB.ID

		// append model store with the new element
		freqency.Name = freqencyDB.Name_Data.String
		freqency.Stage(backRepoFreqency.GetStage())
	}
	freqencyDB.CopyBasicFieldsToFreqency(freqency)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	freqency.Stage(backRepoFreqency.GetStage())

	// preserve pointer to freqencyDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FreqencyDBID_FreqencyDB)[freqencyDB hold variable pointers
	freqencyDB_Data := *freqencyDB
	preservedPtrToFreqency := &freqencyDB_Data
	backRepoFreqency.Map_FreqencyDBID_FreqencyDB[freqencyDB.ID] = preservedPtrToFreqency

	return
}

// BackRepoFreqency.CheckoutPhaseTwo Checkouts all staged instances of Freqency to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFreqency *BackRepoFreqencyStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, freqencyDB := range backRepoFreqency.Map_FreqencyDBID_FreqencyDB {
		backRepoFreqency.CheckoutPhaseTwoInstance(backRepo, freqencyDB)
	}
	return
}

// BackRepoFreqency.CheckoutPhaseTwoInstance Checkouts staged instances of Freqency to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFreqency *BackRepoFreqencyStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, freqencyDB *FreqencyDB) (Error error) {

	freqency := backRepoFreqency.Map_FreqencyDBID_FreqencyPtr[freqencyDB.ID]

	freqencyDB.DecodePointers(backRepo, freqency)

	return
}

func (freqencyDB *FreqencyDB) DecodePointers(backRepo *BackRepoStruct, freqency *models.Freqency) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitFreqency allows commit of a single freqency (if already staged)
func (backRepo *BackRepoStruct) CommitFreqency(freqency *models.Freqency) {
	backRepo.BackRepoFreqency.CommitPhaseOneInstance(freqency)
	if id, ok := backRepo.BackRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]; ok {
		backRepo.BackRepoFreqency.CommitPhaseTwoInstance(backRepo, id, freqency)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFreqency allows checkout of a single freqency (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFreqency(freqency *models.Freqency) {
	// check if the freqency is staged
	if _, ok := backRepo.BackRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]; ok {

		if id, ok := backRepo.BackRepoFreqency.Map_FreqencyPtr_FreqencyDBID[freqency]; ok {
			var freqencyDB FreqencyDB
			freqencyDB.ID = id

			if err := backRepo.BackRepoFreqency.db.First(&freqencyDB, id).Error; err != nil {
				log.Fatalln("CheckoutFreqency : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFreqency.CheckoutPhaseOneInstance(&freqencyDB)
			backRepo.BackRepoFreqency.CheckoutPhaseTwoInstance(backRepo, &freqencyDB)
		}
	}
}

// CopyBasicFieldsFromFreqency
func (freqencyDB *FreqencyDB) CopyBasicFieldsFromFreqency(freqency *models.Freqency) {
	// insertion point for fields commit

	freqencyDB.Name_Data.String = freqency.Name
	freqencyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFreqency_WOP
func (freqencyDB *FreqencyDB) CopyBasicFieldsFromFreqency_WOP(freqency *models.Freqency_WOP) {
	// insertion point for fields commit

	freqencyDB.Name_Data.String = freqency.Name
	freqencyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFreqencyWOP
func (freqencyDB *FreqencyDB) CopyBasicFieldsFromFreqencyWOP(freqency *FreqencyWOP) {
	// insertion point for fields commit

	freqencyDB.Name_Data.String = freqency.Name
	freqencyDB.Name_Data.Valid = true
}

// CopyBasicFieldsToFreqency
func (freqencyDB *FreqencyDB) CopyBasicFieldsToFreqency(freqency *models.Freqency) {
	// insertion point for checkout of basic fields (back repo to stage)
	freqency.Name = freqencyDB.Name_Data.String
}

// CopyBasicFieldsToFreqency_WOP
func (freqencyDB *FreqencyDB) CopyBasicFieldsToFreqency_WOP(freqency *models.Freqency_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	freqency.Name = freqencyDB.Name_Data.String
}

// CopyBasicFieldsToFreqencyWOP
func (freqencyDB *FreqencyDB) CopyBasicFieldsToFreqencyWOP(freqency *FreqencyWOP) {
	freqency.ID = int(freqencyDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	freqency.Name = freqencyDB.Name_Data.String
}

// Backup generates a json file from a slice of all FreqencyDB instances in the backrepo
func (backRepoFreqency *BackRepoFreqencyStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FreqencyDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FreqencyDB, 0)
	for _, freqencyDB := range backRepoFreqency.Map_FreqencyDBID_FreqencyDB {
		forBackup = append(forBackup, freqencyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Freqency ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Freqency file", err.Error())
	}
}

// Backup generates a json file from a slice of all FreqencyDB instances in the backrepo
func (backRepoFreqency *BackRepoFreqencyStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FreqencyDB, 0)
	for _, freqencyDB := range backRepoFreqency.Map_FreqencyDBID_FreqencyDB {
		forBackup = append(forBackup, freqencyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Freqency")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Freqency_Fields, -1)
	for _, freqencyDB := range forBackup {

		var freqencyWOP FreqencyWOP
		freqencyDB.CopyBasicFieldsToFreqencyWOP(&freqencyWOP)

		row := sh.AddRow()
		row.WriteStruct(&freqencyWOP, -1)
	}
}

// RestoreXL from the "Freqency" sheet all FreqencyDB instances
func (backRepoFreqency *BackRepoFreqencyStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFreqencyid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Freqency"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFreqency.rowVisitorFreqency)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFreqency *BackRepoFreqencyStruct) rowVisitorFreqency(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var freqencyWOP FreqencyWOP
		row.ReadStruct(&freqencyWOP)

		// add the unmarshalled struct to the stage
		freqencyDB := new(FreqencyDB)
		freqencyDB.CopyBasicFieldsFromFreqencyWOP(&freqencyWOP)

		freqencyDB_ID_atBackupTime := freqencyDB.ID
		freqencyDB.ID = 0
		query := backRepoFreqency.db.Create(freqencyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFreqency.Map_FreqencyDBID_FreqencyDB[freqencyDB.ID] = freqencyDB
		BackRepoFreqencyid_atBckpTime_newID[freqencyDB_ID_atBackupTime] = freqencyDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FreqencyDB.json" in dirPath that stores an array
// of FreqencyDB and stores it in the database
// the map BackRepoFreqencyid_atBckpTime_newID is updated accordingly
func (backRepoFreqency *BackRepoFreqencyStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFreqencyid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FreqencyDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Freqency file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FreqencyDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FreqencyDBID_FreqencyDB
	for _, freqencyDB := range forRestore {

		freqencyDB_ID_atBackupTime := freqencyDB.ID
		freqencyDB.ID = 0
		query := backRepoFreqency.db.Create(freqencyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFreqency.Map_FreqencyDBID_FreqencyDB[freqencyDB.ID] = freqencyDB
		BackRepoFreqencyid_atBckpTime_newID[freqencyDB_ID_atBackupTime] = freqencyDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Freqency file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Freqency>id_atBckpTime_newID
// to compute new index
func (backRepoFreqency *BackRepoFreqencyStruct) RestorePhaseTwo() {

	for _, freqencyDB := range backRepoFreqency.Map_FreqencyDBID_FreqencyDB {

		// next line of code is to avert unused variable compilation error
		_ = freqencyDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFreqency.db.Model(freqencyDB).Updates(*freqencyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFreqency.ResetReversePointers commits all staged instances of Freqency to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFreqency *BackRepoFreqencyStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, freqency := range backRepoFreqency.Map_FreqencyDBID_FreqencyPtr {
		backRepoFreqency.ResetReversePointersInstance(backRepo, idx, freqency)
	}

	return
}

func (backRepoFreqency *BackRepoFreqencyStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, freqency *models.Freqency) (Error error) {

	// fetch matching freqencyDB
	if freqencyDB, ok := backRepoFreqency.Map_FreqencyDBID_FreqencyDB[idx]; ok {
		_ = freqencyDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFreqencyid_atBckpTime_newID map[uint]uint