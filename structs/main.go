package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string,r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input,err:=r.ReadString('\n')

	return strings.TrimSpace(input),err

}

func createBill() bill{
	reader:=bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name,_:=reader.ReadString('\n')
	// name=strings.TrimSpace(name)

	name,_:=getInput("Create a new bill name: ",reader)

	b:=newBill(name)
	fmt.Println("Created the bill - ",b.name)

	return b
}

func promptOptons(b bill){
	reader:=bufio.NewReader(os.Stdin)

	// name,_:=getInput("Create a new bill name: ",reader)
	opt,_:=getInput("Choose option (a-add item,s-save bill,t-add tip): ",reader)
	// fmt.Println(opt)
	switch opt{
	case "a":
		name,_:=getInput("Item name: ",reader)
		price,_:=getInput("Item price: ",reader)

		// fmt.Println("you chose a")
		fmt.Println(name,price)
	case "t":
		tip,_:=getInput("Enter tip amount ($): ",reader)
		fmt.Println("you chose t",tip)
	case "s":
		fmt.Println("you chose s")
	default:
		fmt.Println("that was not a valid option...")
		// promptOptions(b)
	}


}


func main(){
	//  mybill:=newBill("mario's bill")
	 
	// mybill.addItem("onion soup",4.50)
	// mybill.addItem("veg pie",8.95)
	// mybill.addItem("double latte",3.22)
	//  mybill.updateTip(10)


	//  fmt.Println(mybill.format())

	mybill :=createBill()

	fmt.Println(mybill)



}