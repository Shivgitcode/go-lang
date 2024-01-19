package main

import(
	"fmt"
)

func main(){
	// x :=0

	// for x<5{
	// 	fmt.Println("value of x is:",x)
	// 	x++

	// }

	// for i:=0; i<5 ; i++{
	// 	fmt.Println("value of i is: ",i)

	// }

	names:=[]string{"mario","luigi","yoshi","peach"}

	for i:=0;i<len(names);i++{
		fmt.Println("value of name is : ",names[i])
	}

	for index,value:=range names{
		fmt.Printf("the postion at index %v is %v\n",index,value)
	}

	// if we don't  want to use index then we will run for loop like this


	for _,value:=range names{
		fmt.Printf("%v\n",value)
	}
}