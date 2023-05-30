package src

// RemoveElement https://leetcode.cn/problems/remove-element/

func RemoveElement(nums []int, val int) int {
	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[low] = nums[fast]
			low++
		}
	}
	return low
}
