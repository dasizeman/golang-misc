package main

import (
	"github.com/dasizeman/binarytree"
)

func main() {
	tree := binarytree.GenerateRandomIntTree(3, 1, 10)
	tree.PrintInOrder()
}
