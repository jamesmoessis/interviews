package main

func maxDepth(s string) int {
	n := len(s)
	if n <= 1 && s != "(" && s != ")" {
		return 0
	}

	maxDepth := 0
	currentDepth := 0

	for _, c := range s {
		if c == '(' {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		} else if c == ')' {
			currentDepth--
		}
	}
	return maxDepth
}
