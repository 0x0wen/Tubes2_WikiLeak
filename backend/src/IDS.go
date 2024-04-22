package main

import (
	"fmt"
	"sync"
)

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
