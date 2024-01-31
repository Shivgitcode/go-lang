package main 

import "fmt"

func greet(name string){
	fmt.Println(name)
}

func square(x int) int{
	return x*x
}

func PI() float64{
	return 3.14
}

func addAndSubtract(x int,y int)(int,int){
	return x+y,x-y
}