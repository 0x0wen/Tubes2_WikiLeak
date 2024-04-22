package main

import (
	"fmt"
	"time"
)

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
	// var mutex sync.Mutex
	root := NewTreeNode("", "/wiki/ITB")
	// ScrapeLink(root, "/wiki/Sukarno", &mutex)
	// for i := 0; i < root.GetNumberOfChildren(); i++ {
	// 	ScrapeLink(root.Children[i], "/wiki/Sukarno", &mutex)
	// }
	// root.PrintNode(3)
	listLink := []*TreeNode{}
	result := BFSRace(root, "/wiki/Computer", listLink)
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
