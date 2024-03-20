package main

import "fmt"

func main() {
	m := make(map[int]int)
	a := [...]int{1, 2, 3, 4}

	for i := range a {
		m[i]++
	}
	fmt.Println(m)

	var b []int
	b = append(b, 1)
	fmt.Println(b)
}
