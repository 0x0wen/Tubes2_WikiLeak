package main

import (
	"fmt"
	"sync"
	"time"
)

func IDSRace(node *TreeNode, target string) (*TreeNode, int, int) {

	var wg sync.WaitGroup

	found := false
	visit := 1

	if node.Link == target {
		return node, 1, visit
	}
	var iter int
	cache := NewCache()
	cache.MarkVisited(node.Link)
	ScrapeLink(node, target, cache)

	for i := 1; !found; i++ {
		// fmt.Println("Depth: ", i)
		visit = 1
		queue := []*TreeNode{node}
		for len(queue) != 0 && !found {
			iter = min(len(queue), 450)

			if queue[0].id == 0 {
				for j := 0; j < len(queue[0].Children); j++ {
					// fmt.Println("IDS: ", queue[0].Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)

					visit += 1
					if queue[0].Children[j].Link == target {
						found = true
						return queue[0].Children[j], queue[0].Children[j].id + 1, visit
					}
				}
				queue = queue[1:]
				if i != 1 {
					queue = append(queue, node.Children...)
				}
			} else if queue[0].id < i {
				for j := 0; j < iter; j++ {
					wg.Add(1)
					go func(j int) {
						defer wg.Done()
						if len(queue[j].Children) == 0 {
							ScrapeLink(queue[j], target, cache)
						}
					}(j)
				}
				wg.Wait()
				newNodes := []*TreeNode{}
				for j := 0; j < iter; j++ {
					if queue[j].id == i-1 {
						for k := 0; k < len(queue[j].Children); k++ {
							// fmt.Println("IDS: ", queue[j].Children[k].Parent.Link, " ", queue[j].Children[k].Link, " ", queue[j].Children[k].id)
							visit += 1
							if queue[j].Children[k].Link == target {
								found = true
								return queue[j].Children[k], queue[j].Children[k].id + 1, visit
							}
						}

					} else {
						for k := 0; k < len(queue[j].Children); k++ {
							// fmt.Println("IDS: ", queue[j].Children[k].Parent.Link, " ", queue[j].Children[k].Link, " ", queue[j].Children[k].id)
							visit += 1
							if queue[j].Children[k].Link == target {
								found = true
								return queue[j].Children[k], queue[j].Children[k].id + 1, visit
							}

						}
						newNodes = append(newNodes, queue[j].Children...)
					}
				}
				queue = queue[iter:]
				queue = append(newNodes, queue...)
			}
		}

	}
	return nil, 0, 0
}
func IDSRaceBonus(node *TreeNode, target string) ([]*TreeNode, int, int) {

	var wg sync.WaitGroup

	found := false
	visit := 1
	result := []*TreeNode{}
	if node.Link == target {
		result = append(result, node)
		return result, 1, visit
	}
	var iter int
	cache := NewCache()
	cache.MarkVisited(node.Link)
	ScrapeLink(node, target, cache)
	searchedDepth := -1
	for i := 1; searchedDepth == -1 || i <= searchedDepth; i++ {
		// fmt.Println("Depth: ", i)
		visit = 1
		queue := []*TreeNode{node}
		for len(queue) != 0 && !found {
			iter = min(len(queue), 150)

			if queue[0].id == 0 {
				for j := 0; j < len(queue[0].Children); j++ {
					// fmt.Println("IDS: ", queue[0].Link, " ", queue[0].Children[j].Link, " ", queue[0].Children[j].id)

					visit += 1
					if queue[0].Children[j].Link == target {
						if searchedDepth == -1 {
							searchedDepth = queue[0].Children[j].id
						}
						if !solutionAlreadyExist(queue[0].Children[j], result) {
							result = append(result, queue[0].Children[j])
						}
					}
				}
				queue = queue[1:]
				if i != 1 {
					queue = append(queue, node.Children...)
				}
			} else if queue[0].id < i {
				for j := 0; j < iter; j++ {
					wg.Add(1)
					go func(j int) {
						defer wg.Done()
						if len(queue[j].Children) == 0 {
							ScrapeLink(queue[j], target, cache)
						}
					}(j)
				}
				wg.Wait()
				newNodes := []*TreeNode{}
				for j := 0; j < iter; j++ {
					if queue[j].id == i-1 {
						for k := 0; k < len(queue[j].Children); k++ {
							// fmt.Println("IDS: ", queue[j].Children[k].Parent.Link, " ", queue[j].Children[k].Link, " ", queue[j].Children[k].id)
							visit += 1
							if queue[j].Children[k].Link == target {
								if searchedDepth == -1 {
									searchedDepth = queue[j].Children[k].id
								}
								if !solutionAlreadyExist(queue[j].Children[k], result) {
									result = append(result, queue[j].Children[k])
								}
							}
						}

					} else {
						for k := 0; k < len(queue[j].Children); k++ {
							// fmt.Println("IDS: ", queue[j].Children[k].Parent.Link, " ", queue[j].Children[k].Link, " ", queue[j].Children[k].id)
							visit += 1
							if queue[j].Children[k].Link == target {
								if searchedDepth == -1 {
									searchedDepth = queue[j].Children[k].id
								}
								result = append(result, queue[j].Children[k])
							}

						}
						newNodes = append(newNodes, queue[j].Children...)
					}
				}
				queue = queue[iter:]
				queue = append(newNodes, queue...)
			}
		}

	}
	return result, searchedDepth + 1, visit
}
func IDS(pathAwal string, pathAkhir string) Result {
	startTime := time.Now()

	root := NewTreeNode("", pathAwal)
	result, length, visit := IDSRace(root, pathAkhir)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Milliseconds(), " ms")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path length: ", length)
	var path []Website
	for result != nil {
		path = append([]Website{NewWebsite(result.Link, getTitle(result.Link), getImage(result.Link))}, path...)
		result = result.Parent
	}
	for i := 0; i < len(path); i++ {
		fmt.Println("Title: ", path[i].Title)
		fmt.Println("Image: ", path[i].Imagepath)
		fmt.Println("Link: ", path[i].Link)

	}
	return NewResult(path, length, visit, duration.Milliseconds())

}
func IDSBonus(pathAwal string, pathAkhir string) ResultBonus {
	startTime := time.Now()

	root := NewTreeNode("", pathAwal)
	result, length, visit := IDSRaceBonus(root, pathAkhir)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration.Milliseconds(), " ms")
	fmt.Println("Total Visit: ", visit)
	fmt.Println("Path lengrh: ", length)
	var pathList [][]Website
	for i := 0; i < len(result); i++ {
		path := []Website{}
		for result[i] != nil {
			path = append([]Website{NewWebsite(result[i].Link, getTitle(result[i].Link), getImage(result[i].Link))}, path...)
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
	return NewResultBonus(pathList, length, visit, duration.Milliseconds())
}
