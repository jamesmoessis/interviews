package main

import "strings"

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	currentRow := 0
	increasing := true
	for _, c := range s {
		rows[currentRow].WriteRune(c)
		if increasing {
			currentRow++
		} else {
			currentRow--
		}

		if currentRow == 0 || currentRow == len(rows)-1 {
			increasing = !increasing
		}
	}

	var sb strings.Builder
	for _, row := range rows {
		sb.WriteString(row.String())
	}
	return sb.String()
}
