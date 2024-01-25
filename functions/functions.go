package main

import(
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v\n",n);

}
func sayBye(n string){
	fmt.Printf("Goodbye %v\n",n)
}

func cycleNames(n []string,f func(string)){
	for _,v:=range n{
		f(v)
	}

}

func circleArea(r float64) float64{
	return math.Pi*r*r
}

func getInitials(n string)(string,string){
	s:=strings.ToUpper(n)
	names :=strings.Split(s," ")
	var initials []string;
	for _,v :=range names{
		initials=append(initials,v[:1])
	}

	if len(initials)>1{
		return initials[0],initials[1]
	}

	return initials[0],"_"


}

func main(){
	// sayGreeting("mario")
	// sayGreeting("luigi")

	// sayBye("mario")

	// n := []string{"shivansh","Tanishq","Abhinav","Mann","Viren","Nischay"}

	// cycleNames(n,sayBye)

	// a1:=circleArea(10.5)

	// fmt.Println()
	// fmt.Printf("%0.2f",a1)

	fn1,sn1:=getInitials("tifa lockhart")

	fmt.Println(fn1,sn1)


}