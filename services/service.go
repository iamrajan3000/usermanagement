package services

import (
	"strconv"
	"strings"

	"Assignment_TotalityCorp/models"
)

type userService struct {
	users []*models.User
}

func NewUserService(users []*models.User) UserService {
	return &userService{users: users}
}

func (r *userService) GetUserById(id int32) (*models.User, bool) {
	for _, user := range r.users {
		if user.Id == id {
			return user, true
		}
	}

	return nil, false
}

func (r *userService) GetUsersByIds(ids []int32) []*models.User {
	var result []*models.User

	for _, id := range ids {
		if user, found := r.GetUserById(id); found {
			result = append(result, user)
		}
	}

	return result
}

func (r *userService) SearchUsers(criteria []models.SearchCriteria) []*models.User {
	var result []*models.User

	for _, user := range r.users {
		match := true

		for _, c := range criteria {
			switch c.Field {
			case "fname":
				if !contains(user.Fname, c.Keyword) {
					match = false
				}
			case "city":
				if !contains(user.City, c.Keyword) {
					match = false
				}
			case "phone":
				if !contains(strconv.FormatUint(user.Phone, 10), c.Keyword) {
					match = false
				}
			case "height":
				if !contains(strconv.FormatFloat(float64(user.Height), 'f', 2, 32), c.Keyword) {
					match = false
				}
			case "married":
				if (strings.ToLower(c.Keyword) == "true" && !user.Married) || (strings.ToLower(c.Keyword) == "false" && user.Married) {
					match = false
				}
			default:
				match = false
			}

			if !match {
				break
			}
		}

		if match {
			result = append(result, user)
		}
	}

	return result
}
