package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerinfo owner
	owner
	int
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)
	// output= 0 0
	var myEngine2 gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine = gasEngine{30, 15, owner{"Alex"}, 10}
	// fmt.Println(myEngine3.mpg, myEngine3.gallons, myEngine3.ownerinfo.name)
	fmt.Println(myEngine3.mpg, myEngine3.gallons, myEngine3.name)
	myEngine3.mpg = 20
	fmt.Println(myEngine3.mpg, myEngine3.gallons)
	// anonymous struct define and initialize in the same location
	var myEngine4 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine4.mpg, myEngine4.gallons)
	// but it is not reusable
	// if you need to make another struct, you will have to rewrite the definition
	var myEngine5 = struct {
		mpg     uint8
		gallons uint8
	}{30, 20}
	fmt.Println(myEngine5.mpg, myEngine5.gallons)
}
