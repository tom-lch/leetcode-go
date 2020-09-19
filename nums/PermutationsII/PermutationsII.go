package PermutationsII

import "sort"

var ans [][]int
func PermuteUnique(nums []int) [][]int {
	defer func(){
		ans = nil
	}()
	sort.Ints(nums)
	addNums(nums, 0, len(nums)-1)
	return ans
}

func addNums(nums []int, left, right int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	if left == right {
		ans = append(ans, arr)
		return
	} else {
		for i := left; i <= right; i++ {
			if i != left && arr[left] == arr[i] {
				continue
			}
			arr[left], arr[i] = arr[i], arr[left]
			addNums(arr, left+1, right)

		}
	}
}
