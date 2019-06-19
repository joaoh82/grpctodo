// Client Example
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joaoh82/shelltodo/pb"
	"google.golang.org/grpc"
)

const (
	// "list"
	// "add"
	command = "list"
	addr    = "localhost:50051"
	usage   = `
usage:
  client add <title> <done>
  client list
`
)

func main() {
	title := "buy chicken"
	done := false

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
		fmt.Println(usage)
		os.Exit(1)
	}
}
