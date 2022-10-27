package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
)

type UploadFileStruct struct {
	FilePath       string `json:"file_path"`
	ServerFileName string `json:"server_file_name"`
	BucketName     string `json:"bucket_name"`
}

func UploadHendler(bucketName, serverFileName, filePath string) (response FileOperationsResponseStruct) {

	var jsonAnswer FileOperationsResponseStruct

	err, localion := aws_component.UploadFile(bucketName, serverFileName, filePath)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true
	jsonAnswer.FileName = localion

	return jsonAnswer

}
