package yggdrasil

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestInsert(t *testing.T) {
	trie := New()
	trie.Insert("hello", 5)
	trie.Insert("hell", 42)
	if len(trie.Children) != 1 {
		t.Error("Should have 1 child, but has ", len(trie.Children))
	}
	h := *trie.Children['h']
	assert.Equal(t, EMPTY, h.Value, "Expected ", EMPTY, " actually ", h.Value)
	v, ok := trie.Find("hell")
	assert.True(t, ok, "Couldn't find key")
	assert.Equal(t, 42, v, "Incorrect value")
}

func TestFindPrefixes(t *testing.T) {
	trie := New()
	trie.Insert("tea", 1)
	trie.Insert("ted", 2)
	trie.Insert("ten", 3)
	trie.Insert("dog", 4)
	trie.Insert("to", 5)
	trie.Insert("teddie", 6)
	trie.Insert("teddy", 7)
	result := trie.GetPrefixes("t")
	assert.Equal(t, len(result), 6)
	log.Println(result)
}
