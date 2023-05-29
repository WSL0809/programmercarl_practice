package src

// Search https://leetcode.cn/problems/binary-search/
func Search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	//mid := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
