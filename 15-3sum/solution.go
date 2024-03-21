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
		for _, foundBalancerIndex := range indexes {
			// Only take the case where lower index first, otherwise we will have dupes
			if i < foundBalancerIndex {
				pairs = append(pairs, [2]int{num, balancer})
			}
		}
	}
	return pairs
}

// todo debug this
func twoSumEasy(nums []int, target int) []int {
	// maps value to index
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		balancer := target - num
		balancerIndex, ok := m[balancer]
		if ok && balancerIndex != i {
			return []int{i, balancerIndex}
		}
	}
	return nil
}
