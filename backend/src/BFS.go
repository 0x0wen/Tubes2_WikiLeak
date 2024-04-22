package main

import (
	"fmt"
	"sync"
)

func BFSRace(node *TreeNode, target string, listLink []*TreeNode) *TreeNode {
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
		// mutex.Lock()
		// mutex.Unlock()

		// Check if the current node's link matches the target link
		// if current.Link == target {
		// 	// Modify the title of the target node
		// 	return current // Found the target node, return the modified node
		// }
		// 	scraping
		if i == 0 {
			// mutex.Lock()
			ScrapeLink(queue[0], target, listLink)
			queue = append(queue, queue[0].Children...)
			for j := 0; j < len(queue[0].Children); j++ {
				if queue[0].Children[j] != nil {
					fmt.Println("BFS: ", queue[0].Children[j].Parent.Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)
				}
				if queue[0].Children[j].Link == target {
					return queue[0].Children[j]
				}
			}
			// mutex.Unlock()
		} else if len(queue)-i < 5 {

			// if j*17+18 < len(queue) {
			for k := j*5 + 1; k < j*5+6 && !found; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, listLink)

				}(k)

			}
			wg.Wait()

			for k := j*5 + 1; k < j*5+6; k++ {
				// mutex.Lock()

				queue = append(queue, queue[k].Children...)
				for j := 0; j < len(queue[k].Children); j++ {
					if queue[k].Children[j] != nil {
						fmt.Println("BFS: ", queue[k].Children[j].Parent.Link, " ", queue[k].Children[j].Link, " ", queue[k].Children[j].id)
					}
					if queue[k].Children[j].Link == target {
						return queue[k].Children[j]
					}
				}

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
