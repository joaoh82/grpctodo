package httprouter

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaoh82/shelltodo/pb"
	"google.golang.org/grpc"
)

func SetupRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api/v1")
	{
		// This endpoint could might as well be a health check
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/listtasks", ListTasksHandler)
		api.POST("/addtask", AddTasksHandler)
	}

	return router
}

func AddTasksHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	// Creating connection with gRPC Server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	var task pb.TodoMessage
	// Binding body into custom struct
	if c.ShouldBind(&task) == nil {
		// Logging added for debugging purposes.
		log.Println(task.Title)
		log.Println(task.Done)
	}
	taskRequest := &pb.TodoRequest{
		Task: &task,
	}

	//
	_, err = client.AddTask(context.Background(), taskRequest)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task created",
	})
}

func ListTasksHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	// Creating connection with gRPC Server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	tasks, err := client.ListTasks(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, string(response))
}
