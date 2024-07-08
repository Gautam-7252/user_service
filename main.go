package main

import (
	"context"
	"log"
	"net"

	pb "user_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func NewServer() *server {
	return &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 1234567891, Height: 6.0, Married: false},
			{Id: 3, Fname: "Mello", City: "VG", Phone: 1234567892, Height: 5.11, Married: false},
			{Id: 4, Fname: "Gracia", City: "NY", Phone: 1234567893, Height: 5.7, Married: true},
			{Id: 5, Fname: "David", City: "LA", Phone: 1234567894, Height: 5.5, Married: false},
		},
	}
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserResponse{User: user}, nil
		}
	}
	return nil, grpc.Errorf(codes.NotFound, "User not found")
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, user)
			}
		}
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var users []*pb.User
	for _, user := range s.users {
		match := true
		if req.Id != 0 {
			if user.Id != req.Id {
				match = false
			}
		}
		if req.Fname != "" {
			if user.Fname != req.Fname {
				match = false
			}
		}
		if req.City != "" {
			if user.City != req.City {
				match = false
			}
		}
		if req.Phone != 0 {
			if user.Phone != req.Phone {
				match = false
			}
		}
		if req.Height != 0.0 {
			if user.Height != req.Height {
				match = false
			}
		}
		if req.Married && !user.Married {
			match = false
		}
		if match {
			users = append(users, user)
		}
	}
	return &pb.SearchUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
