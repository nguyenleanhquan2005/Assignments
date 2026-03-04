package exercise1

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	result := make([]int, 0 , size)

	i, j := 0, 0
	for i < len(left) && j < len(right){
		if left[i] < right[j]{
			result = append(result, left[i])
			i++
		}else{
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result 



}