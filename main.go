// package main

// import (
// 	"fmt"
// 	// Import package exercise1.
// 	// Đường dẫn này phụ thuộc vào tên module bạn đặt ở bước 3.
// 	// Giả sử tên module là "myproject"
// 	"app/exercise1"
// )

// func main() {
// 	fmt.Println("--- Bắt đầu chương trình ---")

// 	// 1. Tạo cây
// 	root := exercise1.CreateTree()

// 	// 2. In thử InOrder
// 	fmt.Println("InOrder traversal:")
// 	exercise1.InOrder(root)

// 	fmt.Println("InOrder traversal:")
// 	exercise1.PreOrder(root)

// 	// 3. Tìm số lớn nhất
// 	maxVal := exercise1.GetBiggestNumber(root)
// 	fmt.Printf("Biggest number is: %d\n", maxVal)

// 	// 4. Lấy mảng đã sắp xếp
// 	sortedList := exercise1.GetSortedNumbers(root)
// 	fmt.Println("sorted list:", sortedList)

// 	input := []int{64, 34, 25, 12, 22, 11, 90}
// 	fmt.Println("Before", input)
// 	exercise1.BubbleSort(input)
// 	fmt.Println("After", input)
// 	original := []int{10, 7, 8, 9, 1, 5}
// 	fmt.Printf("Mảng ban đầu: %v\n\n", original)

// 	// 2. Chạy thử BubbleSort
// 	arr1 := make([]int, len(original))
// 	copy(arr1, original)
// 	exercise1.BubbleSort(arr1)
// 	fmt.Println("BubbleSort:", arr1)

// 	// 3. Chạy thử MergeSort (hàm này trả về mảng mới)
// 	arr2 := make([]int, len(original))
// 	copy(arr2, original)
// 	resultMerge := exercise1.MergeSort(arr2)
// 	fmt.Println("MergeSort: ", resultMerge)

// 	// 4. Chạy thử QuickSort
// 	arr3 := make([]int, len(original))
// 	copy(arr3, original)
// 	exercise1.QuickSort(arr3, 0, len(arr3)-1)
// 	fmt.Println("QuickSort: ", arr3)

// }
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"app/exercise1" 
)

// Hàm hỗ trợ đo lường hiệu năng
func trackPerformance(name string, sortFunc func()) {
	var m runtime.MemStats
	
	runtime.GC() 
	runtime.ReadMemStats(&m)
	beforeAlloc := m.TotalAlloc

	start := time.Now() 
	sortFunc()       
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	afterAlloc := m.TotalAlloc
	
	fmt.Printf("=== %s ===\n", name)
	fmt.Printf("Thời gian: %v\n", duration)
	fmt.Printf("Bộ nhớ đã dùng: %v KB\n\n", (afterAlloc-beforeAlloc)/1024)
}

func main() {
	size := 10000
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(100000)
	}

	// 1. Đo BubbleSort (O(n^2))
	arr1 := make([]int, size)
	copy(arr1, data)
	trackPerformance("Bubble Sort", func() {
		exercise1.BubbleSort(arr1)
	})

	// 2. Đo MergeSort (O(n log n) + Tốn thêm bộ nhớ)
	arr2 := make([]int, size)
	copy(arr2, data)
	trackPerformance("Merge Sort", func() {
		exercise1.MergeSort(arr2)
	})

	// 3. Đo QuickSort (O(n log n) + Tối ưu bộ nhớ)
	arr3 := make([]int, size)
	copy(arr3, data)
	trackPerformance("Quick Sort", func() {
		exercise1.QuickSort(arr3, 0, len(arr3)-1)
	})
}
