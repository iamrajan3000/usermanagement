package services

import "Assignment_TotalityCorp/models"

type UserService interface {
	GetUserById(id int32) (*models.User, bool)
	GetUsersByIds(ids []int32) []*models.User
	SearchUsers(criteria []models.SearchCriteria) []*models.User
}
