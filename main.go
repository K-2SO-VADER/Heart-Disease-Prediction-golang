package main

import (
	"flag"
)

func main() {

	view := flag.Bool("v", false, "view data in terms of attributes")
	algorithm := flag.String("a", "decision_tree", "the algorithim to be used")
	path := flag.String("p", "./datasets/hungarian.csv", "path of the dataset")

	flag.Parse()

	if *view {
		treeData(*path)

	} else {
		switch *algorithm {
		case "decision_tree":
			tree(*path)
		case "knn":
			kNN(*path)
		case "crossfold":
			crossfold(*path)

		}

	}

}
