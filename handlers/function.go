package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isValidField(field string) bool {
	validFields := map[string]bool{
		"fname":   true,
		"city":    true,
		"phone":   true,
		"height":  true,
		"married": true,
	}

	return validFields[field]
}

func validateUserID(id int32) error {
	if id <= 0 {
		return status.Errorf(codes.InvalidArgument, "Invalid user ID: %d", id)
	}
	return nil
}

func validateUserIDs(ids []int32) error {
	for _, id := range ids {
		if err := validateUserID(id); err != nil {
			return err
		}
	}

	return nil
}
