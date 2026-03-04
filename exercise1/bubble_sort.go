package exercise1

func BubbleSort(arr []int){
	n := len(arr)

// Vòng lặp chạy từ đầu đến kế cuối
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-1 -i; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
    	}
		if !swapped{
			break
		}
	}
}
