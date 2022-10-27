package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
	data_types "aws_s3_api/datatypes"
	"log"
)

func DownloadHandler(json data_types.DownloadFileStruct) (response data_types.FileOperationsResponseStruct) {

	var jsonAnswer data_types.FileOperationsResponseStruct

	err := aws_component.DownloadFileFromBucket(json.BucketName, json.FileName, json.FileNameOut, json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true
	jsonAnswer.FileName = json.FileNameOut

	return jsonAnswer

}

func ListHandler(json data_types.ListFileStruct) (response data_types.FileOperationsResponseStruct) {

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

func UploadHandler(json data_types.UploadFileStruct) (response data_types.FileOperationsResponseStruct) {

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

func ReplaceHandler(json data_types.ReplaceFileStruct) (response data_types.FileOperationsResponseStruct) {

	var jsonAnswer data_types.FileOperationsResponseStruct

	err := aws_component.CopyObj(json.BucketName, json.InputFile, json.OutputFile, json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	err = aws_component.DeleteObj(json.BucketName, json.InputFile, json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true

	return jsonAnswer

}

func DeleteHandler(json data_types.DeleteObjStruct) (response data_types.FileOperationsResponseStruct) {
	var jsonAnswer data_types.FileOperationsResponseStruct
	err := aws_component.DeleteObj(json.BucketName, json.Key, json.Credentials)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true

	return jsonAnswer
}
