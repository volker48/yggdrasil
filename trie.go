package yggdrasil

const EMPTY = "**EMPTY**"

type Trie struct {
	Children map[rune]*Trie
	leaf bool
	Value interface{}
}

func New() (tree *Trie) {
	tree = &Trie{leaf: false, Value: EMPTY}
	tree.Children = make(map[rune]*Trie)
	return tree
}

func (t *Trie) Find(key string) (interface{}, bool) {
	for _, c := range key {
		_, ok := t.Children[c]
		if !ok {
			return nil, false
		} else {
			t = t.Children[c]
		}
	}
	return t.Value, true
}


func (t *Trie) Insert(key string, value interface{}) {
	for _, c := range key {
		_, ok := t.Children[c]
		if !ok {
			t.Children[c] = New()
		}
		t, _ = t.Children[c]
	}
	t.Value = value
}
