package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var mapCache sync.Map

type Cache struct {
	mu      sync.Mutex
	visited map[string]bool
}

func NewCache() *Cache {
	return &Cache{
		visited: make(map[string]bool),
	}
}

func (c *Cache) IsVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.visited[url]
}

func (c *Cache) MarkVisited(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.visited[url] = true
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

func isLinkValid(node *TreeNode) bool {
	if strings.Contains(node.Link, "#") || strings.Contains(node.Link, "Main_Page") || strings.Contains(node.Link, ":") {
		return false
	} else {
		return true
	}
}

// Format link = /wiki/Title , e.g: /wiki/Albert_Einstein

// func ScrapeLink1(node *TreeNode, target string, listLink []*TreeNode) {
// 	// if node.Parent != nil {
// 	// 	fmt.Println("Scrape : ", node.Parent.Link, "  ", node.Link, "  ", node.id)
// 	// } else {
// 	// 	fmt.Println("Scrape : ", node.Link, "  ", node.id)

// 	// }
// 	if node.Link[0:6] == "/wiki/" {
// 		url := "https://en.wikipedia.org" + node.Link
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Printf("Error fetching %s: %s\n", url, err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			fmt.Printf("Failed to fetch %s: %s\n", url, resp.Status)
// 			return
// 		}

// 		// Parse the HTML response using streaming
// 		tokenizer := html.NewTokenizer(resp.Body)
// 		for {
// 			tokenType := tokenizer.Next()
// 			switch tokenType {
// 			case html.ErrorToken:
// 				// End of the document, or an error occurred
// 				return
// 			case html.StartTagToken, html.SelfClosingTagToken:
// 				token := tokenizer.Token()
// 				if token.Data == "a" {
// 					var link, text string
// 					isLink := false
// 					for _, attr := range token.Attr {
// 						if attr.Key == "href" && len(attr.Val) > 6 {
// 							if attr.Val[0:6] == "/wiki/" && attr.Val != node.Link {
// 								link = attr.Val // Store the href attribute value
// 								isLink = true
// 								break
// 							}

// 						}
// 					}
// 					// Get the text content inside the <a> tag
// 					if isLink {
// 						tokenType = tokenizer.Next()
// 						if tokenType == html.TextToken {
// 							text = strings.TrimSpace(tokenizer.Token().Data)
// 						}
// 						node.AddChild(NewTreeNode(text, link))
// 						if isAlreadyExist(node.Children[len(node.Children)-1], listLink, target) || !isLinkValid(node.Children[len(node.Children)-1]) {
// 							removeChild(node, len(node.Children)-1)
// 						} else {
// 							listLink = append(listLink, node.Children[len(node.Children)-1])

// 						}
// 					}

// 				} else if token.Data == "span" {
// 					for _, a := range token.Attr {
// 						if a.Key == "class" && strings.Contains(a.Val, "mw-page-title-main") {
// 							tokenizer.Next()
// 							pageTitle := tokenizer.Token().Data
// 							node.Title = pageTitle
// 							break
// 						}
// 					}
// 				}
// 			}
// 		}
// 	} else {
// 		fmt.Println("ERROR: Input link is not valid !")
// 	}

// }
var scrap = 0

func ScrapeLink(node *TreeNode, target string, cache *Cache) {
	scrap += 1
	if node.Parent != nil {
		fmt.Println("Scrape ", scrap, ": ", node.Parent.Link, "  ", node.Link, "  ", node.id)
	} else {
		fmt.Println("Scrape ", scrap, ": ", node.Link, "  ", node.id)

	}
	if node.Link[0:6] == "/wiki/" {

		// res, err := http.Get("https://en.wikipedia.org" + node.Link)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer res.Body.Close()
		// if res.StatusCode != 200 {
		// 	// log.Fatal("status code error: %d %s", res.StatusCode, res.Status)
		// 	return
		// }
		if value, ok := mapCache.Load(node.Link); ok {
			for i := 0; i < len(value.([]*TreeNode)); i++ {
				node.AddChild(NewTreeNode(value.([]*TreeNode)[i].Title, value.([]*TreeNode)[i].Link))
			}
			return
		}
		c := colly.NewCollector(
			colly.AllowedDomains("en.wikipedia.org"),
			// colly.Async(true),
		)

		q, _ := queue.New(
			15, // Number of consumer threads
			&queue.InMemoryQueueStorage{MaxSize: 200}, // Use in-memory queue storage
		)

		q.AddURL("https://en.wikipedia.org" + node.Link)

		// Define a callback function to be executed when a link is found
		c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
			// Extract text or any other attribute you want
			node.Title = strings.TrimSpace(e.DOM.Text())
		})
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Extract the href attribute of the <a> element
			link := e.Attr("href")
			teks := e.Text
			// results <- link // Send the link to the results channel
			if !cache.IsVisited(link) {
				if link != target {
					cache.MarkVisited(link)
				}
				if len(link) >= 6 {
					if link[0:6] == "/wiki/" {
						if strings.Contains(link, "#") || strings.Contains(link, "Main_Page") || strings.Contains(link, ":") || link == node.Link {
							// do nothing
						} else {
							node.AddChild(NewTreeNode(teks, link))
						}

					}
				}
			}

		})

		// Define a callback function to be executed when the scraping is complete
		c.OnScraped(func(r *colly.Response) {
			if _, ok := mapCache.Load(node.Link); ok {
				mapCache.Store(node.Link, node.Children)
			}
		})
		q.Run(c)

	}
}
