package main 

import "fmt"

func main (){
	// var ages [3]int=[3]int{20,25,30} length of arrays cannot change
	var ages=[3]int{20,25,30}

	name:=[4]string{"yoshi","mario","peach","bowser"}
	name[1]="luigi"

	fmt.Println(ages,len(ages))
	fmt.Println(name,len(name))

	//slices (use arrays under the hood)

	var scores=[]int{100,50,60} // this is a slice as we didn't specify the length of the array
	scores[2]=25
	scores=append(scores,85)

	//slice ranges
	rangeOne :=name[1:3]
	rangeTwo :=name[2:]
	rangeThree:=name[:3]

	fmt.Println(scores,len(scores))

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)



}