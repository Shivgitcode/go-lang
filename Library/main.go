package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){

	greeting :="hello everyone i am shivansh aggarwal"
	// fmt.Println(strings.Contains(greeting,"hello"))
	// fmt.Println(strings.ReplaceAll(greeting,"hello","hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting,"ll"))

	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45,20,35,30,75,60,50,25}

	sort.Ints(ages)
	fmt.Println(ages)

	index:=sort.SearchInts(ages,30)

	fmt.Println(index)


}