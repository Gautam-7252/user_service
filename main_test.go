package main

import (
	"context"
	"testing"

	pb "user_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUser(t *testing.T) {
	server := NewServer()

	tests := []struct {
		id       int32
		expected *pb.User
	}{
		{1, &pb.User{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}},
		{99, nil},
	}

	for _, tt := range tests {
		req := &pb.GetUserRequest{Id: tt.id}
		res, err := server.GetUser(context.Background(), req)
		if tt.expected == nil {
			if status.Code(err) != codes.NotFound {
				t.Fatalf("expected error code %v, got %v", codes.NotFound, status.Code(err))
			}
		} else {
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if res.User.Id != tt.expected.Id {
				t.Fatalf("expected user id %v, got %v", tt.expected.Id, res.User.Id)
			}
		}
	}
}

func TestListUsers(t *testing.T) {
	server := NewServer()

	tests := []struct {
		ids      []int32
		expected []*pb.User
	}{
		{[]int32{1, 2}, []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 1234567891, Height: 6.0, Married: false},
		}},
		{[]int32{99}, []*pb.User{}},
	}

	for _, tt := range tests {
		req := &pb.ListUsersRequest{Ids: tt.ids}
		res, err := server.ListUsers(context.Background(), req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(res.Users) != len(tt.expected) {
			t.Fatalf("expected %v users, got %v", len(tt.expected), len(res.Users))
		}
	}
}

func TestSearchUsers(t *testing.T) {
	server := NewServer()

	tests := []struct {
		req      *pb.SearchUsersRequest
		expected []*pb.User
	}{
		{&pb.SearchUsersRequest{City: "LA"}, []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 5, Fname: "David", City: "LA", Phone: 1234567894, Height: 5.5, Married: false},
		}},
		{&pb.SearchUsersRequest{Phone: 1234567890}, []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		}},
		{&pb.SearchUsersRequest{Married: true}, []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 4, Fname: "Gracia", City: "NY", Phone: 1234567893, Height: 5.7, Married: true},
		}},
		{&pb.SearchUsersRequest{City: "NY", Married: true}, []*pb.User{
			{Id: 4, Fname: "Gracia", City: "NY", Phone: 1234567893, Height: 5.7, Married: true},
		}},
	}

	for _, tt := range tests {
		res, err := server.SearchUsers(context.Background(), tt.req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(res.Users) != len(tt.expected) {
			t.Fatalf("expected %v users, got %v", len(tt.expected), len(res.Users))
		}
		for i, user := range res.Users {
			if user.Id != tt.expected[i].Id {
				t.Fatalf("expected user id %v, got %v", tt.expected[i].Id, user.Id)
			}
		}
	}
}
