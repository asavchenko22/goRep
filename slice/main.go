package main

import "fmt"

func main() {
	sl := []string{"Andrii", "Anna", "Ivan", "Vasyl", "Solomija", "Wirsawia", "Andrej", "Bogdan"}
	sort(sl)
	fmt.Println(sl)
}

func sort(sl []string) {
	for j := 0; j < len(sl)-1; j++ {
		current := sl[j]
		next := sl[j+1]
		if current < next {
			sl[j+1] = current
			sl[j] = next
		}
	}
	fmt.Println(true)
}
