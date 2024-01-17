package main

import "fmt"


// someName:="shivansh"

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

	nameFour := "yoshi" // replacement of var in the above cases we cant use it outside the function 

	fmt.Println(nameFour)
	

	// integers

	var ageOne int=20
	var ageTwo =30
	var ageThree int 

	ageFour:=40

	fmt.Println(ageOne)
	fmt.Println(ageTwo)
	fmt.Println(ageThree)
	fmt.Println(ageFour)

	//bits & memory : go lets us declare how much should a data take inside our memory


	// var numOne int8=25
	// var numTwo int8=-128
	// var numThree uint8=255


	// float

	var scoreOne float32=25.98 // we need to specify the bits in declaring floats
	var scoreTwo float64=4082390459023.7
	scoreThree :=1.5

	fmt.Println(scoreOne,scoreThree,scoreTwo)

}