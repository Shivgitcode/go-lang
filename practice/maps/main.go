package main

import "fmt"

func main(){
	//maps are just like objects in javascript

	book:=map[string]int{
		"ikigai":180,
		"48 laws of power":180,
		"subtle art of not giving a f***":250,
	}

	
	fmt.Println(book["subtle art of not giving a f***"])


	fmt.Println(book["ikigai"])

	for key,value:=range book{
		// fmt.Println(key,value)
		fmt.Printf("Name: %v\nPrice: %v\n\n\n",key,value)
	}

	

	// fmt.Println(book)
}