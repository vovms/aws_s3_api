package aws_component

import (
	data_types "aws_s3_api/datatypes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteObj(bucketName, key string, Credentials data_types.CredentialsStruct) (err error) {

	// Create client s3
	clients3, err := CreateClientS3(Credentials.AWS_ACCESS_KEY_ID, Credentials.AWS_SECRET_ACCESS_KEY, Credentials.AWS_S3_REGION)
	if err != nil {
		return err
	}

	_, err = clients3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Println(err)
	}

	return err

}
