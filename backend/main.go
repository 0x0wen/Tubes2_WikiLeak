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
func isSameNode(node1 *TreeNode, node2 *TreeNode, target string) bool {
	temp1 := node1
	temp2 := node2
	if temp2.Link == target && temp1.Link == target {
		for temp1 != nil && temp2 != nil {
			if temp1.Link != temp2.Link {
				return false
			}
			temp1 = temp1.Parent
			temp2 = temp2.Parent
		}
	} else {
		return node1.Link == node2.Link
	}
	return true
}
func isAlreadyExist(node *TreeNode, nodeList []*TreeNode, target string) bool {
	for i := 0; i < len(nodeList); i++ {
		if isSameNode(node, nodeList[i], target) {
			return true
		}
	}
	return false
}

func getTitle(link string) string {
	// Instantiate a new collector
	c := colly.NewCollector()
	title := ""
	// Find and visit link
	c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
		// Extract text or any other attribute you want

		title = e.Text
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Scraping finished for", r.Request.URL.String())

	})
	// Visit the URL you want to scrape
	c.Visit("https://en.wikipedia.org" + link)

	c.Wait()
	return title
}

// add children node
func (node *TreeNode) AddChild(child *TreeNode) {
	child.Parent = node
	child.id = node.id + 1
	node.Children = append(node.Children, child)
}
func removeChild(node *TreeNode, index int) {
	// Check if index is valid
	if index < 0 || index >= len(node.Children) {
		return // Index out of range
	}

	// Remove the child at the specified index
	node.Children = append(node.Children[:index], node.Children[index+1:]...)
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

func ScrapeLink(node *TreeNode, target string, listLink []*TreeNode) {
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
		c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
			// Extract text or any other attribute you want
			node.Title = e.Text
		})
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Extract the href attribute of the <a> element
			link := e.Attr("href")
			teks := e.Text
			// results <- link // Send the link to the results channel
			if len(link) >= 6 {
				if link[0:6] == "/wiki/" {
					node.AddChild(NewTreeNode(teks, link))
					if isAlreadyExist(node.Children[len(node.Children)-1], listLink, target) {
						removeChild(node, len(node.Children)-1)
					} else {
						listLink = append(listLink, node.Children[len(node.Children)-1])
					}

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

func BFSRace(node *TreeNode, target string, listLink []*TreeNode, mutex *sync.Mutex) *TreeNode {
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
			ScrapeLink(queue[0], target, listLink)
			queue = append(queue, queue[0].Children...)
			mutex.Unlock()
		} else if len(queue)-i < 17 {

			// if j*17+18 < len(queue) {

			for k := j*17 + 1; k < j*17+18; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, listLink)

				}(k)

			}
			wg.Wait()

			for k := j*17 + 1; k < j*17+18; k++ {
				mutex.Lock()

				queue = append(queue, queue[k].Children...)
				mutex.Unlock()

			}
			j += 1

		}

	}

	// Target node not found
	return nil

}
func BFSRaceBonus(node *TreeNode, target string, listLink []*TreeNode, mutex *sync.Mutex) []*TreeNode {
	// Request the HTML page.
	if node == nil {
		return nil
	}

	// Create a queue for BFS
	queue := []*TreeNode{node}
	result := []*TreeNode{}

	// Perform BFS
	j := 0
	depth := -1
	var wg sync.WaitGroup

	for i := 0; depth == -1 || queue[len(queue)-1].id <= depth; i++ {
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
			if depth == -1 {
				depth = current.id
			}
			result = append(result, current)

		}
		// 	scraping
		if i == 0 {
			mutex.Lock()
			ScrapeLink(queue[0], target, listLink)
			queue = append(queue, queue[0].Children...)
			mutex.Unlock()
		} else if len(queue)-i < 30 {

			// if j*30+31 < len(queue) {

			for k := j*30 + 1; k < j*30+31; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, listLink)

				}(k)

			}
			wg.Wait()

			for k := 30*j + 1; k < j*30+31; k++ {
				queue = append(queue, queue[k].Children...)

			}
			j += 1

			// }
			// } else {
			// 	mutex.Lock()
			// 	for k := j*30 + 1; k < len(queue); k++ {
			// 		wg.Add(1)
			// 		go func(k int) {
			// 			defer wg.Done()
			// 			ScrapeLink(queue[k], target, listLink)
			// 			var newChildren []*TreeNode

			// 			// Iterate over each element in the queue
			// 			for _, element := range queue[k].Children {
			// 				// Append the children of the current element to the newChildren slice
			// 				if !isAlreadyExist(element, queue, target) {
			// 					newChildren = append(newChildren, element)
			// 				}
			// 			}
			// 			queue = append(queue, newChildren...)
			// 		}(k)
			// 	}
			// 	mutex.Unlock()

			// 	j += 1

			// }

		}

	}

	return result

}

