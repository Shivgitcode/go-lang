package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){

	greeting:="Hello everyone , You are learning go programming language"

	fmt.Println(strings.Contains(greeting,"hello"))

	greeting2:=strings.ReplaceAll(greeting,"Hello","hi")

	fmt.Println(greeting2)

	greeting=strings.ToUpper(greeting)
	fmt.Println(greeting)

	fmt.Println(strings.Index(greeting,"YOU"))

	var numbers=[]int{30,50,100,1,2,60}
	sort.Ints(numbers)

	fmt.Println(numbers)


}