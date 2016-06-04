package main

import "fmt"

func main() {
	tree := newBTree(3)

	keys := make(keys, 0, 3)

	for i := 0; i < 3; i++ {
		keys = append(keys, []Key{4})
	}

	tree.Insert(keys...)

	fmt.Println(tree)
}
