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
var dummy_NoteInfo_sql sql.NullBool
var dummy_NoteInfo_time time.Duration
var dummy_NoteInfo_sort sort.Float64Slice

// NoteInfoAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model noteinfoAPI
type NoteInfoAPI struct {
	gorm.Model

	models.NoteInfo_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	NoteInfoPointersEncoding NoteInfoPointersEncoding
}

// NoteInfoPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type NoteInfoPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// NoteInfoDB describes a noteinfo in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model noteinfoDB
type NoteInfoDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field noteinfoDB.Name
	Name_Data sql.NullString

	// Declation for basic field noteinfoDB.IsKept
	// provide the sql storage for the boolan
	IsKept_Data sql.NullBool
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	NoteInfoPointersEncoding
}

// NoteInfoDBs arrays noteinfoDBs
// swagger:response noteinfoDBsResponse
type NoteInfoDBs []NoteInfoDB

// NoteInfoDBResponse provides response
// swagger:response noteinfoDBResponse
type NoteInfoDBResponse struct {
	NoteInfoDB
}

// NoteInfoWOP is a NoteInfo without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type NoteInfoWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsKept bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var NoteInfo_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsKept",
}

type BackRepoNoteInfoStruct struct {
	// stores NoteInfoDB according to their gorm ID
	Map_NoteInfoDBID_NoteInfoDB map[uint]*NoteInfoDB

	// stores NoteInfoDB ID according to NoteInfo address
	Map_NoteInfoPtr_NoteInfoDBID map[*models.NoteInfo]uint

	// stores NoteInfo according to their gorm ID
	Map_NoteInfoDBID_NoteInfoPtr map[uint]*models.NoteInfo

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoNoteInfo *BackRepoNoteInfoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoNoteInfo.stage
	return
}

func (backRepoNoteInfo *BackRepoNoteInfoStruct) GetDB() *gorm.DB {
	return backRepoNoteInfo.db
}

// GetNoteInfoDBFromNoteInfoPtr is a handy function to access the back repo instance from the stage instance
func (backRepoNoteInfo *BackRepoNoteInfoStruct) GetNoteInfoDBFromNoteInfoPtr(noteinfo *models.NoteInfo) (noteinfoDB *NoteInfoDB) {
	id := backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]
	noteinfoDB = backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[id]
	return
}

// BackRepoNoteInfo.CommitPhaseOne commits all staged instances of NoteInfo to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for noteinfo := range stage.NoteInfos {
		backRepoNoteInfo.CommitPhaseOneInstance(noteinfo)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, noteinfo := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr {
		if _, ok := stage.NoteInfos[noteinfo]; !ok {
			backRepoNoteInfo.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoNoteInfo.CommitDeleteInstance commits deletion of NoteInfo to the BackRepo
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CommitDeleteInstance(id uint) (Error error) {

	noteinfo := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[id]

	// noteinfo is not staged anymore, remove noteinfoDB
	noteinfoDB := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[id]
	query := backRepoNoteInfo.db.Unscoped().Delete(&noteinfoDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID, noteinfo)
	delete(backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr, id)
	delete(backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB, id)

	return
}

// BackRepoNoteInfo.CommitPhaseOneInstance commits noteinfo staged instances of NoteInfo to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CommitPhaseOneInstance(noteinfo *models.NoteInfo) (Error error) {

	// check if the noteinfo is not commited yet
	if _, ok := backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]; ok {
		return
	}

	// initiate noteinfo
	var noteinfoDB NoteInfoDB
	noteinfoDB.CopyBasicFieldsFromNoteInfo(noteinfo)

	query := backRepoNoteInfo.db.Create(&noteinfoDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo] = noteinfoDB.ID
	backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID] = noteinfo
	backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[noteinfoDB.ID] = &noteinfoDB

	return
}

// BackRepoNoteInfo.CommitPhaseTwo commits all staged instances of NoteInfo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, noteinfo := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr {
		backRepoNoteInfo.CommitPhaseTwoInstance(backRepo, idx, noteinfo)
	}

	return
}

// BackRepoNoteInfo.CommitPhaseTwoInstance commits {{structname }} of models.NoteInfo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, noteinfo *models.NoteInfo) (Error error) {

	// fetch matching noteinfoDB
	if noteinfoDB, ok := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[idx]; ok {

		noteinfoDB.CopyBasicFieldsFromNoteInfo(noteinfo)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoNoteInfo.db.Save(&noteinfoDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown NoteInfo intance %s", noteinfo.Name))
		return err
	}

	return
}

