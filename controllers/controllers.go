package controllers

import (
	"gim-api-rest-go/database"
	"gim-api-rest-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func ShowStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusNoContent, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func FilterStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}