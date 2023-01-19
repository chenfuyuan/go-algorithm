package binary_search

// BinarySearch 二分查找
// @params:
// nums: 升序数组,且无重复数字
// target: 所要查找的元素
func BinarySearch(nums []int, target int) int {
	leftIndex, rightIndex := 0, len(nums)-1
	for leftIndex <= rightIndex {
		checkIndex := (rightIndex + leftIndex) / 2
		checkElement := nums[checkIndex]
		if checkElement == target {
			return checkIndex
		}
		if target < checkElement {
			rightIndex = checkIndex - 1
		} else {
			leftIndex = checkIndex + 1
		}
	}
	return -1
}
