package main

import "fmt"

type players struct{
	name string
	playerNo int 
	playerTeam string
}

func main(){

	player1:=players{
		name:"Ronaldo",
		playerNo:7,
		playerTeam:"Al nassar",
	}

	fmt.Println(player1)
	fmt.Println(player1.name)

	for i:=0;i<4;i++{
		fmt.Println(player1.name)
	}


}