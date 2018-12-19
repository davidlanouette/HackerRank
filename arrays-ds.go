package main

/*
https://www.hackerrank.com/challenges/arrays-ds/problem

An array is a type of data structure that stores elements of the same type in a contiguous block of memory. In an array, A, of size N, each memory location has some unique index, i (where 0 <= i < N), that can be referenced as A[i].

Given an array, A, of N integers, print each element in reverse order as a single line of space-separated integers.

Input Format

The first line contains an integer, N (the number of integers in A).
The second line contains N space-separated integers describing A.

Constraints

1 <= N <= 10^3
1 <= A[i] <= 10^4, where A[i] is the i-th integer in A.

Output Format

Print all N integers in A in reverse order as a single line of space-separated integers.

Sample Input 1:
4
1 4 3 2

Sample Output 1:
2 3 4 1

*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the reverseArray function below.
func reverseArray(a []int32) []int32 {
	return a
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := reverseArray(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
