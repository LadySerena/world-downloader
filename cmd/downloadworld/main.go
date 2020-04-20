package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic("could not create cloud storage client")
	}
	bucket := client.Bucket("tiede-minecraft-world-bucket")
	query := &storage.Query{Prefix: ""}
	it := bucket.Objects(ctx, query)
	for {
		objectAttributes, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(objectAttributes.Name)
	}
}
