package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var recursiveWalk func(t *tree.Tree)
	recursiveWalk = func(t *tree.Tree) {
		if t != nil {
			recursiveWalk(t.Left)
			ch <- t.Value
			recursiveWalk(t.Right)
		}
	}
	recursiveWalk(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chTree_1 := make(chan int)
	chTree_2 := make(chan int)

	go Walk(t1, chTree_1)
	go Walk(t2, chTree_2)

	for {
		node_1, ok_1 := <-chTree_1
		node_2, ok_2 := <-chTree_2
		if node_1 != node_2 || ok_1 != ok_2 {
			return false
		}
		if !ok_1 || !ok_2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for node := range ch {
		fmt.Println(node)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
