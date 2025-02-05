package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/n3vsk1y/crud-app/handlers"
)

func RegRoutes(r *gin.Engine)  {
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", handlers.CreateTask)
		taskRoutes.GET("/", handlers.GetTasks)
		taskRoutes.PUT("/:id", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}
}