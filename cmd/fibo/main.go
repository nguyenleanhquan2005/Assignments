package main

import (
	excercise3trietree "app/excercise3_trie_tree"
	"app/exercise2"
	"fmt"
)

func main() {
	//  FIBONACCI
	fmt.Println("= FIBONACCI =")
	fibonacci := exercise2.FibonacciImplementByFor(100)
	fmt.Println("F100 =", fibonacci)

	fibonacci3 := exercise2.FiboRecursive(100)
	fmt.Println("F100 =", fibonacci3)

	// TRIE TREE DEM
	fmt.Println("\n=TRIE TREE DEMO =")

	t := excercise3trietree.NewTrie()

	// ---- INSERT ----
	fmt.Println("\n Chèn các từ: cat, car, card, care, dog")
	t.Insert("cat")
	t.Insert("car")
	t.Insert("card")
	t.Insert("care")
	t.Insert("dog")

	// ---- SEARCH ----
	fmt.Println("\n Search (tìm từ chính xác):")
	words := []string{"cat", "car", "card", "care", "dog", "ca", "do", "dart"}
	for _, w := range words {
		fmt.Printf("  Search(%q) = %v\n", w, t.Search(w))
	}
	// ---- DELETE ----
	fmt.Println("\n Xóa từ \"cat\":")
	t.Delete("cat")
	fmt.Printf("  Search(\"cat\") sau khi xóa = %v\n", t.Search("cat"))
	fmt.Printf("  Search(\"car\") vẫn còn    = %v\n", t.Search("car"))

	// ---- PALINDROME ----
	fmt.Println("\n========== PALINDROME ==========")
	testCases := []string{
		"racecar",
		"hello",
		"A man a plan a canal Panama",
		"Was it a car or a cat I saw",
		"golang",
	}
	for _, s := range testCases {
		fmt.Printf("  IsPalindrome(%q) = %v\n", s, excercise3trietree.IsPalindrome(s))
	}

	fmt.Println("\n=====================================")

}
