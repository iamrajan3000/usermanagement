package services

import (
	"testing"

	"Assignment_TotalityCorp/models"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	tests := []struct {
		users          []*models.User
		searchId       int32
		expectedUser   *models.User
		expectedExists bool
	}{
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
			searchId:       2,
			expectedUser:   &models.User{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
			expectedExists: true,
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
			searchId:       4,
			expectedUser:   nil,
			expectedExists: false,
		},
	}

	for i, tc := range tests {
		service := &userService{users: tc.users}

		user, exists := service.GetUserById(tc.searchId)

		assert.Equalf(t, tc.expectedUser, user, "TestCase[%v] Failed", i)
		assert.Equalf(t, tc.expectedExists, exists, "TestCase[%v] Failed", i)
	}
}

func TestGetUsersByIds(t *testing.T) {
	tests := []struct {
		users          []*models.User
		searchIds      []int32
		expectedUsers  []*models.User
		expectedLength int
	}{
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
			searchIds: []int32{2, 3},
			expectedUsers: []*models.User{
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
			expectedLength: 2,
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
			searchIds: []int32{4, 1},
			expectedUsers: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
			},
			expectedLength: 1,
		},
		{
			users:          []*models.User{},
			searchIds:      []int32{1, 2, 3},
			expectedUsers:  nil,
			expectedLength: 0,
		},
	}

	for i, tc := range tests {
		service := &userService{users: tc.users}

		users := service.GetUsersByIds(tc.searchIds)

		assert.Equalf(t, tc.expectedUsers, users, "TestCase[%v] Failed", i)
	}
}

func TestSearchUsers(t *testing.T) {
	tests := []struct {
		users         []*models.User
		criteria      []models.SearchCriteria
		expectedUsers []*models.User
	}{
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
				{Id: 4, Fname: "David", City: "San Francisco", Phone: 1357924680, Height: 5.8, Married: false},
			},
			criteria: []models.SearchCriteria{
				{Field: "fname", Keyword: "Alice"},
			},
			expectedUsers: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
			},
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: false},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
				{Id: 4, Fname: "David", City: "San Francisco", Phone: 1357924680, Height: 5.8, Married: false},
			},
			criteria: []models.SearchCriteria{
				{Field: "city", Keyword: "chicago"},
				{Field: "married", Keyword: "true"},
			},
			expectedUsers: []*models.User{
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
				{Id: 4, Fname: "David", City: "San Francisco", Phone: 1357924680, Height: 5.8, Married: false},
			},
			criteria: []models.SearchCriteria{
				{Field: "height", Keyword: "5.9"},
				{Field: "phone", Keyword: "2468109753"},
			},
			expectedUsers: []*models.User{
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
			},
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
				{Id: 4, Fname: "David", City: "San Francisco", Phone: 1357924680, Height: 5.8, Married: false},
			},
			criteria: []models.SearchCriteria{
				{Field: "phone", Keyword: "9876543210"},
			},
			expectedUsers: []*models.User{
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
			},
		},
		{
			users: []*models.User{
				{Id: 1, Fname: "Alice", City: "New York", Phone: 1234567890, Height: 5.6, Married: true},
				{Id: 2, Fname: "Bob", City: "Los Angeles", Phone: 9876543210, Height: 6.0, Married: false},
				{Id: 3, Fname: "Charlie", City: "Chicago", Phone: 2468109753, Height: 5.9, Married: true},
				{Id: 4, Fname: "David", City: "San Francisco", Phone: 1357924680, Height: 5.8, Married: false},
			},
			criteria: []models.SearchCriteria{
				{Field: "unknown", Keyword: "value"},
			},
			expectedUsers: nil,
		},
	}

	for i, tc := range tests {
		service := &userService{users: tc.users}

		users := service.SearchUsers(tc.criteria)

		assert.Equalf(t, tc.expectedUsers, users, "TestCase[%v] Failed", i)
	}
}
