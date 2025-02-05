package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/models"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.ShouldBindJSON(&todo)
	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Todo{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
