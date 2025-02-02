package main

import (
	"fmt"
)

// def
type TreeNode struct {
	Title     string
	Link      string
	Parent    *TreeNode
	Children  []*TreeNode
	id        int
	imagePath string
}

// ctor
func NewTreeNode(title string, link string) *TreeNode {
	return &TreeNode{
		Title:     title,
		Link:      link,
		Parent:    nil,
		Children:  []*TreeNode{},
		id:        0,
		imagePath: "",
	}
}

// add children node
func (node *TreeNode) AddChild(child *TreeNode) {
	child.Parent = node
	child.id = node.id + 1
	node.Children = append(node.Children, child)
}

// get children num
func (node *TreeNode) GetNumberOfNodes() int {
	count := 1
	for _, child := range node.Children {
		count += child.GetNumberOfNodes()
	}
	return count
}
func (node *TreeNode) GetNumberOfChildren() int {
	return len(node.Children)
}

// print node (for debug)
func (node *TreeNode) PrintNode(indentation int) {
	fmt.Print(node.Title)
	fmt.Print(" ")
	fmt.Print(node.Link)
	fmt.Println()

	for _, child := range node.Children {
		for i := 0; i < indentation; i++ {
			fmt.Print("  ")
		}
		child.PrintNode(indentation + 1)
	}
}
