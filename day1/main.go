package main

import "fmt"

func main() {

	x := []int{3, 4, 2, 1, 3, 3}
	y := []int{4, 3, 5, 3, 9, 3}

	answer := calculateDistance(x, y)

	fmt.Printf("Answer = %d\n", answer)
}

func bubbleSort(unsortedList []int) []int {
	for i := 0; i < len(unsortedList)-1; i++ {
		for j := 0; j < len(unsortedList)-i-1; j++ {
			if unsortedList[j] > unsortedList[j+1] {
				unsortedList[j], unsortedList[j+1] = unsortedList[j+1], unsortedList[j]
			}
		}
	}

	return unsortedList
}

func calculateDistance(l1 []int, l2 []int) int {
	sum := 0

	orderedList1 := bubbleSort(l1)
	orderedList2 := bubbleSort(l2)

	for i := range orderedList1 {
		r := (orderedList1[i] - orderedList2[i])
		if r < 0 {
			r = r * -1
		}

		fmt.Printf("%d\n", r)

		sum += r
	}

	return sum
}
