package main

import (
	"fmt"
)

func main(){
	i:=5;

	for i<10{
		fmt.Printf("%v in loop\n",i)
		i++;
	} // this is a loop which is like for loop 

	for j:=0;j<5;j++{
		fmt.Printf("%v in loop\n",j)

	}

	var animeList[4] string=[4]string{"blue lock","fire force","The misfit of the demon academy","naruto"}

	for index,anime:=range animeList{
		fmt.Printf("%v. %v\n",index+1,anime)
	}


	for _,anime:=range animeList{
		fmt.Printf("%v\n",anime)
	}


	
}