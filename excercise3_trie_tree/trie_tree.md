package excercise3trietree

// ========================================
// BƯỚC 1: Định nghĩa TrieNode
// Mỗi node đại diện cho một ký tự trong cây Trie.
// - children: map từ ký tự → node con tiếp theo
// - isEnd: true nếu đây là ký tự cuối cùng của một từ
// ========================================
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// newTrieNode tạo một TrieNode mới với map children rỗng.
func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// ========================================
 BƯỚC 2: Định nghĩa Trie
 Trie chỉ cần lưu node gốc (root).
 Root không đại diện cho ký tự nào cả.
// ========================================
type Trie struct {
	root *TrieNode
}

// NewTrie tạo một Trie mới với node gốc rỗng.
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// ========================================
// BƯỚC 3: Insert - Chèn một từ vào Trie
//
// Ví dụ: Insert("cat")
//
//	root → 'c' → 'a' → 't' (isEnd=true)
//
// Cách làm:
//  1. Bắt đầu từ root
//  2. Với mỗi ký tự trong từ:
//     - Nếu chưa có node con tương ứng → tạo mới
//     - Di chuyển xuống node con đó
//  3. Đánh dấu node cuối là isEnd = true
//
// ========================================
func (t *Trie) Insert(word string) {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = newTrieNode()
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

// ========================================
// BƯỚC 4: Search - Tìm xem từ có tồn tại không
//
// Ví dụ: Search("cat") → true, Search("ca") → false (nếu "ca" chưa được insert)
//
// Cách làm:
//  1. Bắt đầu từ root
//  2. Với mỗi ký tự:
//     - Nếu không có node con → return false (từ không tồn tại)
//     - Di chuyển xuống node con
//  3. Kiểm tra isEnd của node cuối
//     - isEnd=true  → từ tồn tại hoàn chỉnh
//     - isEnd=false → chỉ là tiền tố, không phải từ
//
// ========================================
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

// ========================================
// BƯỚC 5: StartsWith - Kiểm tra tiền tố (prefix)
//
// Ví dụ: StartsWith("ca") → true (nếu đã insert "cat" hoặc "car")
//
// Cách làm: Giống Search, nhưng KHÔNG cần kiểm tra isEnd.
// Chỉ cần đi qua được hết các ký tự là đủ.
// ========================================
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return true
}

// ========================================
// BƯỚC 6: Delete - Xóa một từ khỏi Trie
//
// Ví dụ: Xóa "cat" nhưng giữ lại "car"
//
// Cách làm (dùng đệ quy):
//  1. Duyệt đệ quy đến cuối từ
//  2. Đánh dấu isEnd = false ở node cuối
//  3. Trên đường quay về:
//     - Nếu node không còn con nào và không phải kết thúc từ khác
//     → xóa node đó khỏi cha
//     Kết quả: Chỉ xóa các node "thừa", không ảnh hưởng từ khác
//
// ========================================
func (t *Trie) Delete(word string) bool {
	return deleteHelper(t.root, word, 0)
}

// deleteHelper là hàm đệ quy hỗ trợ Delete.
// - node: node hiện tại đang xét
// - word: từ cần xóa
// - depth: vị trí ký tự hiện tại trong word
// Trả về true nếu node hiện tại có thể bị xóa khỏi cha.
func deleteHelper(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	// Đã đi qua hết ký tự
	if depth == len(word) {
		if !node.isEnd {
			return false // Từ không tồn tại trong Trie
		}
		node.isEnd = false
		// Có thể xóa node này nếu nó không có con nào
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	child, exists := node.children[ch]
	if !exists {
		return false // Ký tự không tồn tại → từ không có trong Trie
	}

	// Đệ quy xuống node con
	shouldDeleteChild := deleteHelper(child, word, depth+1)

	if shouldDeleteChild {
		delete(node.children, ch) // Xóa node con khỏi map
		// Node hiện tại có thể bị xóa nếu không còn con và không là cuối từ khác
		return len(node.children) == 0 && !node.isEnd
	}

	return false
}
