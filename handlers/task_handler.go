package handlers

import (
	"net/http"

	"github.com/Michael-Ralph/go-task-manage/models"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

var tasks = []models.Task{}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = uuid.New().String()
	newTask.Completed = false
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, tasks := range tasks {
		if tasks.ID == id {
			c.JSON(http.StatusOK, tasks)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Completed = updatedTask.Completed
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted succesfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
