package helpers

import (
	"path/filepath"
)

func PageDirectories(path string) []string {

	var files [] string // Declaring slice to add html pages to itself

	files,_ = filepath.Glob("views/panel/layouts/*.html") // For Layouts pages

	pathFiles,_ := filepath.Glob("views/" + path + "/*.html") // For Index pages

	for _,file := range pathFiles {
		files = append(files,file)
	}
	return files
}
