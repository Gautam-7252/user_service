package main

import (
	"context"
	"testing"

	pb "user_service/proto"
)

func TestGetUser(t *testing.T) {
	s := &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.GetUserRequest{Id: 1}
	res, err := s.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.User.Id != 1 {
		t.Errorf("expected user ID 1, got %v", res.User.Id)
	}
}

func TestListUsers(t *testing.T) {
	s := &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.ListUsersRequest{Ids: []int32{1, 2}}
	res, err := s.ListUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 2 {
		t.Errorf("expected 2 users, got %v", len(res.Users))
	}
}

func TestSearchUsers(t *testing.T) {
	s := &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.SearchUsersRequest{City: "LA"}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 1 {
		t.Errorf("expected 1 user, got %v", len(res.Users))
	} else {
		t.Logf("Search result: %v", res.Users[0])
	}
}

func TestSearchUsersPartialCriteria(t *testing.T) {
	s := &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	// Search by phone number
	req := &pb.SearchUsersRequest{Phone: 1234567890}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 1 {
		t.Errorf("expected 1 user, got %v", len(res.Users))
	} else {
		t.Logf("Search result by phone: %v", res.Users[0])
	}

	// Search by marital status
	req = &pb.SearchUsersRequest{Married: true}
	res, err = s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 1 {
		t.Errorf("expected 1 user, got %v", len(res.Users))
	} else {
		t.Logf("Search result by marital status: %v", res.Users[0])
	}
}
