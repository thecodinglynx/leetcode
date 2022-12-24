package leetcode

// this solution uses a Trie data structure to determine the
// longest common prefix for the provided slices of strings.
// The Trie data structure is very performant O(K) for both
// insertion and search.
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	root := NewNode()

	// we insert the first word - it doesn't matter if
	// this word is the shortest, the longest or somehwere
	// in the middle, the algorithm works the same way either way.
	root.Insert([]rune(strs[0]))
	res := strs[0]

	for i := 1; i < len(strs); i++ {
		depth := root.getDepth([]rune(strs[i]))
		if depth < len(res) {
			val := strs[i]
			res = val[:depth]
		}
	}

	return res
}

type TrieNode struct {
	children *map[rune]TrieNode
}

func NewNode() *TrieNode {
	children := make(map[rune]TrieNode)
	return &TrieNode{children: &children}
}

func (t *TrieNode) Insert(val []rune) {
	cur := t
	for _, v := range val {
		new := NewNode()
		(*cur.children)[v] = *new
		cur = new
	}
}

// returns how deep the provided slice of runes goes into the Trie
func (t *TrieNode) getDepth(val []rune) int {
	curNode := t
	depth := 0
	for _, v := range val {
		if next, ok := (*curNode.children)[v]; ok {
			curNode = &next
			depth++
		} else {
			p := make(map[rune]TrieNode)
			curNode.children = &p
			return depth
		}
	}
	return depth
}
