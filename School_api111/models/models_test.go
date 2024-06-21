package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchoolBeforeSave(t *testing.T) {
	school := School{
		Classes: []Class{
			{
				ClassName:  "Class 1",
				StudentIDs: []uint{1, 2, 3},
			},
		},
	}

	err := school.BeforeSave(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, school.ClassesJSON)
}

func TestSchoolAfterFind(t *testing.T) {
	classes := []Class{
		{
			ClassName:  "Class 1",
			StudentIDs: []uint{1, 2, 3},
		},
	}
	classesJSON, _ := json.Marshal(classes)

	school := School{
		ClassesJSON: string(classesJSON),
	}

	err := school.AfterFind(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, school.Classes)
}

func TestStudentBeforeSave(t *testing.T) {
	address := &Address{
		Street:  "old street",
		City:    "hyd",
		State:   "TS",
		ZipCode: "535001",
	}

	student := Student{
		AddressStruct: address,
	}

	err := student.BeforeSave(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, student.AddressDb)
}

func TestStudentAfterFind(t *testing.T) {
	address := &Address{
		Street:  "old street",
		City:    "hyd",
		State:   "TS",
		ZipCode: "535001",
	}
	addressJSON, _ := json.Marshal(address)

	student := Student{
		AddressDb: string(addressJSON),
	}

	err := student.AfterFind(nil)
	assert.Nil(t, err)
	assert.NotNil(t, student.AddressStruct)
}
