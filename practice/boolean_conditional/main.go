package main 


import (
	"fmt"
)


func main(){
	// var age int;

	// fmt.Println("Enter your age:")
	// fmt.Scan(&age)

	// if age>=18{
	// 	fmt.Println("you are eligible to vote")
	// }else{
	// 	fmt.Println("you are not eligible to vote")
	// }

	ages:=[]int{10,20,15,30}

	for _,age:=range ages{
		if age<18{
			continue
		}
		fmt.Println(age)
	}


}