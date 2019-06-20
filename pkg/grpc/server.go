package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joaoh82/shelltodo/pb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Task struct {
	title string `json:"title" binding:"required"`
	done  bool   `json:"done" binding:"required"`
}

var tasksDB = []Task{
	Task{"get eggs", false},
	Task{"fill car with gas", false},
	Task{"pay electric bill", false},
}

type todoServer struct {
}

const (
	port = ":50051"
)

func (s *todoServer) AddTask(ctx context.Context, req *pb.TodoRequest) (*pb.TodoRequest, error) {
	tasksDB = append(tasksDB, Task{req.GetTask().Title, req.GetTask().Done})

	return req, nil
}

func (s *todoServer) ListTasks(ctx context.Context, req *pb.Empty) (*pb.TodoResponse, error) {
	var (
		tasks []*pb.TodoMessage
	)

	for _, v := range tasksDB {
		tasks = append(tasks, &pb.TodoMessage{
			Title: v.title,
			Done:  v.done,
		})
	}

	return &pb.TodoResponse{
		Tasks: tasks,
	}, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoServer{})

	fmt.Printf("gRPC Server started at port %v", port)
	s.Serve(lis)
}
