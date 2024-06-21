package services

import (
	"School_api/models"
	"School_api/repositories"
	"School_api/utils"
)

type SchoolService interface {
	GetSchools() ([]models.School, error)
	GetSchoolByID(id uint) (models.School, error)
	CreateSchool(school models.School) (models.School, error)
	UpdateSchool(id uint, school models.School) (models.School, error)
	DeleteSchool(id uint) error
}

type StudentService interface {
	GetStudents() ([]models.Student, error)
	GetStudentByID(id uint) (models.Student, error)
	CreateStudent(student models.Student) (models.Student, error)
	UpdateStudent(id uint, student models.Student) (models.Student, error)
	DeleteStudent(id uint) error
}

type schoolService struct {
	repo repositories.SchoolRepository
}

type studentService struct {
	repo repositories.StudentRepository
}

func NewSchoolService() SchoolService {
	db := utils.InitDB()
	repo := repositories.NewSchoolRepository(db)
	return &schoolService{repo}
}

func NewStudentService() StudentService {
	db := utils.InitDB()
	repo := repositories.NewStudentRepository(db)
	return &studentService{repo}
}

func (s *schoolService) GetSchools() ([]models.School, error) {
	return s.repo.FindAll()
}

func (s *schoolService) GetSchoolByID(id uint) (models.School, error) {
	return s.repo.FindByID(id)
}

func (s *schoolService) CreateSchool(school models.School) (models.School, error) {
	return s.repo.Create(school)
}

func (s *schoolService) UpdateSchool(id uint, school models.School) (models.School, error) {
	existingSchool, err := s.repo.FindByID(id)
	if err != nil {
		return existingSchool, err
	}
	existingSchool.Name = school.Name
	existingSchool.SchoolId = school.SchoolId
	existingSchool.SchoolAddress = school.SchoolAddress
	existingSchool.Classes = school.Classes
	return s.repo.Update(existingSchool)
}

func (s *schoolService) DeleteSchool(id uint) error {
	school, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(school)
}

func (s *studentService) GetStudents() ([]models.Student, error) {
	return s.repo.FindAll()
}

func (s *studentService) GetStudentByID(id uint) (models.Student, error) {
	return s.repo.FindByID(id)
}

func (s *studentService) CreateStudent(student models.Student) (models.Student, error) {
	return s.repo.Create(student)
}

func (s *studentService) UpdateStudent(id uint, student models.Student) (models.Student, error) {
	existingStudent, err := s.repo.FindByID(id)
	if err != nil {
		return existingStudent, err
	}
	existingStudent.Name = student.Name
	existingStudent.Age = student.Age
	existingStudent.AddressStruct = student.AddressStruct
	return s.repo.Update(existingStudent)
}

func (s *studentService) DeleteStudent(id uint) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(student)
}
