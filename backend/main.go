package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// def
type TreeNode struct {
	Title    string
	Link     string
	Parent   *TreeNode
	Children []*TreeNode
	id       int
}

// ctor
func NewTreeNode(title string, link string) *TreeNode {
	return &TreeNode{
		Title:    title,
		Link:     link,
		Parent:   nil,
		Children: []*TreeNode{},
		id:       0,
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

// Format link = /wiki/Title , e.g: /wiki/Albert_Einstein

func ScrapeLink(node *TreeNode, target string) {
	// Request the HTML page.
	if node.Parent != nil {
		fmt.Println("Scrapping: ", node.Parent.Link, " ", node.Link)

	} else {
		fmt.Println("Scrapping: ", node.Link)
	}

	if node.Link[0:6] == "/wiki/" {
		res, err := http.Get("https://en.wikipedia.org" + node.Link)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			// log.Fatal("status code error: %d %s", res.StatusCode, res.Status)
			return
		}
		// Create a new Collector with concurrency settings
		c := colly.NewCollector(
			colly.Async(true),
		)

		// Define a callback function to be executed when a link is found
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Extract the href attribute of the <a> element
			link := e.Attr("href")
			// results <- link // Send the link to the results channel
			if len(link) >= 6 {
				if link[0:6] == "/wiki/" {
					node.AddChild(NewTreeNode("", link))

				}
			}

		})

		// Define a callback function to be executed when the scraping is complete
		c.OnScraped(func(r *colly.Response) {
			// fmt.Println("Scraping finished for", r.Request.URL.String())
		})

		// Start scraping by visiting the URL
		c.Visit("https://en.wikipedia.org" + node.Link)

		// Wait for scraping to finish
		c.Wait()

	}

}

func BFSRace(node *TreeNode, target string, mutex *sync.Mutex) *TreeNode {
	// Request the HTML page.
	if node == nil {
		return nil
	}

	// Create a queue for BFS
	queue := []*TreeNode{node}

	// Perform BFS
	found := false
	j := 0
	var wg sync.WaitGroup

	for i := 0; !found; i++ {
		// Dequeue a node from the front of the queue
		mutex.Lock()
		current := queue[i]
		mutex.Unlock()

		if current.Parent != nil {
			fmt.Println("BFS: ", current.Parent.Link, " ", current.Link, " ", current.id)
		}

		// Check if the current node's link matches the target link
		if current.Link == target {
			// Modify the title of the target node
			return current // Found the target node, return the modified node
		}
		// 	scraping
		if i == 0 {
			mutex.Lock()
			ScrapeLink(queue[0], target)
			queue = append(queue, queue[0].Children...)
			mutex.Unlock()
		} else if len(queue)-i < 30 {

			if j*4+5 < len(queue) {
				mutex.Lock()

				for k := j*20 + 1; k < j*20+21; k++ {
					wg.Add(1)
					go func(k int) {
						defer wg.Done()
						ScrapeLink(queue[k], target)
						queue = append(queue, queue[k].Children...)
					}(k)

				}
				j += 1
				mutex.Unlock()

			} else {
				mutex.Lock()
				wg.Add(1)
				for k := j*20 + 1; k < len(queue); k++ {
					go func(k int) {
						defer wg.Done()
						ScrapeLink(queue[k], target)
						queue = append(queue, queue[k].Children...)
					}(k)
				}
				mutex.Unlock()

				j += 1

			}

		}
		wg.Wait()

	}

	// Target node not found
	return nil

}

func main() {
	startTime := time.Now()
	// handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	resp := []byte(`{"status": "ok"}`)
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	// 	rw.Write(resp)
	// })

	// log.Println("Server is available at http://localhost:8000")
	// log.Fatal(http.ListenAndServe(":8000", handler))
	var mutex sync.Mutex
	root := NewTreeNode("", "/wiki/ITB")
	// ScrapeLink(root, "/wiki/Sukarno", &mutex)
	// for i := 0; i < root.GetNumberOfChildren(); i++ {
	// 	ScrapeLink(root.Children[i], "/wiki/Sukarno", &mutex)
	// }
	// root.PrintNode(3)
	result := BFSRace(root, "/wiki/Computer", &mutex)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Seconds(), " s")
	for result != nil {
		fmt.Println("Result")
		fmt.Println(result.Link)
		result = result.Parent
	}
}
