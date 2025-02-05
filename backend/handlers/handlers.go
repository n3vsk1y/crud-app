package handlers

import (
 	"net/http"
 	"github.com/gin-gonic/gin"
 	"github.com/n3vsk1y/crud-app/models"
)

var tasks []models.Task
var idCounter = 1

func CreateTask(c *gin.Context) {
	var newTask models.Task
 	if err := c.ShouldBindJSON(&newTask); err != nil {
  		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
 	}
 	newTask.ID = idCounter
 	idCounter++
 	tasks = append(tasks, newTask)
 	c.JSON(http.StatusCreated, newTask)
}

func GetTasks(c *gin.Context) {
 	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
 	id := c.Param("id")
 	for i, task := range tasks {
  		if id == string(rune(task.ID)) {
   			c.ShouldBindJSON(&tasks[i])
   			c.JSON(http.StatusOK, tasks[i])
   			return
  		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if id == string(rune(task.ID)) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Status(http.StatusNoContent)
   			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}