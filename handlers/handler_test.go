package handlers

import (
	pb "Assignment_TotalityCorp/grpc"
	"Assignment_TotalityCorp/models"
	"Assignment_TotalityCorp/services"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := services.NewMockUserService(ctrl)

	server := &userServiceServer{service: mockService}
	ctx := context.Background()

	tests := []struct {
		name        string
		id          int32
		isUserFound bool
		mockUser    *models.User
		expUser     *pb.UserResponse
		expErr      error
	}{
		{
			name:        "Valid user ID",
			id:          1,
			isUserFound: true,
			mockUser:    &models.User{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
			expUser: &pb.UserResponse{Message: "User found",
				User: &pb.User{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true}},
		},
		{
			name:        "Invalid user ID",
			id:          -1,
			isUserFound: true,
			expErr:      status.Errorf(codes.InvalidArgument, "Invalid user ID: %d", -1),
		},
		{
			name:        "User not found",
			id:          3,
			isUserFound: false,
			expErr:      status.Errorf(codes.NotFound, "User with ID %d not found", 3),
		},
	}

	for i, tc := range tests {
		req := &pb.UserRequest{Id: tc.id}
		mockService.EXPECT().GetUserById(tc.id).Return(tc.mockUser, tc.isUserFound).AnyTimes()
		resp, err := server.GetUserById(ctx, req)

		assert.Equalf(t, tc.expErr, err, "TestCase[%v] Failed, Name : %v", i, tc.name)
		assert.Equalf(t, tc.expUser, resp, "TestCase[%v] Failed, Name : %v", i, tc.name)
	}
}

func TestGetUsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := services.NewMockUserService(ctrl)
	mockStream := pb.NewMockUserService_GetUsersByIdsServer(ctrl)

	server := &userServiceServer{service: mockService}

	tests := []struct {
		name      string
		ids       []int32
		mockUsers []*models.User
		expResp   []*pb.UserResponse
		expErr    error
	}{
		{
			name: "Valid user IDs",
			ids:  []int32{1, 2},
			mockUsers: []*models.User{{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false}},
			expResp: []*pb.UserResponse{
				{User: &pb.User{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true}, Message: "User found"},
				{User: &pb.User{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false}, Message: "User found"},
			},
		},
		{
			name:    "Invalid user IDs",
			ids:     []int32{-1},
			expErr:  status.Errorf(codes.InvalidArgument, "Invalid user ID: %d", -1),
			expResp: []*pb.UserResponse{},
		},
	}

	for i, tc := range tests {
		req := &pb.UsersRequest{Ids: tc.ids}

		mockService.EXPECT().GetUsersByIds(tc.ids).Return(tc.mockUsers).AnyTimes()
		for _, user := range tc.mockUsers {
			mockStream.EXPECT().Send(&pb.UserResponse{
				User: &pb.User{
					Id:      user.Id,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				},
				Message: "User found",
			})

		}

		err := server.GetUsersByIds(req, mockStream)

		assert.Equal(t, tc.expErr, err, "TestCase[%v] Failed, Name : %v", i, tc.name)
	}
}

func TestSearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockUserService(ctrl)
	mockStream := pb.NewMockUserService_SearchUsersServer(ctrl)

	server := &userServiceServer{service: mockService}

	tests := []struct {
		name           string
		criteria       []*pb.SearchCriteria
		mockCriteria   []models.SearchCriteria
		mockUsers      []*models.User
		expResp        []*pb.UserResponse
		expErr         error
		expectedErrMsg string
	}{
		{
			name: "Valid search criteria",
			criteria: []*pb.SearchCriteria{
				{Field: "fname", Keyword: "Alice"},
			},
			mockCriteria: []models.SearchCriteria{
				{Field: "fname", Keyword: "Alice"},
			},
			mockUsers: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
			},
			expResp: []*pb.UserResponse{
				{
					User: &pb.User{
						Id:      1,
						Fname:   "Alice",
						City:    "New York",
						Phone:   1234567890,
						Height:  5.6,
						Married: true,
					},
					Message: "User found",
				},
			},
		},

		{
			name: "Invalid search field",
			criteria: []*pb.SearchCriteria{
				{Field: "invalid", Keyword: "Alice"},
			},
			expErr: fmt.Errorf("invalid search field: %s", "invalid"),
		},
		{
			name: "No users found",
			criteria: []*pb.SearchCriteria{
				{Field: "fname", Keyword: "Nonexistent"},
			},
			mockCriteria: []models.SearchCriteria{
				{Field: "fname", Keyword: "Nonexistent"},
			},
			mockUsers: []*models.User{},
			expResp:   []*pb.UserResponse{},
		},
	}

	for i, tc := range tests {
		req := &pb.SearchRequest{Criteria: tc.criteria}

		mockService.EXPECT().SearchUsers(tc.mockCriteria).Return(tc.mockUsers).AnyTimes()

		for _, expected := range tc.expResp {
			mockStream.EXPECT().Send(expected).Return(nil)
		}

		err := server.SearchUsers(req, mockStream)

		assert.Equalf(t, tc.expErr, err, "TestCase[%v] Failed, Name : %v", i, tc.name)
	}
}
