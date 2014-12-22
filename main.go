package main

import "fmt"

func sort(unsorted []int) []int {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(unsorted)-1; i++ {
			if unsorted[i+1] < unsorted[i] {
				temp := unsorted[i+1]
				unsorted[i+1] = unsorted[i]
				unsorted[i] = temp
				changed = true
			}
		}
	}
	return unsorted
}

func main() {
	unsorted := []int{5, 9, 1, 4, 2}
	fmt.Print("Unsorted list ", unsorted)
	foo := sort(unsorted)
	fmt.Println(foo)
}
