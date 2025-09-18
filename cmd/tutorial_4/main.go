package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "rèsumè"
	fmt.Println(myString)
	var indexed = myString[1]
	fmt.Println(indexed)
	fmt.Printf("%v %T\n", indexed, indexed)
	for ind, val := range myString {
		fmt.Println(ind, val)
	}
	fmt.Printf("Length: %v\n", len(myString))

	var myString2 = []rune("rèsumè")
	var indexed2 = myString2[1]
	fmt.Println(indexed2)
	fmt.Printf("%v %T\n", indexed2, indexed2)
	for ind, val := range myString2 {
		fmt.Println(ind, val)
	}
	fmt.Printf("Length: %v\n", len(myString2))

	var rune1 = 'a'
	fmt.Printf("\nrune1= %v", rune1)

	var strSlice = []string{"a", "n", "a", "c", "o", "n", "d", "a"}
	var catstr = ""
	for i := range strSlice {
		catstr += strSlice[i]
	}
	fmt.Printf("\n%v", catstr)
	// catstr[5]='a' // strings are immutable

	var strSlice2 = []string{"p", "y", "t", "h", "o", "n"}
	var strBuilder strings.Builder
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i])
	}
	var catstr2 = strBuilder.String()
	fmt.Printf("\n%v", catstr2)
}
