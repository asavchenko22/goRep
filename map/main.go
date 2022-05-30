package main

import "fmt"

func main() {
	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	numbersOfSymb := map[rune]int{}
	for _, tmp := range str {
		numbersOfSymb[tmp] += 1
	}
	for key, value := range numbersOfSymb {
		fmt.Printf("%c", key)
		fmt.Print(" ", value, "\n")
	}
}
