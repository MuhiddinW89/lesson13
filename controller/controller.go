package controller

import (
	"asd/models"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)


var Users []models.User

func CreateUsr(data models.User){
	Users = append(Users, data)
}

func GetListUsr(req models.GetListRequset) ([]models.User,error) {
	if req.Limit + req.Offset > len(Users) {
		return []models.User{}, errors.New("Out of range")
	}
	return Users[req.Offset:req.Limit+req.Offset],nil
}

func UpdateUsr(id int, name string, surname string){
	for i, v := range Users {
		if id == v.Id{
			Users[i].Name = name
			Users[i].Surname = surname
			fmt.Println(v.Id, "user updated")
		}
	}	
}

func DleteUsr(id int){
	slice := []models.User{}
	for _, v := range Users {
		if id != v.Id{
			slice = append(slice, v)
		}
	}
	Users = slice
}

func GenereteUsr(count int){
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id: i+1,
			Name: faker.FirstName(),
			Surname: faker.LastName(),
			BirthDate: faker.Date(),
		})
	}
}

func Pagination(page int) ([]models.User, error){
	fmt.Println(Users)
	pages := len(Users)/10
	if len(Users)%10 == 1  {
		pages = len(Users)/10 + 1
	}
	if page > pages {
		return []models.User{}, errors.New("page out of range")
	}
	fmt.Println("Pages: ", pages)
	return Users[page*10-10 : page*10], nil
}

func FindUsrByName(name string)([]models.User){
	slice := []models.User{}
	for i := 0; i < len(Users); i++ {
	  if strings.Contains( strings.ToLower(Users[i].Name + Users[i].Surname),  strings.ToLower(name)  ){
			slice = append(slice, Users[i])
	  }
	}
	return slice
}

func BirthDateFilter(fromDate, toDate string) []models.User{
	slice := []models.User{}
	strart, _ := time.Parse("2006-01-02", fromDate)
	end, _ := time.Parse("2006-01-02", toDate)
	for _, v := range Users{
		usrdate, _ := time.Parse("2006-01-02", v.BirthDate)
		if usrdate.After(strart) && usrdate.Before(end) {
			slice = append(slice, v)
		}
	}
	return slice
}

