package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// println("AWS_ACCESS_KEY_ID     :", os.Getenv("AWS_ACCESS_KEY_ID"))
	// println("AWS_SECRET_ACCESS_KEY :", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)

		// presignClient := s3.NewPresignClient(client, func(s *s3.PresignOptions) {
		// 	s.Expires = time.Until(time.Now().Add(1 * time.Minute))
		// })
		presignClient := s3.NewPresignClient(client)

		responseExpires := time.Now().Add(30 * time.Second) // msec
		request, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket:          aws.String(os.Getenv("AWS_S3_BUCKET")),
			Key:             aws.String(*object.Key),
			ResponseExpires: &responseExpires,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", request.URL)
		fmt.Printf("header : %+v\n", request.SignedHeader)

	}
}

func getContent(method, url string) error {
	return nil
}
