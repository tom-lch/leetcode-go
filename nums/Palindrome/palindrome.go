package Palindrome



func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var num int
	y := x
	for y > 0 {
		num *= 10
		num += y % 10
		y /= 10
	}
	return num == x
}