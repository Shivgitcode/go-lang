package main 

import "fmt"


func updateName(x string){
	x="wedge"
}

func updateMenu(x map[string]float64){
	x["coffee"]=2.99
}

func main (){
	// group A types -> strings,ints,bools,floats,arrays,structs 
	 name:="tifa"

	 updateName(name)

	 fmt.Println(name)

	// group B types -> slices,maps,functions

	menu:=map[string]float64{
		"pie":5.95,
		"ice cream":3.99,
	}


	updateMenu(menu)

	fmt.Println(menu)
}