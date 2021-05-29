package main

import "fmt"

func double(s []int) {
	s[len(s)-1] = 2222
}

func main() {
	s := []int{1, 2, 3}
	s = append(s, 1000)
	double(s)
	fmt.Println(s, len(s)) // prints [1 2 3] 3
}