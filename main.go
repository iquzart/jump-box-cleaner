package main

import (
	"io/fs"
	"jump-box-cleaner/configs"
	"jump-box-cleaner/helpers"
	"jump-box-cleaner/models"
	"jump-box-cleaner/notification"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//generate config
	cfgPath, err := configs.ParseFlags()
	check(err)

	cfg, err := configs.NewConfig(cfgPath)
	check(err)

	// get items in parant directory
	items, err := os.ReadDir(cfg.Cleanup.Path)
	check(err)

	// get directories in parant directory
	directories := getDirectories(items)
	log.Println("Parent Directories: ", directories)

	// create slice of struct
	results := make([]models.ParentDirectories, 0)

	for _, pd := range directories {
		pdPath := (cfg.Cleanup.Path + "/" + pd)
		log.Printf("Checking size of parent directory: %v", pdPath)
		pdSize := dirSize(pdPath)

		items, err := os.ReadDir(pdPath)
		check(err)
		subdirs := getDirectories(items)
		log.Printf("Sub Directories identified for %v: %v ", pd, subdirs)

		//get child directories and size for the pd
		sds := getSubdirs(subdirs, pdPath)

		// update struct with data
		results = append(results,
			models.ParentDirectories{
				Name:           pd,
				Size:           pdSize,
				SubDirectories: sds,
			},
		)

	}
	log.Println("Completed checking directoy sizes")

	log.Println("Send Email Notification")
	notification.EmailNotification(cfg, results)

}

func getSubdirs(subdirs []string, path string) []models.SubDirectories {
	sdResults := make([]models.SubDirectories, 0)

	for _, subd := range subdirs {
		log.Printf("Checking size of subdirectory: %v", path+"/"+subd)
		subdSize := dirSize(path + "/" + subd)
		sdResults = append(sdResults,
			models.SubDirectories{
				Name: subd,
				Size: subdSize,
			},
		)

	}

	return sdResults

}

func getDirectories(items []fs.DirEntry) []string {
	var directories []string
	for _, item := range items {
		if item.IsDir() {
			directories = append(directories, item.Name())
		}
	}
	return directories
}

func dirSize(path string) float64 {
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

	sizeGBF := float64(size) / 1024.0 / 1024.0 / 1024.0

	//sizeGB = math.Round(sizeGB)
	sizeGB := helpers.RoundUp(sizeGBF, 2)
	return sizeGB
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
