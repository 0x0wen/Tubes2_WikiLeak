package main

import "sync"

type Website struct {
	Link  string
	Title string
}

// list of path
type Result struct {
	Path        []Website
	Pathlength  int
	Pathvisited int
	Duration    int64
}

// List of List of path (khusus bonus)
type ResultBonus struct {
	PathList    [][]Website
	Pathlength  int
	Pathvisited int
	Duration    int64
}

func NewWebsite(link string, title string) Website {
	return Website{
		Link:  link,
		Title: title,
	}

}
func NewResult(path []Website, pathlength int, pathvisited int, duration int64) Result {
	return Result{
		Path:        path,
		Pathlength:  pathlength,
		Pathvisited: pathvisited,
		Duration:    duration,
	}
}

func NewResultBonus(pathlist [][]Website, pathlength int, pathvisited int, duration int64) ResultBonus {
	return ResultBonus{
		PathList:    pathlist,
		Pathlength:  pathlength,
		Pathvisited: pathvisited,
		Duration:    duration,
	}
}

var (
	once     sync.Once
	instance *Singleton
)

type Singleton struct {
	Data map[string][]string
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: make(map[string][]string),
		}
	})
	return instance
}
