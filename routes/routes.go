package routes

import (
	"gim-api-rest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("/", controllers.ListStudents)
			students.GET("/:id", controllers.ShowStudent)
			students.POST("/", controllers.CreateStudent)
			students.DELETE("/:id", controllers.DeleteStudent)
			students.PATCH("/:id", controllers.UpdateStudent)
			students.GET("/filter/:cpf", controllers.FilterStudentByCPF)
		}
	}

	router.Run()
}