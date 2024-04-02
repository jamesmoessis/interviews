package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	ans := sort.Search(len(a), func(i int) bool {
		return a[i] >= 4
	})
	fmt.Println(ans)
}
