package main


import (
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"

)

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
// Format link = /wiki/Title , e.g: /wiki/Albert_Einstein
func ScrapeLink(link string, node *TreeNode) {
  // Request the HTML page.
  if (link[0:6]=="/wiki/"){
	res, err := http.Get("https://en.wikipedia.org" + link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find all links
	doc.Find("a").Each(func(i int, s *goquery.Selection){
		link, exitst:= s.Attr("href")
		if (exitst && (len(link) >=6)){
			if (link[0:6]=="/wiki/"){
				fmt.Println("Link: ",link)				
			}
		}
	})
  }else{
	log.Fatal("Link Format is not correct!")
  }
  
}


func main() {
	// handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	resp := []byte(`{"status": "ok"}`)
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	// 	rw.Write(resp)
	// })

	// log.Println("Server is available at http://localhost:8000")
	// log.Fatal(http.ListenAndServe(":8000", handler))
	root := NewTreeNode("Orange","/wiki/Orange")
	root.PrintNode(2)
  	ScrapeLink(root.Link,root)
}

// test
// func main() {
//     root := NewTreeNode("Wikipedia","www.wikipedia.org")

//     child1 := NewTreeNode("Einstein","www.einstein.com")
//     child2 := NewTreeNode("Hawkins", "www.Hawkins.com")
//     child3 := NewTreeNode("Isaac", "www.Isaac.newt")

//     root.AddChild(child1)
//     root.AddChild(child2)
//     root.AddChild(child3)

//     subchild1 := NewTreeNode("Thomas","www.thomas.ed")
//     subchild2 := NewTreeNode("Tesla","www.tesla.nc")

//     child1.AddChild(subchild1)
//     child1.AddChild(subchild2)

//     fmt.Println("Number of nodes:", root.GetNumberOfNodes())
//     root.PrintNode(2);
// }