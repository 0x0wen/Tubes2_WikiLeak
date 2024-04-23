package main

import (
	"fmt"
	"sync"
	"time"
)

// return node, panjang path, jumlah web yang diperiksa
func BFSRace(node *TreeNode, target string, listLink []*TreeNode) (*TreeNode, int, int) {
	// Request the HTML page.
	if node == nil {
		return nil, 0, 0
	}

	// Create a queue for BFS
	queue := []*TreeNode{node}

	// Perform BFS
	found := false
	j := 0
	var wg sync.WaitGroup
	// jumlah web yang diperiksa

	visit := 1
	if node.Link == target {
		return node, 1, visit
	}
	var iter int
	for i := 0; !found; i++ {
		if i == 0 {
			// mutex.Lock()
			ScrapeLink(queue[0], target, listLink)
			queue = append(queue, queue[0].Children...)
			for j := 0; j < len(queue[0].Children); j++ {
				visit += 1
				// if queue[0].Children[j] != nil {
				// 	fmt.Println("BFS: ", queue[0].Children[j].Parent.Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)
				// }
				if queue[0].Children[j].Link == target {
					return queue[0].Children[j], queue[0].Children[j].id + 1, visit
				}
			}
			// mutex.Unlock()
			iter = min(len(queue[0].Children), 10)
		} else if len(queue)-i < iter {

			// if j*17+18 < len(queue) {
			for k := j*iter + 1; k < j*iter+iter+1 && !found; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, listLink)

				}(k)

			}
			wg.Wait()

			for k := j*iter + 1; k < j*iter+iter+1; k++ {
				// mutex.Lock()

				queue = append(queue, queue[k].Children...)
				for j := 0; j < len(queue[k].Children); j++ {
					visit += 1
					// if queue[k].Children[j] != nil {
					// 	fmt.Println("BFS: ", queue[k].Parent.Link, " ", queue[k].Children[j].Link, " ", queue[k].Children[j].id)
					// }
					if queue[k].Children[j].Link == target {
						return queue[k].Children[j], queue[k].Children[j].id + 1, visit
					}
				}

			}
			j += 1

		}

	}

	// Target node not found (impossible)
	return nil, 0, 0

}
func BFSRaceBonus(node *TreeNode, target string, listLink []*TreeNode) ([]*TreeNode, int, int) {
	// Request the HTML page.
	if node == nil {
		return nil, 0, 0
	}

	// Create a queue for BFS
	queue := []*TreeNode{node}
	result := []*TreeNode{}

	// Perform BFS
	j := 0
	depth := -1
	var wg sync.WaitGroup
	visit := 1

	// fmt.Println("BFS: ", queue[0].Link, " ", queue[0].id)
	if node.Link == target {
		if depth == -1 {
			depth = node.id
		}
		result = append(result, node)
	}
	var iter int
	for i := 0; depth == -1 || queue[len(queue)-1].id <= depth; i++ {
		// Dequeue a node from the front of the queue

		// 	scraping
		if i == 0 {
			ScrapeLink(queue[0], target, listLink)
			queue = append(queue, queue[0].Children...)
			for j := 0; j < len(queue[0].Children); j++ {
				visit += 1
				// fmt.Println("BFS: ", queue[0].Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)
				if queue[0].Children[j].Link == target {
					// Modify the title of the target node
					if depth == -1 {
						depth = queue[0].Children[j].id
					}
					result = append(result, queue[0].Children[j])

				}
			}
			iter = min(len(queue[0].Children), 20)
		} else if len(queue)-i <= iter {

			// if j*30+31 < len(queue) {

			for k := j*iter + 1; k < j*iter+iter+1; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, listLink)

				}(k)

			}
			wg.Wait()

			for k := iter*j + 1; k < j*iter+iter+1; k++ {
				queue = append(queue, queue[k].Children...)
				for j := 0; j < len(queue[k].Children); j++ {
					visit += 1
					// fmt.Println("BFS: ", queue[k].Link, " ", queue[k].Children[j].Link, " ", queue[k].Children[j].id)
					if queue[k].Children[j].Link == target {
						// Modify the title of the target node
						if depth == -1 {
							depth = queue[k].Children[j].id
						}
						result = append(result, queue[k].Children[j])

					}
				}

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

	return result, depth + 1, visit

}

// format path: /wiki/Albert_Einstein
func BFS(pathAwal string, pathAkhir string) Result {
	startTime := time.Now()

	root := NewTreeNode("", pathAwal)

	listLink := []*TreeNode{}
	result, length, visit := BFSRace(root, pathAkhir, listLink)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Seconds(), " s")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path lengrh: ", length)
	var path []Website
	for result != nil {
		path = append([]Website{NewWebsite(result.Link, result.Title)}, path...)
		result = result.Parent
	}
	for i := 0; i < len(path); i++ {
		fmt.Println("Title: ", path[i].Title)
		fmt.Println("Link: ", path[i].Link)
	}
	return NewResult(path, length, visit, duration.Seconds())

}

func BFSBonus(pathAwal string, pathAkhir string) ResultBonus {
	startTime := time.Now()

	root := NewTreeNode("", pathAwal)

	listLink := []*TreeNode{}
	result, length, visit := BFSRaceBonus(root, pathAkhir, listLink)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Seconds(), " s")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path lengrh: ", length)
	var pathList [][]Website
	for i := 0; i < len(result); i++ {
		path := []Website{}
		for result[i] != nil {
			path = append([]Website{NewWebsite(result[i].Link, result[i].Title)}, path...)
			result[i] = result[i].Parent
		}
		pathList = append(pathList, path)

	}
	for i := 0; i < len(pathList); i++ {
		fmt.Println("Result ", i+1)
		for j := 0; j < len(pathList[i]); j++ {
			fmt.Println("Title: ", pathList[i][j].Title)
			fmt.Println("Link: ", pathList[i][j].Link)
		}
	}
	return NewResultBonus(pathList, length, visit, duration.Seconds())
}
