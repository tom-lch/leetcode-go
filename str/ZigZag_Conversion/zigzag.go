package ZigZag_Conversion

import "strings"

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans := make([]string, numRows)
	curRow := 0
	down := 0
	for i := 0; i < len(s); i++ {
		ans[curRow] += string(s[i])
		if curRow == 0 || curRow == numRows-1 {
			down ^= 1
		}
		if down == 1 {
			curRow++
		} else {
			curRow--
		}
	}
	return strings.Join(ans, "")
}
