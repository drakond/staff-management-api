package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage *MongoStorage
}

func NewHandler(storage *MongoStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateDepartment(c *gin.Context) {
	var department Department

	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.InsertDepartment(&department); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": department.ID})
}

func (h *Handler) GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	department, err := h.storage.GetDepartment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (h *Handler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.DeleteDepartment(id); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "department deleted successfully"})
}

func (h *Handler) AddEmployeeToDepartment(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.AddEmployeeToDepartment(departmentID, employeeID); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "employee added to department successfully"})
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.Insert(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": employee.ID})
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	employees, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.Update(id, employee); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "employee updated successfully"})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "employee deleted successfully"})
}
