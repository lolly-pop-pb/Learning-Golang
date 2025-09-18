package main

import "fmt"

func main() {
	fmt.Println("Hi")
	arr1 := [...]int32{3, 6, 9}
	fmt.Println(arr1)

	var slice1 []int32 = []int32{4, 5, 6}
	fmt.Println(slice1)
	fmt.Printf("The length is %v and the capacity is %v\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 7)
	fmt.Println(slice1)
	fmt.Printf("The length is %v and the capacity is %v\n", len(slice1), cap(slice1))
	// fmt.Println(slice1[5])

	var slice2 []int32 = []int32{8, 9}
	fmt.Println(slice2)
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println(slice1[5])

	var slice3 []int32 = make([]int32, 3, 8)
	fmt.Println(slice3)
	fmt.Printf("The length is %v and the capacity is %v\n", len(slice3), cap(slice3))

}
