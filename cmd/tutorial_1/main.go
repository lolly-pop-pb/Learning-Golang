package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello World!")
	// var intNum int = 32767
	// intNum = intNum + 1
	// fmt.Println(intNum)
	// var floatNum float64 = 12345678.9
	// fmt.Println(floatNum)

	// var a float32 = 10.1
	// var b int32 =2
	// var result float32 = a + float32(b)
	// fmt.Println(result)

	// var x int = 3
	// var y int =2
	// fmt.Println(x/y)
	// fmt.Println(x%y)
	// var str1 string = "Hi"
	// fmt.Println(str1)
	// fmt.Println(len(str1))
	// fmt.Println(len("α"))
	fmt.Println(utf8.RuneCountInString("α"))

	// var rune1 rune = 'a'
	// fmt.Println(rune1)

	// var boole bool=false
	// fmt.Println(boole)

	var intNum3 int
	fmt.Println(intNum3)

	var myText = "test"
	fmt.Println(myText)

	myText2 := "test2"
	fmt.Println(myText2)

	// initialize multiple variables at once
	var x1, x2 int = 1, 2
	var y1, y2 = 4, 5
	z1, z2 := 7, 8

	fmt.Println(x1, x2)
	fmt.Println(y1, y2)
	fmt.Println(z1, z2)

	const myConst int = 5
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
