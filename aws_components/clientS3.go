package aws_component

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	AWS_S3_REGION = "eu-central-1" // Region
)

func CreateClientS3() (client *s3.Client, err error) {

	creds := credentials.NewStaticCredentialsProvider("", "", "")

	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(AWS_S3_REGION))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create an S3 client using the loaded configuration
	client = s3.NewFromConfig(cfg)

	return client, nil

}
