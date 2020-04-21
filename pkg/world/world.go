package world

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const filename = "world.tar.gz"

type byTime []storage.ObjectAttrs

func (objects byTime) Len() int {
	return len(objects)
}

func (objects byTime) Swap(i, j int) {
	objects[i], objects[j] = objects[j], objects[i]
}

func (objects byTime) Less(i, j int) bool {
	return objects[i].Created.Before(objects[j].Created)
}

func DownloadWorldFromBucket(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic("could not create cloud storage client")
	}
	defer client.Close()
	bucket := client.Bucket(bucketName)
	query := &storage.Query{Prefix: ""}
	it := bucket.Objects(ctx, query)
	objectsByTime := byTime{}
	for {
		objectAttributes, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		objectsByTime = append(objectsByTime, *objectAttributes)
	}
	sort.Sort(sort.Reverse(objectsByTime))
	object := bucket.Object(objectsByTime[0].Name)
	reader, readerErr := object.NewReader(ctx)
	if readerErr != nil {
		log.Println(readerErr.Error())
		return readerErr
	}
	defer reader.Close()

	buffer := new(bytes.Buffer)
	_, bufferErr := buffer.ReadFrom(reader)
	if bufferErr != nil {
		log.Println(bufferErr.Error())
		return bufferErr
	}
	writeErr := ioutil.WriteFile(filename, buffer.Bytes(), os.FileMode(0640))
	if writeErr != nil {
		log.Println(writeErr)
	}
	log.Println("downloaded minecraft world!")
	return nil
}
