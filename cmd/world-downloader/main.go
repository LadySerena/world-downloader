package main

import (
	"flag"
	"log"
	"world-downloader/pkg/world"
)

func main() {
	var bucketName string
	flag.StringVar(&bucketName, "bucket", "", "name of your gcp storage bucket")
	flag.Parse()
	if bucketName == "" {
		flag.Usage()
		return
	}
	downloadErr := world.DownloadWorldFromBucket(bucketName)
	if downloadErr != nil {
		log.Println(downloadErr)
		return
	}
	log.Println("downloaded world from: " + bucketName)
}
