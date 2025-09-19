package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

// func canMakeIt(e gasEngine, miles uint8){
// 	if miles<=e.milesLeft(){
// 		fmt.Println("You can make it there!")
// 	}else{
// 		fmt.Println("Need to fuel up first!")
// 	}
// }

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine1 gasEngine = gasEngine{25, 1}
	fmt.Println(myEngine1.mpg, myEngine1.gallons)
	fmt.Printf("The miles left in tank: %v\n", myEngine1.milesLeft())
	var myEngine2 electricEngine = electricEngine{25, 10}
	fmt.Println(myEngine2.mpkwh, myEngine2.kwh)
	fmt.Printf("The miles left in tank: %v\n", myEngine2.milesLeft())
	canMakeIt(myEngine1, 50)
	canMakeIt(myEngine2, 50)
}