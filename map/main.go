package main 

import "fmt"


func main(){

	menu:=map[string]float64{
		"soup":4.99,
		"pie":7.99,
		"salad":6.99,
		"toffee pudding":3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through maps 

	for k,y :=range menu {
		fmt.Println(k,"-",y)
	}

	// ints as key type

	phonebook:=map[int]string{
		234253223:"mario",
		242345342:"luigi",
		834834245:"peach",
	}

	for k,y:=range phonebook{
		fmt.Println(k,"-",y)
	}
}