// BackRepoNoteInfo.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CheckoutPhaseOne() (Error error) {

	noteinfoDBArray := make([]NoteInfoDB, 0)
	query := backRepoNoteInfo.db.Find(&noteinfoDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	noteinfoInstancesToBeRemovedFromTheStage := make(map[*models.NoteInfo]any)
	for key, value := range backRepoNoteInfo.stage.NoteInfos {
		noteinfoInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, noteinfoDB := range noteinfoDBArray {
		backRepoNoteInfo.CheckoutPhaseOneInstance(&noteinfoDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		noteinfo, ok := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]
		if ok {
			delete(noteinfoInstancesToBeRemovedFromTheStage, noteinfo)
		}
	}

	// remove from stage and back repo's 3 maps all noteinfos that are not in the checkout
	for noteinfo := range noteinfoInstancesToBeRemovedFromTheStage {
		noteinfo.Unstage(backRepoNoteInfo.GetStage())

		// remove instance from the back repo 3 maps
		noteinfoID := backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]
		delete(backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID, noteinfo)
		delete(backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB, noteinfoID)
		delete(backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr, noteinfoID)
	}

	return
}

// CheckoutPhaseOneInstance takes a noteinfoDB that has been found in the DB, updates the backRepo and stages the
// models version of the noteinfoDB
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CheckoutPhaseOneInstance(noteinfoDB *NoteInfoDB) (Error error) {

	noteinfo, ok := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]
	if !ok {
		noteinfo = new(models.NoteInfo)

		backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID] = noteinfo
		backRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo] = noteinfoDB.ID

		// append model store with the new element
		noteinfo.Name = noteinfoDB.Name_Data.String
		noteinfo.Stage(backRepoNoteInfo.GetStage())
	}
	noteinfoDB.CopyBasicFieldsToNoteInfo(noteinfo)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	noteinfo.Stage(backRepoNoteInfo.GetStage())

	// preserve pointer to noteinfoDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_NoteInfoDBID_NoteInfoDB)[noteinfoDB hold variable pointers
	noteinfoDB_Data := *noteinfoDB
	preservedPtrToNoteInfo := &noteinfoDB_Data
	backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[noteinfoDB.ID] = preservedPtrToNoteInfo

	return
}

// BackRepoNoteInfo.CheckoutPhaseTwo Checkouts all staged instances of NoteInfo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, noteinfoDB := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB {
		backRepoNoteInfo.CheckoutPhaseTwoInstance(backRepo, noteinfoDB)
	}
	return
}

// BackRepoNoteInfo.CheckoutPhaseTwoInstance Checkouts staged instances of NoteInfo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteInfo *BackRepoNoteInfoStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, noteinfoDB *NoteInfoDB) (Error error) {

	noteinfo := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr[noteinfoDB.ID]

	noteinfoDB.DecodePointers(backRepo, noteinfo)

	return
}

