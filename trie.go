package yggdrasil

const EMPTY = "**EMPTY**"

type Trie struct {
	Children map[rune]*Trie
	leaf bool
	c rune
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

func (t *Trie) GetPrefixes(prefix string) (matches []string) {
	for _, c := range prefix {
		child, ok := t.Children[c]
		if !ok {
			return
		}

		t = child
	}
	queue := []*Trie{t}
	suffix := ""
	for len(queue) > 0 {
		t, queue = queue[len(queue)-1], queue[:len(queue)-1]
		for _, child := range t.Children {
			queue = append(queue, child)
		}
		suffix += string(t.c)
		if t.leaf {
			matches = append(matches, suffix)
			suffix = suffix[:len(suffix) - 1]
		}
	}
	return
}


func (t *Trie) Insert(key string, value interface{}) {
	for _, c := range key {
		_, ok := t.Children[c]
		if !ok {
			t.Children[c] = New()
			t.Children[c].c = c
		}
		t, _ = t.Children[c]
	}
	t.Value = value
	t.leaf = true
}

