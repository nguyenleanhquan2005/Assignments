package main

import (
	"fmt"
	// Import package exercise1. 
	// Đường dẫn này phụ thuộc vào tên module bạn đặt ở bước 3.
	// Giả sử tên module là "myproject"
	"assignments/exercise1" 
)

func main() {
	fmt.Println("--- Bắt đầu chương trình ---")

	// 1. Tạo cây
	root := exercise1.CreateTree()

	// 2. In thử InOrder
	fmt.Println("InOrder traversal:")
	exercise1.InOrder(root)

	fmt.Println("InOrder traversal:")
	exercise1.PreOrder(root)


	// 3. Tìm số lớn nhất
	maxVal := exercise1.GetBiggestNumber(root)
	fmt.Printf("Biggest number is: %d\n", maxVal)

	// 4. Lấy mảng đã sắp xếp
	sortedList := exercise1.GetSortedNumbers(root)
	fmt.Println("sorted list:", sortedList)


	input := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Before", input)
	exercise1.BubbleSort(input)
	fmt.Println("After", input)
}