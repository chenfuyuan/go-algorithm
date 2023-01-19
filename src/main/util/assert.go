package util

func IsArraySortAsc(array []int) bool {
	if array == nil {
		return false
	}
	length := len(array)
	if length == 1 {
		return true
	}
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}
