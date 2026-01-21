package exercise1

import (
	"fmt"
)

func SortedBinaryTree() {


}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InOrder(node *TreeNode) {
    // 1. Kiểm tra điểm dừng (Base case)
    if node == nil {
        return 
    }
    // 2. Đệ quy sang nhánh Trái
    InOrder(node.Left)

    // 3. Xử lý nút hiện tại (Gốc) - In giá trị ra
	fmt.Println(node.Val)
    // 4. Đệ quy sang nhánh Phải
    
	InOrder(node.Right)
}

func PreOrder(node *TreeNode) {
    // 1. Kiểm tra điểm dừng (Base case)
    if node == nil {
        return 
    }
    // 2. Xử lý nút hiện tại (Gốc) - In giá trị ra

    fmt.Println(node.Val)

    // 3. Đệ quy sang nhánh Trái

    PreOrder(node.Left)

    // 4. Đệ quy sang nhánh Phải
    
    PreOrder(node.Right)
}

func PostOrder(node *TreeNode) {
    // 1. Kiểm tra điểm dừng (Base case)
    if node == nil {
        return 
    }
    // 2. Đệ quy sang nhánh Trái
    PostOrder(node.Left)
    
	PostOrder(node.Right)

	fmt.Println(node.Val)

}

func CreateTree()*TreeNode{
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 2}
    root.Left.Right = &TreeNode{Val: 7}

    root.Right = &TreeNode{Val: 15}
    root.Right.Left = &TreeNode{Val: 11}
    root.Right.Right = &TreeNode{Val: 19}
    root.Right.Right.Right = &TreeNode{Val: 26}

    return root 
}

func GetBiggestNumber(root *TreeNode) int{
    if root == nil{
        return 0
    }
    current := root
    for current.Right != nil{
        current = current.Right
    }
    return current.Val

}

func GetSortedNumbers(root *TreeNode) [] int{
    var result []int

    var traverse func(node * TreeNode)

    traverse = func(node *TreeNode){
        if node == nil{
            return 
        }
        traverse(node.Left)

        result = append(result, node.Val)

        traverse(node.Right)

    }
    
    traverse(root)

    return result

} 


