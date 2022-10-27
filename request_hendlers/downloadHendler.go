package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
	data_types "aws_s3_api/datatypes"
)

func DownloadHendler(json data_types.DownloadFileStruct) (response data_types.FileOperationsResponseStruct) {

	var jsonAnswer data_types.FileOperationsResponseStruct

	err := aws_component.DownloadFileFromBucket(json.BucketName, json.FileName, json.FileNameOut,json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true
	jsonAnswer.FileName = json.FileNameOut

	return jsonAnswer

}
