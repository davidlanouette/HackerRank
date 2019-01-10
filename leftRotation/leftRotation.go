package main

import "fmt"

func main() {

	n := int32(5)
	d := int32(4)

	a := []int32{1, 2, 3, 4, 5}
	b := a

	fmt.Print("starting ")
	printSlice(a)

	for i := int32(0); i < d; i++ {
		b = rotate(b, n)
		// fmt.Println(" - ", b)
	}
	printSlice(b)

	c := rotate2(a, n, d)
	printSlice(c)
}

func printSlice(a []int32) {
	for i, n := range a {
		fmt.Print(n)
		if i != len(a) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func rotate2(a []int32, n, d int32) []int32 {
	result := make([]int32, n, n)

	for i := int32(0); i < n; i++ {
		// fmt.Println(i, i+n, i+n+d, i+d, i+d/n, (i+d)/n, (i+d+n)%n)
		newIndex := (i + d + n) % n
		// fmt.Println("- ", i, newIndex, a[newIndex])
		result[i] = a[newIndex]
		// fmt.Println("-- ", result)
	}
	return result
}

func rotate(a []int32, n int32) []int32 {
	result := make([]int32, n, n)

	result[n-1] = a[0]
	for i := int32(0); i < n-1; i++ {
		// fmt.Println(i, a[i], a[i+1])
		result[i] = a[i+1]
		// fmt.Println(" -- ", result)
	}
	return result
}
