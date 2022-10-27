package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
	data_types "aws_s3_api/datatypes"
)

func UploadHendler(json data_types.UploadFileStruct) (response data_types.FileOperationsResponseStruct) {

	var jsonAnswer data_types.FileOperationsResponseStruct

	err, localion := aws_component.UploadFile(json.BucketName, json.ServerFileName, json.FilePath, json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true
	jsonAnswer.FileName = localion

	return jsonAnswer

}
