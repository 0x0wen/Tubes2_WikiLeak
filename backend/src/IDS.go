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
	node.Title = getTitle(node.Link)

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
		stack := []*TreeNode{node}
		for len(stack) != 0 && !found {
			iter = min(len(stack), 150)

			if stack[0].id == 0 {
				for j := 0; j < len(stack[0].Children); j++ {
					// fmt.Println("IDS: ", stack[0].Link, " ", stack[0].Children[j].Link, " ", stack[0].Children[j].id)

					visit += 1
					if stack[0].Children[j].Link == target {
						found = true
						return stack[0].Children[j], stack[0].Children[j].id + 1, visit
					}
				}
				stack = stack[1:]
				if i != 1 {
					stack = append(stack, node.Children...)
				}
			} else if stack[0].id < i {
				for j := 0; j < iter; j++ {
					wg.Add(1)
					go func(j int) {
						defer wg.Done()
						if len(stack[j].Children) == 0 {
							ScrapeLink(stack[j], target, cache)
						}
					}(j)
				}
				wg.Wait()
				newNodes := []*TreeNode{}
				for j := 0; j < iter; j++ {
					if stack[j].id == i-1 {
						for k := 0; k < len(stack[j].Children); k++ {
							// fmt.Println("IDS: ", stack[j].Children[k].Parent.Link, " ", stack[j].Children[k].Link, " ", stack[j].Children[k].id)
							visit += 1
							if stack[j].Children[k].Link == target {
								found = true
								return stack[j].Children[k], stack[j].Children[k].id + 1, visit
							}
						}

					} else {
						for k := 0; k < len(stack[j].Children); k++ {
							// fmt.Println("IDS: ", stack[j].Children[k].Parent.Link, " ", stack[j].Children[k].Link, " ", stack[j].Children[k].id)
							visit += 1
							// if stack[j].Children[k].Link == target {
							// 	found = true
							// 	return stack[j].Children[k], stack[j].Children[k].id + 1, visit
							// }

						}
						newNodes = append(newNodes, stack[j].Children...)
					}
				}
				stack = stack[iter:]
				stack = append(newNodes, stack...)
			}
		}

	}
	return nil, 0, 0
}
func IDSRaceBonus(node *TreeNode, target string) ([]*TreeNode, int, int) {

	var wg sync.WaitGroup
	node.Title = getTitle(node.Link)

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
		stack := []*TreeNode{node}
		for len(stack) != 0 && !found {
			iter = min(len(stack), 150)

			if stack[0].id == 0 {
				for j := 0; j < len(stack[0].Children); j++ {
					// fmt.Println("IDS: ", stack[0].Link, " ", stack[0].Children[j].Link, " ", stack[0].Children[j].id)

					visit += 1
					if stack[0].Children[j].Link == target {
						if searchedDepth == -1 {
							searchedDepth = stack[0].Children[j].id
						}
						if !solutionAlreadyExist(stack[0].Children[j], result) {
							result = append(result, stack[0].Children[j])
						}
					}
				}
				stack = stack[1:]
				if i != 1 {
					stack = append(stack, node.Children...)
				}
			} else if stack[0].id < i {
				for j := 0; j < iter; j++ {
					wg.Add(1)
					go func(j int) {
						defer wg.Done()
						if len(stack[j].Children) == 0 {
							ScrapeLink(stack[j], target, cache)
						}
					}(j)
				}
				wg.Wait()
				newNodes := []*TreeNode{}
				for j := 0; j < iter; j++ {
					if stack[j].id == i-1 {
						for k := 0; k < len(stack[j].Children); k++ {
							// fmt.Println("IDS: ", stack[j].Children[k].Parent.Link, " ", stack[j].Children[k].Link, " ", stack[j].Children[k].id)
							visit += 1
							if stack[j].Children[k].Link == target {
								if searchedDepth == -1 {
									searchedDepth = stack[j].Children[k].id
								}
								if !solutionAlreadyExist(stack[j].Children[k], result) {
									result = append(result, stack[j].Children[k])
								}
							}
						}

					} else {
						for k := 0; k < len(stack[j].Children); k++ {
							// fmt.Println("IDS: ", stack[j].Children[k].Parent.Link, " ", stack[j].Children[k].Link, " ", stack[j].Children[k].id)
							visit += 1
							// if stack[j].Children[k].Link == target {
							// 	if searchedDepth == -1 {
							// 		searchedDepth = stack[j].Children[k].id
							// 	}
							// 	result = append(result, stack[j].Children[k])
							// }

						}
						newNodes = append(newNodes, stack[j].Children...)
					}
				}
				stack = stack[iter:]
				stack = append(newNodes, stack...)
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
	// 	fmt.Println("Image: ", path[i].Imagepath)
	// 	fmt.Println("Link: ", path[i].Link)

	// }
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
