package GenerateParentheses



func GenerateParenthesis(n int) []string {
	var res []string
	add(&res, "", 0, 0, n)
	return res
}

func add(arr *[]string, s string, l, r, n int) {
	if l > n || r > n || r > l {
		return
	}
	if l == n && r == n {
		*arr = append(*arr, s)
		return
	}
	add(arr, s+"(", l+1, r, n)

	add(arr, s+")", l, r+1, n)
}