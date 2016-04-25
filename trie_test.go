package yggdrasil

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
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

func TestInsertLongShort(t *testing.T) {
	trie := New()
	trie.Insert("tenth", 1)
	trie.Insert("tea", 2)
	trie.Insert("ten", 3)
	tNode, ok := trie.Children['t']
	assert.True(t, ok)
	eNode, ok := tNode.Children['e']
	assert.True(t, ok)
	nNode, ok := eNode.Children['n']
	assert.True(t, ok)
	otherTNode, ok := nNode.Children['t']
	assert.True(t, ok)
	_, ok = otherTNode.Children['h']
	assert.True(t, ok)

}

func TestFindPrefixes(t *testing.T) {
	trie := New()
	trie.Insert("tenth", 4)
	trie.Insert("tea", 1)
	trie.Insert("ted", 2)
	trie.Insert("ten", 3)
	trie.Insert("dog", 4)
	trie.Insert("to", 5)
	trie.Insert("teddie", 6)
	trie.Insert("teddy", 7)
	result := GetPrefixes(trie, "te")
	assert.Equal(t, 6, len(result))
	log.Println(result)
}
