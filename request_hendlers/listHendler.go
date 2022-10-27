package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
	"log"
)

type ListFileStruct struct {
	Prefix     string `json:"prefix"`
	BucketName string `json:"bucket_name"`
}

func ListHendler(bucketName, prefix string) (response FileOperationsResponseStruct) {

	var delimiter = ""
	var jsonResponse FileOperationsResponseStruct

	itemsList, err := aws_component.GetListObjectsInABucket(bucketName, prefix, delimiter)

	if err != nil {
		log.Println(err)
		jsonResponse.Succesful = false
		jsonResponse.Description = err.Error()
		return jsonResponse
	}

	jsonResponse.Succesful = true
	jsonResponse.ItemsList = itemsList

	return jsonResponse
}
