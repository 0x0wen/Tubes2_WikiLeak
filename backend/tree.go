package main

import "fmt"
// def
type TreeNode struct {
    Title    string
    Link     string
    Parent   *TreeNode
    Children []*TreeNode
}
// ctor
func NewTreeNode(title string, link string) *TreeNode {
    return &TreeNode{
        Title   : title,
        Link    : link,
        Parent:   nil,
        Children: []*TreeNode{},
    }
}
// add children node
func (node *TreeNode) AddChild(child *TreeNode) {
    child.Parent = node
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
// test
func main() {
    root := NewTreeNode("Wikipedia","www.wikipedia.org")

    child1 := NewTreeNode("Einstein","www.einstein.com")
    child2 := NewTreeNode("Hawkins", "www.Hawkins.com")
    child3 := NewTreeNode("Isaac", "www.Isaac.newt")

    root.AddChild(child1)
    root.AddChild(child2)
    root.AddChild(child3)

    subchild1 := NewTreeNode("Thomas","www.thomas.ed")
    subchild2 := NewTreeNode("Tesla","www.tesla.nc")

    child1.AddChild(subchild1)
    child1.AddChild(subchild2)

    fmt.Println("Number of nodes:", root.GetNumberOfNodes())
    root.PrintNode(2);
}
