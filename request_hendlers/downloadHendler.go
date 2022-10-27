package request_hendlers

import (
	aws_component "aws_s3_api/aws_components"
)

type DownloadFileStruct struct {
	FileName    string `json:"file_name"`
	FileNameOut string `json:"file_name_out"`
	BucketName  string `json:"bucket_name"`
}

type FileOperationsResponseStruct struct {
	Succesful   bool     `json:"successful"`
	Description string   `json:"description"`
	FileName    string   `json:"file_name"`
	ItemsList   []string `json:"items_list"`
}

func DownloadHendler(bucketName, fileName, fileNameOut string) (response FileOperationsResponseStruct) {

	var jsonAnswer FileOperationsResponseStruct

	err := aws_component.DownloadFileFromBucket(bucketName, fileName, fileNameOut)

	if err != nil {
		jsonAnswer.Succesful = false
		jsonAnswer.Description = err.Error()
		return jsonAnswer
	}

	jsonAnswer.Succesful = true
	jsonAnswer.FileName = fileNameOut

	return jsonAnswer

}
