package grpc

import (
	"context"
	"testing"

	"github.com/joaoh82/shelltodo/pb"
)

func TestAddTast(t *testing.T) {
	s := todoServer{}

	tests := []struct {
		task *pb.TodoRequest
		want *pb.TodoRequest
	}{
		{
			task: &pb.TodoRequest{
				Task: &pb.TodoMessage{
					Title: "Test",
					Done:  true,
				},
			},
			want: &pb.TodoRequest{
				Task: &pb.TodoMessage{
					Title: "Test",
					Done:  true,
				},
			},
		},
		{
			task: &pb.TodoRequest{
				Task: &pb.TodoMessage{
					Title: "Test2",
					Done:  true,
				},
			},
			want: &pb.TodoRequest{
				Task: &pb.TodoMessage{
					Title: "Test2",
					Done:  true,
				},
			},
		},
	}

	for _, tt := range tests {
		req := tt.task
		resp, err := s.AddTask(context.Background(), req)
		if err != nil {
			t.Errorf("AddTask(%v) got unexpected error", err)
		}
		if resp.GetTask().Title != tt.want.GetTask().Title {
			t.Errorf("AddTask(%v)=%v, wanted %v", tt.task.GetTask().Title, resp.GetTask().Title, tt.want.GetTask().Title)
		}
	}
}

func TestNumberOfReturnedItemsListTasks(t *testing.T) {
	s := todoServer{}

	req := &pb.Empty{}
	resp, err := s.ListTasks(context.Background(), req)
	if err != nil {
		t.Errorf("ListTasks(%v) got unexpected error", err)
	}
	if len(resp.GetTasks()) <= 0 {
		t.Errorf("%v", resp.GetTasks())
		t.Errorf("ListTasks(%v)=%v, wanted more then 0", len(resp.GetTasks()), len(resp.GetTasks()))
	}
}
