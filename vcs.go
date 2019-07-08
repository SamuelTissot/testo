package testo

import (
	"fmt"
	"os"
	"path/filepath"
)

//GitRoot travers the directory and report the dir path of the closes .git folder
func GitRoot(root string) (string, error) {
	if len(root) <= 1 && root != "." {
		return root, fmt.Errorf("could not find a vcs root. returning: %s", root)
	}

	_, err := os.Stat(filepath.Join(root, ".git"))

	switch true {
	case err == nil:
		return root, nil
	case os.IsNotExist(err):
		return GitRoot(filepath.Dir(root))
	default:
		return "", err
	}
}
