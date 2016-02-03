package yggdrasil


type Trie struct {
	Children map[rune]Trie
	leaf bool
}

func New() (tree *Trie) {
	tree = &Trie{leaf: false}
	tree.Children = make(map[byte]Trie)
	return tree
}

func (t *Trie) Find(key string) (string, bool) {
	for _, c := range key {
		_, ok := t.Children[c]
		if !ok {
			return "", false
		} else {
			t = t.Children[c]
		}
	}
	return true
}


func (t *Trie) Insert(value string) {
	for _, c := range value {
		_, ok := t.Children[c]
		if !ok {
			children := make(map[byte]Trie)
			t.Children[c] = &Trie{Children: &Trie{Children: children}}
		}
	}
	t.leaf = true
}
