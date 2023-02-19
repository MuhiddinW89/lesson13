package main

import (
	"asd/controller"
	"fmt"
)

func main() {
	controller.GenereteUsr(51)
	// users, err := controller.GetListUsr(models.GetListRequset{Offset: 10, Limit: 10})	
	// for _, v := range users{
	// 	fmt.Println(v)
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(controller.Pagination(7))

	// fmt.Println(controller.FindUsrByName("be"))

	fmt.Println(controller.BirthDateFilter("2000-01-01","2010-01-02"))
}	