package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxInt = 8

func main() {
	var intArr [maxInt]int
	intArr = fillingArr(intArr)
	checkArr(intArr)
	printArr(intArr)
}

func checkArr(arr [maxInt]int) {
	output := false
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		for j := 0; j < len(arr); j++ {
			y := arr[j]
			if x == y && i != j {
				output = true
			}
		}

	}
	fmt.Println(output)
}

func printArr(arr [maxInt]int) {
	for _, number := range arr {
		fmt.Print(number, " ")
	}
}

func fillingArr(arr [maxInt]int) [maxInt]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxInt; i++ {
		arr[i] = rand.Intn(maxInt)
	}
	return arr
}

//func fillingArr(arr [...]int) {
//	for i := minInt; i < maxInt; i++ {
//		arr[i] = rand.Intn(maxInt)
//	}
//}
