// Client Example
// Client created mainly for testing gRPC Server and connection to database
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joaoh82/shelltodo/pb"
	"google.golang.org/grpc"
)

const (
	// "list"
	// "add"
	// command = "list"
	addr = "localhost:50051"
	help = `
help:
  client add <title> <done>
  client list
`
)

func main() {
	// testing data
	// title := "buy chicken"
	// done := false
	var command, title string
	var done bool

	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	command = os.Args[1]

	if command == "add" {
		if len(os.Args) != 4 {
			fmt.Println(help)
			os.Exit(1)
		}

		title = os.Args[2]
		done, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			fmt.Printf("done Type no correct: %T \n", done)
			fmt.Println(help)
			os.Exit(1)
		}
		// done = os.Args[3]
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	switch command {
	case "add":
		task := &pb.TodoRequest{
			Task: &pb.TodoMessage{
				Title: title,
				Done:  done,
			},
		}

		_, err = client.AddTask(context.Background(), task)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Suceeded.")

	case "list":
		tasks, err := client.ListTasks(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		for _, task := range tasks.Tasks {
			fmt.Printf("%s: %v\n", task.Title, task.Done)
		}

	default:
		fmt.Println(help)
		os.Exit(1)
	}
}
