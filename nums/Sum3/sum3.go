package Sum3

import "sort"


func Sum3(nums []int) [][]int {
	// 使用枚举
	var ans [][]int
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < n; first ++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n-1
		target := -nums[first]
		for second := first+1; second < n; second++ {
			if second > first+1 && nums[second]==nums[second-1] {
				continue
			}
			// third指针一直想做移动，一直到找到目标
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
