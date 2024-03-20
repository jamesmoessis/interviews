package main

// I'm making this harder by returning all pairs.
func twoSum(nums []int, target int) [][2]int {
	n := len(nums)
	m := make(map[int][]int, n*2)

	for i, num := range nums {
		indexes, ok := m[num]
		if !ok {
			m[num] = []int{i}
		} else {
			m[num] = append(indexes, i)
		}
	}

	pairs := make([][2]int, 0)

	for i, num := range nums {
		balancer := target - num
		indexes, ok := m[balancer]
		if !ok {
			continue
		}
		for foundBalancerIndex := range indexes {
			var pair [2]int
			if i < foundBalancerIndex {
				pair = [2]int{i, foundBalancerIndex}
			} else {
				pair = [2]int{foundBalancerIndex, i}
			}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}
