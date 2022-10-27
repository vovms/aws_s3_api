package aws_component

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFile(bucketName, fileName, uploadFileFullPath string) (err error, localion string) {

	file, err := os.Open(uploadFileFullPath) // For read access.
	if err != nil {
		log.Println(err)
		return err, localion
	}

	// Create client s3
	clients3, err := CreateClientS3()
	if err != nil {
		return err, localion
	}

	uploader := manager.NewUploader(clients3)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		log.Println(err)
		return err, localion
	}

	return nil, result.Location

}
