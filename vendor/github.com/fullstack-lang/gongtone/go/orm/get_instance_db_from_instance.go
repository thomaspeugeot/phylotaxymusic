// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtone/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | FreqencyDB | NoteDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Freqency:
		freqencyInstance := any(concreteInstance).(*models.Freqency)
		ret2 := backRepo.BackRepoFreqency.GetFreqencyDBFromFreqencyPtr(freqencyInstance)
		ret = any(ret2).(*T2)
	case *models.Note:
		noteInstance := any(concreteInstance).(*models.Note)
		ret2 := backRepo.BackRepoNote.GetNoteDBFromNotePtr(noteInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Freqency:
		tmp := GetInstanceDBFromInstance[models.Freqency, FreqencyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note:
		tmp := GetInstanceDBFromInstance[models.Note, NoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Freqency:
		tmp := GetInstanceDBFromInstance[models.Freqency, FreqencyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note:
		tmp := GetInstanceDBFromInstance[models.Note, NoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}