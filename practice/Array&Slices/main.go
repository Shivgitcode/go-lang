package main 

import "fmt"

func main(){
	// arrays
	var name[3] string=[3]string{"shivansh","aggarwal","hello"}

	var score=[3]float32{30.2,50.3,60.5}

	age:=[3]int{18,21,15}

	fmt.Println(age)

	fmt.Println(score[0])
	fmt.Println(name)

	// slices

	score2:=[]int{10,290,50,70}
	fmt.Println(score2)

	var animeList[4] string=[4]string{"blue lock","fire force","The misfit of the demon academy","naruto"}
	fmt.Println(animeList)
	// slicing a slice 
	subscore:=score2[0:2]
	subscore2:=score2[0:]
	subscore3:=score2[1:]
	subscore2=append(subscore2,90)
	fmt.Println(subscore2)
	fmt.Println(subscore)
	fmt.Println(subscore2)
	fmt.Println(subscore3)
}