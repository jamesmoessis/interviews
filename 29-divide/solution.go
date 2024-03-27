package main

// how many times does divisor go into dividend?
// can find out by adding

func divide(dividend int, divisor int) int {
	// optimisations and corner cases
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}

	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	ans := 0
	built := divisor
	for built <= dividend {
		built += divisor
		ans++
	}
	if sign == 1 {
		return ans
	}
	return -ans
}
