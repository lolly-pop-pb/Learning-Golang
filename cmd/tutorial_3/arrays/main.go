package main

import "fmt"

func main() {
	var arr1 [3]int32

	arr1[0] = 23
	arr1[1] = 24
	arr1[2] = 25
	fmt.Println(arr1)
	fmt.Println(arr1[0])
	fmt.Println(arr1[0:2])

	fmt.Println(&arr1[0])
	fmt.Println(&arr1[1])
	fmt.Println(&arr1[2])

	var arr2 [4]int32 = [4]int32{-1, -2, -3, -4}
	fmt.Println(arr2)

	arr3 := [4]int32{0, 3, 6, 9}
	fmt.Println(arr3)

	arr4 := [...]int32{4, 6, 8}
	fmt.Println(arr4)

}
