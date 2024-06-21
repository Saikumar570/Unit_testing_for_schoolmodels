package repositories

import (
	"School_api/models"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type SchoolRepository interface {
	FindAll() ([]models.School, error)
	FindByID(id uint) (models.School, error)
	Create(school models.School) (models.School, error)
	Update(school models.School) (models.School, error)
	Delete(school models.School) error
}

type StudentRepository interface {
	FindAll() ([]models.Student, error)
	FindByID(id uint) (models.Student, error)
	Create(student models.Student) (models.Student, error)
	Update(student models.Student) (models.Student, error)
	Delete(student models.Student) error
}

type schoolRepository struct {
	db *gorm.DB
}

type studentRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{db}
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *schoolRepository) FindAll() ([]models.School, error) {
	var schools []models.School
	err := r.db.Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) FindByID(id uint) (models.School, error) {
	var school models.School
	err := r.db.First(&school, id).Error
	return school, err
}

func (r *schoolRepository) Create(school models.School) (models.School, error) {
	err := r.db.Create(&school).Error
	return school, err
}

func (r *schoolRepository) Update(school models.School) (models.School, error) {
	err := r.db.Save(&school).Error
	return school, err
}

func (r *schoolRepository) Delete(school models.School) error {
	return r.db.Delete(&school).Error
}

func (r *studentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByID(id uint) (models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	return student, err
}

func (r *studentRepository) Create(student models.Student) (models.Student, error) {
	err := r.db.Create(&student).Error
	return student, err
}

func (r *studentRepository) Update(student models.Student) (models.Student, error) {
	err := r.db.Save(&student).Error
	return student, err
}

func (r *studentRepository) Delete(student models.Student) error {
	return r.db.Delete(&student).Error
}
