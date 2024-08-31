package handlers

import (
	"context"
	"fmt"

	pb "Assignment_TotalityCorp/grpc"
	"Assignment_TotalityCorp/models"
	"Assignment_TotalityCorp/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	service services.UserService
}

func NewUserServiceServer(service services.UserService) pb.UserServiceServer {
	return &userServiceServer{service: service}
}

func (s *userServiceServer) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if err := validateUserID(req.Id); err != nil {
		return nil, err
	}

	user, found := s.service.GetUserById(req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "User with ID %d not found", req.Id)
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:      user.Id,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		},
		Message: "User found",
	}, nil
}

func (s *userServiceServer) GetUsersByIds(req *pb.UsersRequest, stream pb.UserService_GetUsersByIdsServer) error {
	users := s.service.GetUsersByIds(req.Ids)

	if err := validateUserIDs(req.Ids); err != nil {
		return err
	}

	for _, user := range users {
		if err := stream.Send(&pb.UserResponse{
			User: &pb.User{
				Id:      user.Id,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			},
			Message: "User found",
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *userServiceServer) SearchUsers(req *pb.SearchRequest, stream pb.UserService_SearchUsersServer) error {
	var criteria []models.SearchCriteria

	for _, c := range req.Criteria {
		if !isValidField(c.Field) {
			return fmt.Errorf("invalid search field: %s", c.Field)
		}

		criteria = append(criteria, models.SearchCriteria{Field: c.Field, Keyword: c.Keyword})
	}

	users := s.service.SearchUsers(criteria)

	for _, user := range users {
		if err := stream.Send(&pb.UserResponse{
			User: &pb.User{
				Id:      user.Id,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			},
			Message: "User found",
		}); err != nil {
			return err
		}
	}

	return nil
}
