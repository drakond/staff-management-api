package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoStorage, err := NewMongoStorage("mongodb://localhost:27017", "your_database_name")
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB storage: %v", err)
	}

	handler := NewHandler(mongoStorage)
	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.GET("/employee", handler.GetAllEmployees)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.POST("/department", handler.CreateDepartment)
	router.GET("/department/:id", handler.GetDepartment)
	router.DELETE("/department/:id", handler.DeleteDepartment)
	router.PUT("/department/:department_id/employee/:employee_id", handler.AddEmployeeToDepartment)

	router.Run()
}
