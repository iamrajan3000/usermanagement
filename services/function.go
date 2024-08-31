package services

import "strings"

func contains(source, keyword string) bool {
	return strings.Contains(strings.ToLower(source), strings.ToLower(keyword))
}
