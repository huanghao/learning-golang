package main

import "fmt"

func main() {
	var s = make([]int, 3, 100)
	fmt.Println(s)
	var s2 = make([]int, 0, 10)
	s2 = append(s2, 1, 2, 3)
	s = append(s, s2[2:5]...)
	fmt.Println(s)
}
