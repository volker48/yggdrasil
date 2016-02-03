#Yggdrasil is an immense ash tree that is central and considered very holy

[Yggdrasil](https://en.wikipedia.org/wiki/Yggdrasil) is a Golang implementation of a [Trie](https://en.wikipedia.org/wiki/Trie) data structure.

###Usage
```go
package main

import (
	"github.com/volker48/yggdrasil"
	"log"
)

func main() {
	trie := yggdrasil.New()
	trie.Insert("hello", 42)
	trie.Insert("help", 3.14159)
	val, ok := trie.Find("hello")
	if !ok {
		log.Fatal("Couldn't find the key: ", "hello")
	}
	log.Printf("Key %s found its value is: %d", "hello", val)
	val, ok = trie.Find("help")
	if !ok {
		log.Fatal("Couldn't find key: ", "help")
	}
	log.Printf("Key %s found its value is: %f", "help", val)
}

```