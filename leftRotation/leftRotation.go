package main

import "fmt"

func main() {

	n := 5
	d := 5

	a := []int{1, 2, 3, 4, 5}

	fmt.Println("starting ", a)

	for i := 0; i < d; i++ {
		a = rotate(a, n)
		fmt.Println(" - ", a)
	}

	fmt.Println("final array: ", a)
}

func rotate(a []int, n int) []int {
	result := make([]int, n, n)

	result[n-1] = a[0]
	for i := 0; i < n-1; i++ {
		// fmt.Println(i, a[i], a[i+1])
		result[i] = a[i+1]
		// fmt.Println(" -- ", result)
	}
	return result
}
