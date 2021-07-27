package main

import (
	"jump-box-cleaner/configs"
	"jump-box-cleaner/helpers"
	"jump-box-cleaner/models"
	"jump-box-cleaner/notification"
	"log"
	"os"
	"time"
)

func main() {
	//generate config
	cfgPath, err := configs.ParseFlags()
	helpers.Check(err)

	cfg, err := configs.NewConfig(cfgPath)
	helpers.Check(err)

	// get items in from the path
	items, err := os.ReadDir(cfg.Cleanup.Path)
	helpers.Check(err)

	// filter directories from the path
	directories := helpers.GetDirectories(items)
	log.Println("Parent Directories: ", directories)

	// create slice of struct
	results := make([]models.Data, 0)

	hostname, err := os.Hostname()
	now := time.Now()
	time := now.Format(time.RFC822)
	pds := helpers.GetParentdirs(directories, cfg)

	results = append(results,
		models.Data{
			Hostname:    hostname,
			Date:        time,
			Path:        cfg.Cleanup.Path,
			Directories: pds,
		},
	)

	log.Println("Completed checking directoy sizes")

	log.Println("Send Email Notification")
	notification.EmailNotification(cfg, results)
}
