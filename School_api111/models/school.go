package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name          string        `json:"name"`
	SchoolId      string        `json:"school_id"`
	SchoolAddress SchoolAddress `gorm:"embedded" json:"school_address"`
	Classes       []Class       `gorm:"-" json:"classes"`
	ClassesJSON   string        `gorm:"column:classes_json" json:"-"`
}

type SchoolAddress struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Class struct {
	ClassName      string `json:"class_name"`
	StudentCount   int    `json:"student_count"`
	StudentIDs     []uint `gorm:"-" json:"student_ids"`
	StudentIDsJSON string `gorm:"column:student_ids_json" json:"-"`
}

type Student struct {
	gorm.Model
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	AddressStruct *Address `gorm:"-" json:"address"`
	AddressDb     string   `gorm:"column:student_address" json:"-"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func (s *School) BeforeSave(tx *gorm.DB) (err error) {
	if len(s.Classes) > 0 {
		for i, class := range s.Classes {
			if len(class.StudentIDs) > 0 {
				data, err := json.Marshal(class.StudentIDs)
				if err != nil {
					return err
				}
				s.Classes[i].StudentIDsJSON = string(data)
			}
		}
		data, err := json.Marshal(s.Classes)
		if err != nil {
			return err
		}
		s.ClassesJSON = string(data)
	}
	return nil
}

func (s *School) AfterFind(tx *gorm.DB) (err error) {
	if s.ClassesJSON != "" {
		if err := json.Unmarshal([]byte(s.ClassesJSON), &s.Classes); err != nil {
			return err
		}
		for i, class := range s.Classes {
			if class.StudentIDsJSON != "" {
				if err := json.Unmarshal([]byte(class.StudentIDsJSON), &s.Classes[i].StudentIDs); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *Student) BeforeSave(tx *gorm.DB) (err error) {
	if s.AddressStruct != nil {
		data, err := json.Marshal(s.AddressStruct)
		if err != nil {
			return err
		}
		s.AddressDb = string(data)
	}
	return nil
}

func (s *Student) AfterFind(tx *gorm.DB) (err error) {
	if s.AddressDb != "" {
		if err := json.Unmarshal([]byte(s.AddressDb), &s.AddressStruct); err != nil {
			return err
		}
	}
	return nil
}
