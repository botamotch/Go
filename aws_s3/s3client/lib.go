package s3client

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ClientStruct struct {
	Client        *s3.Client
	PresignClient *s3.PresignClient
}

// Load the Shared AWS Configuration (~/.aws/config) and create an Amazon S3 service client
func (c *ClientStruct) Init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	c.Client = s3.NewFromConfig(cfg)
	c.PresignClient = s3.NewPresignClient(c.Client)
}

func (c *ClientStruct) GetListObjects() (*s3.ListObjectsV2Output, error) {
	input := s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
	}
	return c.Client.ListObjectsV2(context.TODO(), &input)
}

func GetObjectKeys() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
	})
	if err != nil {
		return err
	}
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

	return nil
}
