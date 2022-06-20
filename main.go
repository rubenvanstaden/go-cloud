package main

import (
	"os"
	"fmt"
	"log"
	"context"
	"io/ioutil"

	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
)

func writeBlob(bucket *blob.Bucket) error {

	data, err := ioutil.ReadFile("job1.tar.gz")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

    err = bucket.WriteAll(context.Background(), "my-job.tar.gz", data, nil)
    if err != nil {
        return err
    }

    return nil
}

func readBlob(bucket *blob.Bucket) error {

    data, err := bucket.ReadAll(context.Background(), "my-job")
    if err != nil {
        return err
    }

    fmt.Println(string(data))

    return nil
}

func main()  {

    log.Println("Creating...")

    // The directory you pass to fileblob.OpenBucket must exist first.
    const myDir = "/tmp/go_cloud"
    if err := os.MkdirAll(myDir, 0777); err != nil {
        fmt.Printf("%s", err)
    }

    // Create a file-based bucket.
    bucket, err := fileblob.OpenBucket(myDir, nil)
    if err != nil {
        fmt.Printf("%s", err)
    }
    defer bucket.Close()

    log.Println("Uploading...")

    err = writeBlob(bucket)
    if err != nil {
        log.Println(err)
    }

    log.Println("Reading...")

    err = readBlob(bucket)
    if err != nil {
        log.Println(err)
    }

    log.Println("Ending...")
}
