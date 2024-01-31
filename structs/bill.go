package main

import "fmt"

type bill struct{
	name string 
	items map[string]float64
	tip float64
}


// make new bills

func newBill(name string) bill{
	b :=bill{
		name:name,
		items:map[string]float64{},
		tip:0,
	}

	return b;

}

// format the bill 

func (b *bill) format() string{
	fs :="bill breakdown :\n"
	var total float64=0

	//list items;
	for k,v:=range b.items{
		fs +=fmt.Sprintf("%-25v ...$%v\n",k+":",v) // -25 signifies that the key takes a 25 character long value but the key is only 4 characters long so it will fill rest of the space with empty spaces and also -25 signifies that we have to give space to the right and 25 signifies space to the left
		total+=v;
	}

	// add tip

	fs +=fmt.Sprintf("%-25v ...$%0.2f\n","tip:",b.tip)

	// total 

	fs +=fmt.Sprintf("%-25v ...$%0.2f","total:",total+b.tip)


	return fs

}

// update tip

func (b *bill) updateTip(tip float64){
	// (*b).tip=tip we can do this but 
	b.tip=tip //when we pass pointer struct go automatically dereference the pointer values
}

//add an item to the bill

func(b *bill) addItem(name string,price float64){
	b.items[name]=price
}