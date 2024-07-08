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
		if req.Married {
			if user.Married != req.Married {
				match = false
			}
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
	pb.RegisterUserServiceServer(s, &server{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			// Add more users as needed
		},
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
