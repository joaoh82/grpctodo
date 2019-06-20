package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/joaoh82/shelltodo/pb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type todoServer struct {
	DB *sql.DB
}

func (s *todoServer) AddTask(ctx context.Context, req *pb.TodoRequest) (*pb.TodoRequest, error) {
	_, err := s.DB.Exec("INSERT INTO todos(title, done) VALUES($1, $2)", req.GetTask().Title, req.GetTask().Done)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *todoServer) ListTasks(ctx context.Context, req *pb.Empty) (*pb.TodoResponse, error) {
	var (
		tasks []*pb.TodoMessage
		title string
		done  bool
	)

	rows, err := s.DB.Query("SELECT title, done FROM todos")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&title, &done)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &pb.TodoMessage{
			Title: title,
			Done:  done,
		})
	}

	return &pb.TodoResponse{
		Tasks: tasks,
	}, nil
}

const (
	dbHost     = "localhost"
	dbName     = "grpc_todo"
	dbPassword = "1234"
	dbPort     = "5432"
	dbUser     = "postgres"
	port       = ":50051"
)

func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (title varchar(50), done BOOLEAN)")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoServer{
		DB: db,
	})

	fmt.Println("=====> Server started.")
	s.Serve(lis)
}
