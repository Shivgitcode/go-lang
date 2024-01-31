package main

import "fmt"

func updateName(x *string){
	*x="Ronaldo"
}

func main(){
	name:="Messi"
	// updateName(name)
	fmt.Println(name)

	fmt.Println(&name)

	pname:=&name;

	updateName(pname)
	fmt.Println(*pname)
	fmt.Println(name)

}