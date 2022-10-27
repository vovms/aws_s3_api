package aws_component

import (
	"context"
	"log"
	"os"

	data_types "aws_s3_api/datatypes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFileFromBucket(bucketName, fileName, fileNameOut string, Credentials data_types.CredentialsStruct) (err error) {

	// Create the file
	newFile, err := os.Create(fileNameOut)
	if err != nil {
		log.Println(err)
		return err
	}
	defer newFile.Close()

	// Create client s3
	clients3, err := CreateClientS3(Credentials.AWS_ACCESS_KEY_ID, Credentials.AWS_SECRET_ACCESS_KEY, Credentials.AWS_S3_REGION)
	if err != nil {
		return err
	}

	// Create a downloader passing it the S3 client
	downloader := manager.NewDownloader(clients3)
	numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		log.Println(err)
		return err
	}

	_ = numBytes

	return nil

}
