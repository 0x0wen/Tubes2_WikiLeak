package main

import (
	"log"
	"net/http"
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
func solutionAlreadyExist(node *TreeNode, solutionList []*TreeNode) bool {
	for i := 0; i < len(solutionList); i++ {
		if node.Parent == nil {
			if node.Link == solutionList[i].Link {
				return true
			}
		} else {
			if node.Link == solutionList[i].Link && node.Parent.Link == solutionList[i].Parent.Link {
				return true
			}
		}
	}
	return false
}

func getImage(link string) string {
	// Instantiate a new collector
	c := colly.NewCollector()
	q, _ := queue.New(
		15, // Number threads
		&queue.InMemoryQueueStorage{MaxSize: 200},
	)

	q.AddURL("https://en.wikipedia.org" + link)

	// Find and visit link
	img := ""
	var gained bool
	c.OnHTML("a.mw-file-description img", func(e *colly.HTMLElement) {
		// extract image
		if !gained {
			img = e.Attr("src")

			gained = true
		}

	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Scraping finished for", r.Request.URL.String())

	})
	q.Run(c)

	return img

}
func getTitle(link string) string {
	// Instantiate a new collector
	c := colly.NewCollector()
	q, _ := queue.New(
		15, // Number threads
		&queue.InMemoryQueueStorage{MaxSize: 200}, 
	)

	q.AddURL("https://en.wikipedia.org" + link)

	// Find and visit link
	src := ""
	c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
		// Extract title
		src = strings.TrimSpace(e.DOM.Text())
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Scraping finished for", r.Request.URL.String())

	})

	q.Run(c)
	return src
}
func ScrapeLink(node *TreeNode, target string, cache *Cache) {
	// if node.Parent != nil {
	// 	fmt.Println("Scrape :", node.Parent.Link, "  ", node.Link, "  ", node.id)
	// } else {
	// 	fmt.Println("Scrape : ", node.Link, "  ", node.id)

	// }
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

		if value, ok := mapCache.Load(node.Link); ok {
			if value != nil {
				for i := 0; i < len(value.([]*TreeNode)); i++ {
					node.AddChild(NewTreeNode(value.([]*TreeNode)[i].Title, value.([]*TreeNode)[i].Link))
				}
				return
			}

		}
		// rp, err := proxy.RoundRobinProxySwitcher("socks5://158.180.52.194:1080")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		c := colly.NewCollector(
			colly.AllowedDomains("en.wikipedia.org"),
			// colly.Async(true),
		)
		// c.SetProxyFunc(rp)

		q, _ := queue.New(
			15, // Number threads
			&queue.InMemoryQueueStorage{MaxSize: 200}, 
		)

		q.AddURL("https://en.wikipedia.org" + node.Link)

		
		c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
			// Extract title
			node.Title = strings.TrimSpace(e.DOM.Text())
		})
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Extract the href attribute of the <a> element

			link := e.Attr("href")
			teks := e.Text
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

		// callback function 
		c.OnScraped(func(r *colly.Response) {
			if _, ok := mapCache.Load(node.Link); !ok {
				mapCache.Store(node.Link, node.Children)
			}
		})
		q.Run(c)

	}
}
