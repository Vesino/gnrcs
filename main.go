package main

import "fmt"

func splitSlice[T any](s []T) ([]T, []T) {
	if len(s) == 0 {
		var zeroVAl T
		return []T{zeroVAl}, []T{zeroVAl}
	}
	mid := len(s)/2
	return s[:mid], s[mid:]
}

func main() {
	myEmails := []string{"sfb@udjf.com", "jsahfb@jdikf.com", "shjnk@kd.com"}
	fmt.Println(splitSlice(myEmails))
	myZeroEmails := []string{}
	fmt.Println(splitSlice(myZeroEmails))

	myNumbers := []int{5273, 3948, 78, 893, 9}
	fmt.Println(splitSlice(myNumbers))
}
