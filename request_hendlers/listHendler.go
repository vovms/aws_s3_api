package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
	data_types "aws_s3_api/datatypes"
	"log"
)

type ListFileStruct struct {
	Prefix     string `json:"prefix"`
	BucketName string `json:"bucket_name"`
}

func ListHendler(json data_types.ListFileStruct) (response data_types.FileOperationsResponseStruct) {

	var jsonResponse data_types.FileOperationsResponseStruct

	itemsList, err := aws_component.GetListObjectsInABucket(json.BucketName, json.Prefix, json.Credentials)

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
