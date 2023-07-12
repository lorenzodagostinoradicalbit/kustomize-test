package compare

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type PossibleOut int

const (
	NotKustomizeDir PossibleOut = iota
	Different
	Equal
)

func Compare(dir string) PossibleOut {
	kustomizeVers := []string{"v5.0.1", "v4.5.7"}
	dirInputStr := Join(dir)

	if !checkKustomizeFolder(dirInputStr) {
		return NotKustomizeDir
	}

	mountInput := fmt.Sprintf("%s:/test", dirInputStr)

	kustOuts := []string{}

	cmd := "docker"
	for _, version := range kustomizeVers {
		kustDocker := fmt.Sprintf("registry.k8s.io/kustomize/kustomize:%s", version)
		cmdArgs := []string{"run", "--rm", "-v", mountInput, kustDocker, "build", "/test"}
		out, err := exec.Command(cmd, cmdArgs...).Output()
		if err != nil {
			log.Fatal(err)
		}
		kustOuts = append(kustOuts, string(out))
	}
	if kustOuts[0] == kustOuts[1] {
		return Equal
	}
	return Different
}

func checkKustomizeFolder(dirInputStr string) bool {
	// Check for kustomization.yaml file
	files, err := os.ReadDir(dirInputStr)
	if err != nil {
		log.Fatal("Error reading folder", err)
	}
	isKustomize := false
	for _, file := range files {
		if file.Name() == "kustomization.yaml" {
			isKustomize = true
		}
		if isKustomize {
			break
		}
	}

	return isKustomize
}

func Join(target string) string {
	if path.IsAbs(target) {
		return target
	}
	ret, err := filepath.Abs(target)
	if err != nil {
		log.Fatal("Unable to find abs path", err)
	}
	return ret
}
