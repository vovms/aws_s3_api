package aws_component

import (
	"context"
	"log"

	data_types "aws_s3_api/datatypes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetListObjectsInABucket(bucketName, prefix string, Credentials data_types.CredentialsStruct) (objectsArray []string, err error) {

	//var BucketName = "gorgany-stoky-na-sajt-1c-test"

	// Create an Amazon S3 service client
	client, err := CreateClientS3(Credentials.AWS_ACCESS_KEY_ID, Credentials.AWS_SECRET_ACCESS_KEY, Credentials.AWS_S3_REGION)

	if err != nil {
		return objectsArray, err
	}

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucketName),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String(""),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return objectsArray, err
		}
		for _, obj := range page.Contents {
			// Do whatever you need with each object "obj"
			objectsArray = append(objectsArray, *obj.Key)
			//log.Println(*obj.Key)
		}
	}

	return objectsArray, err
}
