package main 

import "fmt"

func main(){
	fmt.Println("hello world")

	// variables types in go lang

	// string type variabels

	var stringOne string="shivansh"
	var stringTwo="aggarwal"
	stringThree:="Abhiraj Garg"

	fmt.Println(stringOne)
	fmt.Println(stringTwo)
	fmt.Println(stringThree)

	// interger datatypes in go lang

	var numberOne int=40
	var numberTwo int32=50 // we can also declare how much space a variable should take but it is not compulsory to declare
	numberThree:=70;

	// float datatypes in gol lang 

	var scoreOne float32=32.5

	var scoreThree uint32=70

	// scoreOne=50.3
	// but it is compulsory to declare the bit for float type values
	scoreTwo:=70.6

	fmt.Println(scoreTwo) //go lang also supports type inferrence 
	fmt .Println(scoreThree)
	fmt.Println(scoreOne)


	
	


	fmt.Println(numberThree)


	fmt.Println(numberOne)
	fmt.Println(numberTwo)

}