func IDSrecursion(node []*TreeNode, depth int, target string, found *bool, result *TreeNode, listLink []*TreeNode,
	wg *sync.WaitGroup, sem chan struct{}, mutex *sync.Mutex) {

	if depth == 0 {
		for i := 0; !*found && i < len(node); i++ {
			if node[i].Parent != nil {
				fmt.Println("IDS 1 : ", node[i].Parent.Link, " ", node[i].Link, " ", node[i].id)
			}
			if node[i].Link == target {
				*result = *node[i]
				*found = true
			}
		}

	} else {

		for i := 0; !*found && i < len(node); i++ {
			if node[i].Parent != nil {
				fmt.Println("IDS ", depth+1, " :  ", node[i].Parent.Link, " ", node[i].Link, " ", node[i].id)
			}
			if node[i].Link == target {
				*result = *node[i]
				*found = true
			} else {
				if len(node[i].Children) == 0 {
					wg.Add(1)
					sem <- struct{}{}

					go func(i int) {
						defer func() {
							<-sem // Release semaphore after scraping is done
							wg.Done()
						}()

						ScrapeLink(node[i], target, listLink)

					}(i)

				}
				wg.Wait()

				// Acquire semaphore before launching a new scraping task

				// Launch a new scraping task in a goroutine
				// mutex.Lock()
				// fmt.Println("length: ", len(node[i].Children))
				// for j := 0; j < len(node[i].Children) && !*found; j++ {
				// 	if node[i].Children[j].Link == target {
				// 		*result = *node[i].Children[j]
				// 		*found = true
				// 	} else {
				// 		fmt.Println(node[i].Children[j].Link, " bukan linknya")
				// 	}
				// }
				// mutex.Unlock()
				// fmt.Println("test")

				if !*found {
					IDSrecursion(node[i].Children, depth-1, target, found, result, listLink, wg, sem, mutex)
				}

			}

		}

	}
}

func IDSRace(node *TreeNode, target string, listLink []*TreeNode, mutex *sync.Mutex) *TreeNode {
	if node == nil {
		return nil
	}

	result := NewTreeNode("", "")
	queue := []*TreeNode{node}
	found := false
	current := queue[0]
	sem := make(chan struct{}, 20)
	var wg sync.WaitGroup

	ScrapeLink(node, target, listLink)

	for depth := 1; !found; depth++ {
		if queue[0].Link == target {
			return queue[0]
		}

		IDSrecursion(current.Children, depth-1, target, &found, result, listLink, &wg, sem, mutex)

	}
	return result
}

func main() {
	startTime := time.Now()
	// handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	resp := []byte(`{"status": "ok"}`)
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	// 	rw.Write(resp)
	// })

	// title := getTitle("/wiki/Albert_Einstein")
	// fmt.Println("Ttitle: ", title)
	// log.Println("Server is available at http://localhost:8000")
	// log.Fatal(http.ListenAndServe(":8000", handler))
	var mutex sync.Mutex
	root := NewTreeNode("", "/wiki/ITB")
	// ScrapeLink(root, "/wiki/Sukarno", &mutex)
	// for i := 0; i < root.GetNumberOfChildren(); i++ {
	// 	ScrapeLink(root.Children[i], "/wiki/Sukarno", &mutex)
	// }
	// root.PrintNode(3)
	listLink := []*TreeNode{}
	result := IDSRace(root, "/wiki/Computer", listLink, &mutex)
	// result := IDSRace(root, "/wiki/Computer", &mutex)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Seconds(), " s")

	for result != nil {
		fmt.Println("Title: ", result.Title)
		fmt.Println("Link: ", result.Link)
		result = result.Parent
	}
	// Target node not found
	// for i := 0; i < len(result); i++ {
	// 	fmt.Println("Result ", i+1, " : ")
	// 	for result[i] != nil {
	// 		fmt.Println("Title: ", result[i].Title)
	// 		fmt.Println("Link: ", result[i].Link)
	// 		result[i] = result[i].Parent
	// 	}
	// }

}
