package yggdrasil

import (
	"fmt"
)

const EMPTY = "**EMPTY**"

type Trie struct {
	Children map[rune]*Trie
	isWord   bool
	c        rune
	Value    interface{}
}

func (t *Trie) String() string {
	return fmt.Sprintf("Trie{c: %q, Value: %q, isWord: %q}", t.c, t.Value, t.isWord)
}

func New() (tree *Trie) {
	tree = &Trie{isWord: false, Value: EMPTY}
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

func sub(node *Trie, prefix, currentSuffix string) []string {
	matches := make([]string, 0)
	for c, child := range node.Children {
		if child.isWord {
			word := prefix+currentSuffix+string(c)
			matches = append(matches, word)
		}
		subMatches := sub(child, prefix, currentSuffix+string(c))
		matches = append(matches, subMatches...)
	}
	return  matches
}

func GetPrefixes(node *Trie, prefix string) []string {
	matches := make([]string, 0)
	for _, c := range prefix {
		child, found := node.Children[c]
		if !found {
			return matches
		}
		node = child
	}
	return sub(node, prefix, "")
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
	t.isWord = true
}
