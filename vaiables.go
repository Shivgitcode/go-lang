package main

import "fmt"

func main(){

	// declaring variables in go

	// strings

	var nameOne string="mario"
	var nameTwo="luigi"
	var nameThree string

	fmt.Println(nameOne,nameThree,nameTwo)

	nameOne="peach"
	nameThree="browser"

	fmt.Println(nameOne,nameTwo)

	// another syntax of declaring 

	nameFour := "yoshi"

	fmt.Println(nameFour)

}