// helper functions
package helpers

import (
	"io/fs"
	"jump-box-cleaner/configs"
	"jump-box-cleaner/models"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/inhies/go-bytesize"
)

// BySize implements sort.Interface based on Size field.
type BySize []models.ParentDirectories

func (a BySize) Len() int {
	return len(a)
}

func (a BySize) Less(i, j int) bool {
	return a[i].Size > a[j].Size
}

func (a BySize) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetParentdirs(directories []string, cfg *configs.Config) []models.ParentDirectories {
	pdResults := make([]models.ParentDirectories, 0)

	for _, pd := range directories {
		pdPath := (cfg.Cleanup.Path + "/" + pd)
		log.Printf("Checking size of parent directory: %v", pdPath)
		pdSize := DirSize(pdPath)

		items, err := os.ReadDir(pdPath)
		Check(err)
		subdirs := GetDirectories(items)
		log.Printf("Sub Directories identified for %v: %v ", pd, subdirs)

		//get child directories and size for the pd
		sds := GetSubdirs(subdirs, pdPath)

		// update struct with data
		pdResults = append(pdResults,
			models.ParentDirectories{
				Name:           pd,
				Size:           pdSize,
				SubDirectories: sds,
			},
		)

	}
	// sort the by size
	sort.Sort(BySize(pdResults))
	return pdResults
}

func GetSubdirs(subdirs []string, path string) []models.SubDirectories {
	sdResults := make([]models.SubDirectories, 0)

	for _, subd := range subdirs {
		log.Printf("Checking size of subdirectory: %v", path+"/"+subd)
		subdSize := DirSize(path + "/" + subd)
		sdResults = append(sdResults,
			models.SubDirectories{
				Name: subd,
				Size: subdSize,
			},
		)

	}

	return sdResults

}

func GetDirectories(items []fs.DirEntry) []string {
	var directories []string
	for _, item := range items {
		if item.IsDir() {
			directories = append(directories, item.Name())
		}
	}
	return directories
}

func DirSize(path string) bytesize.ByteSize {
	sizes := make(chan int64)
	readSize := func(path string, file os.FileInfo, err error) error {
		if err != nil || file == nil {
			return nil // Ignore errors
		}
		if !file.IsDir() {
			sizes <- file.Size()
		}
		return nil
	}

	go func() {
		filepath.Walk(path, readSize)
		close(sizes)
	}()

	size := int64(0)
	for s := range sizes {
		size += s
	}

	return bytesize.New(float64(size))
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
