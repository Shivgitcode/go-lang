package main

import "fmt"

func main (){
	age:=45;

	fmt.Println(age<=50)
	fmt.Println(age>=50)
	fmt.Println(age==45)
	fmt.Println(age!=50)

	name:=[]string{"shivansh","Abhinav","viren","tanishq"}

	for index,value:=range name{
		if index==1{
			fmt.Println("Continuing at pos",index)
			continue
		}
		fmt.Printf("the value at pos %v is %v\n",index,value)
	}

	// if age<30{
	// 	fmt.Println("age is less than 30")

	// }else if age<40{
	// 	fmt.Println("age is less than 40")
	// }else{
	// 	fmt.Println("age is not less than 45")
	// }

}