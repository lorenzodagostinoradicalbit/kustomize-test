package main

import (
	"flag"
	"fmt"
	"testing-kustomize/compare"
	"testing-kustomize/folders"
)

func main() {

	dirInputPtr := flag.String("input-dir", ".", "The kustomize folder to evaluate")

	flag.Parse()

	foldersInput := folders.GetFolders(*dirInputPtr)

	for _, folder := range foldersInput {
		compared := compare.Compare(folder)

		if compared != compare.NotKustomizeDir {
			outTemp := fmt.Sprintf("Comparison in %s: %t", folder, compared == compare.Equal)
			fmt.Println(outTemp)
		}
	}
}