func (noteinfoDB *NoteInfoDB) DecodePointers(backRepo *BackRepoStruct, noteinfo *models.NoteInfo) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitNoteInfo allows commit of a single noteinfo (if already staged)
func (backRepo *BackRepoStruct) CommitNoteInfo(noteinfo *models.NoteInfo) {
	backRepo.BackRepoNoteInfo.CommitPhaseOneInstance(noteinfo)
	if id, ok := backRepo.BackRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]; ok {
		backRepo.BackRepoNoteInfo.CommitPhaseTwoInstance(backRepo, id, noteinfo)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitNoteInfo allows checkout of a single noteinfo (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutNoteInfo(noteinfo *models.NoteInfo) {
	// check if the noteinfo is staged
	if _, ok := backRepo.BackRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]; ok {

		if id, ok := backRepo.BackRepoNoteInfo.Map_NoteInfoPtr_NoteInfoDBID[noteinfo]; ok {
			var noteinfoDB NoteInfoDB
			noteinfoDB.ID = id

			if err := backRepo.BackRepoNoteInfo.db.First(&noteinfoDB, id).Error; err != nil {
				log.Fatalln("CheckoutNoteInfo : Problem with getting object with id:", id)
			}
			backRepo.BackRepoNoteInfo.CheckoutPhaseOneInstance(&noteinfoDB)
			backRepo.BackRepoNoteInfo.CheckoutPhaseTwoInstance(backRepo, &noteinfoDB)
		}
	}
}

// CopyBasicFieldsFromNoteInfo
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsFromNoteInfo(noteinfo *models.NoteInfo) {
	// insertion point for fields commit

	noteinfoDB.Name_Data.String = noteinfo.Name
	noteinfoDB.Name_Data.Valid = true

	noteinfoDB.IsKept_Data.Bool = noteinfo.IsKept
	noteinfoDB.IsKept_Data.Valid = true
}

// CopyBasicFieldsFromNoteInfo_WOP
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsFromNoteInfo_WOP(noteinfo *models.NoteInfo_WOP) {
	// insertion point for fields commit

	noteinfoDB.Name_Data.String = noteinfo.Name
	noteinfoDB.Name_Data.Valid = true

	noteinfoDB.IsKept_Data.Bool = noteinfo.IsKept
	noteinfoDB.IsKept_Data.Valid = true
}

// CopyBasicFieldsFromNoteInfoWOP
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsFromNoteInfoWOP(noteinfo *NoteInfoWOP) {
	// insertion point for fields commit

	noteinfoDB.Name_Data.String = noteinfo.Name
	noteinfoDB.Name_Data.Valid = true

	noteinfoDB.IsKept_Data.Bool = noteinfo.IsKept
	noteinfoDB.IsKept_Data.Valid = true
}

// CopyBasicFieldsToNoteInfo
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsToNoteInfo(noteinfo *models.NoteInfo) {
	// insertion point for checkout of basic fields (back repo to stage)
	noteinfo.Name = noteinfoDB.Name_Data.String
	noteinfo.IsKept = noteinfoDB.IsKept_Data.Bool
}

// CopyBasicFieldsToNoteInfo_WOP
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsToNoteInfo_WOP(noteinfo *models.NoteInfo_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	noteinfo.Name = noteinfoDB.Name_Data.String
	noteinfo.IsKept = noteinfoDB.IsKept_Data.Bool
}

// CopyBasicFieldsToNoteInfoWOP
func (noteinfoDB *NoteInfoDB) CopyBasicFieldsToNoteInfoWOP(noteinfo *NoteInfoWOP) {
	noteinfo.ID = int(noteinfoDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	noteinfo.Name = noteinfoDB.Name_Data.String
	noteinfo.IsKept = noteinfoDB.IsKept_Data.Bool
}

// Backup generates a json file from a slice of all NoteInfoDB instances in the backrepo
func (backRepoNoteInfo *BackRepoNoteInfoStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "NoteInfoDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteInfoDB, 0)
	for _, noteinfoDB := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB {
		forBackup = append(forBackup, noteinfoDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json NoteInfo ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json NoteInfo file", err.Error())
	}
}

// Backup generates a json file from a slice of all NoteInfoDB instances in the backrepo
func (backRepoNoteInfo *BackRepoNoteInfoStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteInfoDB, 0)
	for _, noteinfoDB := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB {
		forBackup = append(forBackup, noteinfoDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("NoteInfo")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&NoteInfo_Fields, -1)
	for _, noteinfoDB := range forBackup {

		var noteinfoWOP NoteInfoWOP
		noteinfoDB.CopyBasicFieldsToNoteInfoWOP(&noteinfoWOP)

		row := sh.AddRow()
		row.WriteStruct(&noteinfoWOP, -1)
	}
}

// RestoreXL from the "NoteInfo" sheet all NoteInfoDB instances
func (backRepoNoteInfo *BackRepoNoteInfoStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoNoteInfoid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["NoteInfo"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoNoteInfo.rowVisitorNoteInfo)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoNoteInfo *BackRepoNoteInfoStruct) rowVisitorNoteInfo(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var noteinfoWOP NoteInfoWOP
		row.ReadStruct(&noteinfoWOP)

		// add the unmarshalled struct to the stage
		noteinfoDB := new(NoteInfoDB)
		noteinfoDB.CopyBasicFieldsFromNoteInfoWOP(&noteinfoWOP)

		noteinfoDB_ID_atBackupTime := noteinfoDB.ID
		noteinfoDB.ID = 0
		query := backRepoNoteInfo.db.Create(noteinfoDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[noteinfoDB.ID] = noteinfoDB
		BackRepoNoteInfoid_atBckpTime_newID[noteinfoDB_ID_atBackupTime] = noteinfoDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "NoteInfoDB.json" in dirPath that stores an array
// of NoteInfoDB and stores it in the database
// the map BackRepoNoteInfoid_atBckpTime_newID is updated accordingly
func (backRepoNoteInfo *BackRepoNoteInfoStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoNoteInfoid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "NoteInfoDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json NoteInfo file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*NoteInfoDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_NoteInfoDBID_NoteInfoDB
	for _, noteinfoDB := range forRestore {

		noteinfoDB_ID_atBackupTime := noteinfoDB.ID
		noteinfoDB.ID = 0
		query := backRepoNoteInfo.db.Create(noteinfoDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[noteinfoDB.ID] = noteinfoDB
		BackRepoNoteInfoid_atBckpTime_newID[noteinfoDB_ID_atBackupTime] = noteinfoDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json NoteInfo file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<NoteInfo>id_atBckpTime_newID
// to compute new index
func (backRepoNoteInfo *BackRepoNoteInfoStruct) RestorePhaseTwo() {

	for _, noteinfoDB := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB {

		// next line of code is to avert unused variable compilation error
		_ = noteinfoDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoNoteInfo.db.Model(noteinfoDB).Updates(*noteinfoDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoNoteInfo.ResetReversePointers commits all staged instances of NoteInfo to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNoteInfo *BackRepoNoteInfoStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, noteinfo := range backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoPtr {
		backRepoNoteInfo.ResetReversePointersInstance(backRepo, idx, noteinfo)
	}

	return
}

func (backRepoNoteInfo *BackRepoNoteInfoStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, noteinfo *models.NoteInfo) (Error error) {

	// fetch matching noteinfoDB
	if noteinfoDB, ok := backRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB[idx]; ok {
		_ = noteinfoDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoNoteInfoid_atBckpTime_newID map[uint]uint