package tst

import "fmt"

// import "fmt"

type node struct {
	c     byte
	left  *node
	right *node
	mid   *node
	val   interface{}
}

// Trie represents a Ternary Search Trie. An empty value can be used to insert
// key/value pairs.
type Trie struct {
	root *node
	n    int
}

// Size returns number of elements in the trie.
func (t *Trie) Size() int {
	return t.n
}

// Put inserts a key/value pair into the trie. If the key is already in the
// trie, it's value will be updated. Note: empty strings will be ignored.
func (t *Trie) Put(key string, value interface{}) {
	t.put(key, value)
}

// Get returns the value associated with key.
func (t *Trie) Get(key string) interface{} {
	return t.get(key)
}

// // GetShortestPrefix searches for the shortest key which is a prefix of the
// // given key and returns the value. For example, if trie contains key "com"
// // and "com.g", search for "com.google" will return value associated with
// // "com".
// func (t *Trie) GetShortestPrefix(key string) interface{} {
// 	return t.getShortestPrefix(key, false)
// }

// // GetShortestPrefixRK is the same as GetShortestPrefix, but the key is
// // reversely iterated when searching for the value.
// func (t *Trie) GetShortestPrefixRK(key string) interface{} {
// 	return t.getShortestPrefix(key, true)
// }

func iterStartEndStep(key string) (start, end, step int) {
	return 0, len(key) - 1, 1
}

func (t *Trie) put(key string, value interface{}) {
	if len(key) < 1 {
		return
	}

	start, end, step := iterStartEndStep(key)

	t.n++
	pnd := &t.root
	i := start
	c := key[start]
	for {
		if *pnd == nil {
			// fmt.Printf("new node: %c\n", c)
			*pnd = &node{c: c}
		}
		switch {
		case c < (*pnd).c:
			pnd = &(*pnd).left
		case c > (*pnd).c:
			pnd = &(*pnd).right
		case i != end:
			pnd = &(*pnd).mid
			i += step
			c = key[i]
		default:
			(*pnd).val = value
			// fmt.Println("add value:", value)
			return
		}
	}
}

func (t *Trie) get(key string) interface{} {
	if len(key) < 1 {
		return nil
	}

	start, end, step := iterStartEndStep(key)

	nd := t.root
	i := start
	c := key[start]
	// go down the tree
	for nd != nil {
		switch {
		case c < nd.c:
			nd = nd.left
		case c > nd.c:
			nd = nd.right
		case i != end:
			nd = nd.mid
			i += step
			c = key[i]
		default:
			return nd.val
		}
	}
	return nil
}

// func (t *Trie) getShortestPrefix(key string, reverseKey bool) interface{} {
// 	// Most code copied from Get.
// 	if len(key) < 1 {
// 		return nil
// 	}

// 	start, end, step := iterStartEndStep(key, reverseKey)

// 	nd := t.root
// 	i := start
// 	c := key[start]
// 	for nd != nil {
// 		switch {
// 		case c < nd.c:
// 			nd = nd.left
// 		case c > nd.c:
// 			nd = nd.right
// 		// This is the only added code to get.
// 		case nd.val != nil:
// 			return nd.val
// 		case i != end:
// 			nd = nd.mid
// 			i += step
// 			c = key[i]
// 		default:
// 			return nd.val
// 		}
// 	}
// 	return nil
// }

func (t *Trie) inorderTraversal(n node, word string) {
	if n.val != nil {
		fmt.Println(word + string(n.c) + ": ")
	}

	if n.left != nil {
		t.inorderTraversal(*n.left, word)
	}

	if n.mid != nil {
		t.inorderTraversal(*n.mid, word+string(n.c))
	}

	if n.right != nil {
		t.inorderTraversal(*n.right, word)
	}
}

func (t *Trie) Traverse() {
	t.inorderTraversal(*t.root, "")
}
