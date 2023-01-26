package main

// Outra forma de fazer

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
)

// Access Size() in Less() to sort by file size.
type ByFileSize []fs.FileInfo

func (a ByFileSize) Len() int           { return len(a) }
func (a ByFileSize) Less(i, j int) bool { return a[i].Size() > a[j].Size() }
func (a ByFileSize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	folder := "teste/"

	// Read directory.
	dirRead, _ := os.Open(folder)
	dirFiles, _ := dirRead.Readdir(0)

	// Sort by file size.
	sort.Sort(ByFileSize(dirFiles))

	// Write results.
	for dirIndex := range dirFiles {
		fileHere := dirFiles[dirIndex]
		fmt.Println(fileHere.Name(), fileHere.Size())
	}

}
