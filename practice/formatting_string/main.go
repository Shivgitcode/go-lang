package main 

import "fmt"

func main(){
	fmt.Print("shivansh aggarwal\n")
	fmt.Print("Abhiraj Garg\n")

	name:="shivansh aggarwal"
	var age int=21
	var score float64=9.8


	fmt.Printf("My name is %v and my age is %v and my cgpa is %0.1f\n",name,age,score)
	
	fmt.Printf("type of age variable %T",age)

	fmt.Printf("type of age is %q",name)

	var intro string=fmt.Sprintf("My name is %v and my age is %v and my cgpa is %0.1f\n",name,age,score)

	fmt.Println(intro)
}