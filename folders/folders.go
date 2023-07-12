package folders

import (
	"log"
	"os"
	"path/filepath"
	"testing-kustomize/compare"
)

func GetFolders(dir string) []string {
	ret := []string{}
	dirTrue := compare.Join(dir)
	files, err := os.ReadDir(dirTrue)
	if err != nil {
		log.Fatal("Error reading folder ", err)
	}
	for _, file := range files {
		if file.IsDir() {
			ret = append(ret, filepath.Join(dir, file.Name()))
		}
	}
	return ret
}
