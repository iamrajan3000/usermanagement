package models

type User struct {
	Id      int32
	Fname   string
	City    string
	Phone   uint64
	Height  float32
	Married bool
}

type SearchCriteria struct {
	Field   string
	Keyword string
}
