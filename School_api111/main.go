package main

import (
	"School_api/controllers"
	"School_api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	schoolService := services.NewSchoolService()
	studentService := services.NewStudentService()

	// School routes
	r.GET("/schools", controllers.GetSchools(schoolService))
	r.GET("/schools/:id", controllers.GetSchoolByID(schoolService))
	r.POST("/schools", controllers.CreateSchool(schoolService))
	r.PUT("/schools/:id", controllers.UpdateSchool(schoolService))
	r.DELETE("/schools/:id", controllers.DeleteSchool(schoolService))

	// Student routes
	r.GET("/students", controllers.GetStudents(studentService))
	r.GET("/students/:id", controllers.GetStudentByID(studentService))
	r.POST("/students", controllers.CreateStudent(studentService))
	r.PUT("/students/:id", controllers.UpdateStudent(studentService))
	r.DELETE("/students/:id", controllers.DeleteStudent(studentService))

	r.Run(":8080")
}
