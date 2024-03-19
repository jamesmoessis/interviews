package main

import "strings"

var (
	M = [...]string{"", "M", "MM", "MMM"}
	C = [...]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X = [...]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I = [...]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	var sb strings.Builder
	sb.WriteString(M[num/1000])
	sb.WriteString(C[(num%1000)/100])
	sb.WriteString(X[(num%100)/10])
	sb.WriteString(I[num%10])
	return sb.String()
}