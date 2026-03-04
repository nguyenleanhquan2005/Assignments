package excercise3trietree


type TrieNode struct {
	children map[rune]*TrieNode
	isEnd bool

}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children : make(map[rune]*TrieNode),
		isEnd : false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie{
	return &Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) Insert(word string){
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = newTrieNode()
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

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



func (t *Trie) Delete(word string) bool {
	return deleteHelper(t.root, word, 0)
}

func deleteHelper(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	child, exists := node.children[ch]
	if !exists {
		return false
	}

	shouldDeleteChild := deleteHelper(child, word, depth+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEnd
	}

	return false
}