package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var svc *s3.S3
var destBucket string

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = s3.New(sess)
	destBucket = os.Getenv("DEST_BUCKET")
}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		sourceBucket := record.S3.Bucket.Name
		item := record.S3.Object.Key
		source := sourceBucket + "/" + item

		log.Printf("Start copying %q from bucket %q to bucket %q\n", item, sourceBucket, destBucket)

		_, err := svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(destBucket), CopySource: aws.String(source), Key: aws.String(item)})
		if err != nil {
			panic(fmt.Errorf("Unable to copy item from bucket %q to bucket %q, %v", sourceBucket, destBucket, err))
		}

		log.Printf("Item %q successfully copied from bucket %q to bucket %q\n", item, sourceBucket, destBucket)
	}
}

func main() {
	lambda.Start(handler)
}
