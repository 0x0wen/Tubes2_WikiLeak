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
	var wg sync.WaitGroup
	// jumlah web yang diperiksa

	visit := 1
	node.Title = getTitle(node.Link)
	if node.Link == target {
		return node, 1, visit
	}
	var iter int
	cache := NewCache()
	cache.visited[node.Link] = true
	for !found {
		if queue[len(queue)-1].id == 0 {
			// mutex.Lock()
			wg.Add(1)
			go func() {
				defer wg.Done()
				ScrapeLink(queue[0], target, cache)
			}()
			wg.Wait()
			queue = append(queue, queue[0].Children...)
			queue = queue[1:]
			for j := 0; j < len(queue); j++ {
				visit += 1
				// if queue[0].Children[j] != nil {
				// fmt.Println("BFS ", visit, " : ", queue[j].Parent.Link, " ", queue[j].Link, " ", queue[j].id)
				// }
				if queue[j].Link == target {
					return queue[j], queue[j].id + 1, visit
				}
			}
			// mutex.Unlock()
		} else {
			iter = min(len(queue), 150)
			// if j*17+18 < len(queue) {
			for k := 0; k < iter && !found; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, cache)

				}(k)

			}
			wg.Wait()

			for k := 0; k < iter; k++ {
				// mutex.Lock()

				queue = append(queue, queue[k].Children...)
				for j := 0; j < len(queue[k].Children); j++ {
					visit += 1
					// if queue[k].Children[j] != nil {
					// fmt.Println("BFS ", visit, " : ", queue[k].Link, " ", queue[k].Children[j].Link, " ", queue[k].Children[j].id)
					// }
					if queue[k].Children[j].Link == target {
						return queue[k].Children[j], queue[k].Children[j].id + 1, visit
					}
				}

			}
			queue = queue[iter:]

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
	node.Title = getTitle(node.Link)

	// Create a queue for BFS
	queue := []*TreeNode{node}
	result := []*TreeNode{}

	// Perform BFS
	depth := -1
	var wg sync.WaitGroup
	visit := 1
	cache := NewCache()
	cache.visited[node.Link] = true
	fmt.Println("BFS: ", queue[0].Link, " ", queue[0].id)
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
		if queue[len(queue)-1].id == 0 {
			ScrapeLink(queue[0], target, cache)
			queue = append(queue, queue[0].Children...)
			for j := 0; j < len(queue[0].Children); j++ {
				visit += 1
				// fmt.Println("BFS ", visit, " : ", queue[0].Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)
				if queue[0].Children[j].Link == target {
					// Modify the title of the target node
					if depth == -1 {
						depth = queue[0].Children[j].id
					}
					if !solutionAlreadyExist(queue[0].Children[j], result) {
						result = append(result, queue[0].Children[j])
					}

				}
			}
			queue = queue[1:]
		} else {
			iter = min(len(queue), 150)

			// if j*30+31 < len(queue) {

			for k := 0; k < iter; k++ {
				wg.Add(1)
				go func(k int) {
					defer wg.Done()
					ScrapeLink(queue[k], target, cache)

				}(k)

			}
			wg.Wait()

			for k := 0; k < iter; k++ {
				queue = append(queue, queue[k].Children...)
				for j := 0; j < len(queue[k].Children) && (depth == -1 || depth > queue[k].id); j++ {
					visit += 1
					// fmt.Println("BFS ", visit, " : ", queue[k].Link, " ", queue[k].Children[j].Link, " ", queue[k].Children[j].id)
					if queue[k].Children[j].Link == target {
						// Modify the title of the target node
						if depth == -1 {
							depth = queue[k].Children[j].id
						}
						if !solutionAlreadyExist(queue[k].Children[j], result) {
							result = append(result, queue[k].Children[j])
						}

					}
				}

			}
			queue = queue[iter:]

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
	fmt.Println("Duration: ", duration.Milliseconds(), " ms")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path lengrh: ", length)
	var path []Website
	temp := result
	for temp != nil {
		temp.Title = getTitle(temp.Link)
		temp.imagePath = getImage(temp.Link)
		temp = temp.Parent
	}
	for result != nil {

		path = append([]Website{NewWebsite(result.Link, result.Title, result.imagePath)}, path...)
		result = result.Parent
	}
	// for i := 0; i < len(path); i++ {
	// 	fmt.Println("Title: ", path[i].Title)
	// 	fmt.Println("Link: ", path[i].Link)
	// 	fmt.Println("Img: ", path[i].Imagepath)
	// }
	return NewResult(path, length, visit, duration.Milliseconds())

}

func BFSBonus(pathAwal string, pathAkhir string) ResultBonus {
	startTime := time.Now()

	root := NewTreeNode("", pathAwal)

	listLink := []*TreeNode{}
	result, length, visit := BFSRaceBonus(root, pathAkhir, listLink)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Milliseconds(), " ms")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path lengrh: ", length)
	var wg sync.WaitGroup
	var pathList [][]Website
	for i := 0; i < len(result); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			temp := result[i]
			for temp != nil {

				temp.Title = getTitle(temp.Link)
				temp.imagePath = getImage(temp.Link)
				temp = temp.Parent
			}
		}(i)

	}
	wg.Wait()
	for i := 0; i < len(result); i++ {
		path := []Website{}
		for result[i] != nil {

			path = append([]Website{NewWebsite(result[i].Link, result[i].Title, result[i].imagePath)}, path...)
			result[i] = result[i].Parent
		}
		pathList = append(pathList, path)

	}
	// for i := 0; i < len(pathList); i++ {
	// 	fmt.Println("Result ", i+1)
	// 	for j := 0; j < len(pathList[i]); j++ {
	// 		fmt.Println("Title: ", pathList[i][j].Title)
	// 		fmt.Println("Link: ", pathList[i][j].Link)
	// 		fmt.Println("Image: ", pathList[i][j].Imagepath)
	// 	}
	// }
	return NewResultBonus(pathList, length, visit, duration.Milliseconds())
}
