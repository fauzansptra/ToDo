package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/todos", controllers.GetTodos)
		api.POST("/todos", controllers.CreateTodo)
		api.PUT("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)
	}
	return router
}
