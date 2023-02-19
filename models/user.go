package models

type User struct{
	Id  int
	Name string
	Surname string
	BirthDate string
} 

type GetListRequset struct {
	Offset int
	Limit int
}