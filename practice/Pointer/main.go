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

	pname:=&name; //ampersant(&) is used to access memory location of the variable and

	updateName(pname)
	fmt.Println(*pname)
	fmt.Println(name)

}