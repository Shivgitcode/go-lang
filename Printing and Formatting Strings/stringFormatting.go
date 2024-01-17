package main

import "fmt"

func main(){
	age:=20
	name:="shivansh"

	fmt.Print("hello\n");
	fmt.Print("world!\n");
	fmt.Print("new line\n")


	fmt.Println("Hello, ")
	fmt.Println("World!")

	fmt.Println("my age is",age,"and my name is",name)

	// formatting string

	fmt.Printf("my age is %v and my name is %v",age,name)
	fmt.Printf("my age is %q and my name is %q",age,name) // here %q will find the variables and print them out in double quotes
	fmt.Printf("age is of type %T",age) // %T get us the type of the varaible we declared
	fmt.Printf("you scored %0.2f points!\n",255.55)

	//Sprintf (save formatted strings)

	var str=fmt.Sprintf("my age is %v and my name is %v\n",age,name)
	fmt.Println(str)
}