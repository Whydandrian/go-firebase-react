package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func CloudStorage() *storage.BucketHandle {
	// [START cloud_storage_golang]
	config := &firebase.Config{
		StorageBucket: "whydandrian-v2.appspot.com",
	}
	opt := option.WithCredentialsFile("serviceAccount.json.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	log.Printf("Created bucket handle: %v\n", bucket)
	return bucket
}

func cloudStorageCustomBucket(app *firebase.App) {
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// [START cloud_storage_custom_bucket_golang]
	bucket, err := client.Bucket("my-custom-bucket")
	// [END cloud_storage_custom_bucket_golang]
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Created bucket handle: %v\n", bucket)
}