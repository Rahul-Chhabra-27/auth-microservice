package main

import (
	"auth-micro/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. SaveTaskToDB
// 2.GenerateTask
func GenerateTask(context *gin.Context) {
	var task model.Task;
	err := context.ShouldBindJSON(&task);
	fmt.Println(task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : "Bad Request",
			"error":err,
		});
	}  else {
		task.ID = 10;
		task.SaveTask();
		context.JSON(http.StatusCreated, gin.H{
			"message" : "Task has been successfully created!",
			"task": task,
		})
	}
}
func GetAllTheTasks(context *gin.Context) {
	tasks := model.GetAllTheTasks();
	context.JSON(http.StatusOK, gin.H{
		"message":"Successfully fetched all the tasks",
		"tasks":tasks,
	})
}
// entry point
func main() {
	// configuration of the http server.
	httpServer := gin.Default();

	//? Method : @POST
	// ? Endpoint Route : /save-task 
	httpServer.POST("/save-task", GenerateTask);
	//? Method : @GET
	// ? Endpoint Route : /get-tasks 
	httpServer.GET("/get-tasks",GetAllTheTasks);
	
	// running the server
	httpServer.Run(":8080");  // Infinite loop
	
}