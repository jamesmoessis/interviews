package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/?envType=list&envId=5l31jv7r
func search(nums []int, target int) int {
	offset := findOffset(&nums)
	return binarySearch(&nums, offset, target)
}

func findOffset(nums *[]int) int {
	// low, mid, high
	// if mid > mid + 1, then offset = mid + 1
	// if high < mid, then "start" of array in top half
	// if low > mid, then "start" of array in bottom half
	a := *nums
	low := 0
	high := len(a) - 1
	mid := 0
	if high == 0 {
		return 0
	}

	for low <= high {
		mid = (low + high) / 2
		if a[mid] > a[mid+1] {
			return mid + 1
		} else if mid != 0 && a[mid] < a[mid-1] {
			return mid
		} else if a[high] < a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}

func binarySearch(nums *[]int, offset, target int) int {
	low := 0
	n := len(*nums)
	high := n - 1
	mid := 0
	val := 0
	for low <= high {
		mid = (high + low) / 2
		val = (*nums)[(offset+mid)%n]
		if val < target {
			// target in top half
			low = mid + 1
		} else if val > target {
			// target in bottom half
			high = mid - 1
		} else {
			return (offset + mid) % n
		}
	}
	return -1
}
