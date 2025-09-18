package main

import "fmt"

func main() {
	arr1 := [...]int32{1, 2, 3}
	fmt.Println(arr1)

	var arr2 [3]int32 = [3]int32{12, 13, 15}
	fmt.Println(arr2)

	for index, value := range arr2 {
		fmt.Printf("Index: %v Value: %v\n", index, value)
	}

	var slice1 []int32 = []int32{4, 5, 6}
	fmt.Println(slice1)

	var slice2 []int32 = make([]int32, 3, 8)
	fmt.Println(slice2)

	var map1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(map1)

	var map2 = map[string]uint8{"apple": 1, "banana": 2, "cherry": 3, "durian": 4}
	fmt.Println(map2)
	fmt.Println(map2["banana"])
	fmt.Println(map2["orange"])

	var freq, ok = map2["apple"]
	if ok {
		fmt.Printf("The freq is %v\n", freq)
	} else {
		fmt.Printf("Invalid\n")
	}

	for fruit, freq := range map2 {
		fmt.Printf("Fruit: %v Freq: %v\n", fruit, freq)
	}


	var map3 = map[string]uint8{"Adam": 23, "Bob": 24, "Jason": 25}
	fmt.Println(map3)
	delete(map3, "Adam")
	fmt.Println(map3)

	for name := range map3 {
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range map3 {
		fmt.Printf("Name: %v Age: %v\n", name, age)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	var j int = 10
	for {
		if j < 0 {
			break
		}
		fmt.Println(j)
		j = j - 1
	}

	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}
}
