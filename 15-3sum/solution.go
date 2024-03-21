package main

import (
	"slices"
	"sort"
)

// threeSum solves the 3-sum problem by reducing it to n 2-sum problems
// So if we have an O(n) 2-sum we can get an O(n^2) 3-sum.
func threeSum(nums []int) [][]int {
	m := make(map[int][]int)
	for i, num := range nums {
		indexes, ok := m[num]
		if !ok {
			m[num] = []int{i}
		} else {
			m[num] = append(indexes, i)
		}
	}

	solSet := make(map[[3]int]bool)

	for i, num := range nums {
		target := -1 * num
		twoSums := twoSum(removeIndex(nums, i), target)
		for _, pair := range twoSums {
			arr := [3]int{num, pair[0], pair[1]}
			slices.Sort(arr[:])
			solSet[arr] = true
		}
	}
	sol := make([][]int, 0, len(solSet))
	for k := range solSet {
		sol = append(sol, k[:])
	}
	return sol
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// I'm making this harder by returning all pairs.
// Not quite O(n), actually o(n^2) -- is it possible to get this to O(n)?
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

func twoSumEasy(nums []int, target int) []int {
	// maps value to index
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		balancer := target - num
		balancerIndex, ok := m[balancer]
		if ok && balancerIndex != i {
			return []int{balancerIndex, i}
		}
		m[num] = i
	}

	return nil
}

// Very smart solution I found on leetcode solutions
func threeSumSmart(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
