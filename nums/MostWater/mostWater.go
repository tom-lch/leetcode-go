package MostWater


func MaxArea(height []int) int {
	var count int
	left, right := 0, len(height)-1
	for left < right {
		minH := min(height[left], height[right])
		if count < minH * (right-left) {
			count = minH * (right-left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
