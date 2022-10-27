package datatypes

type DownloadFileStruct struct {
	FileName    string            `json:"file_name"`
	FileNameOut string            `json:"file_name_out"`
	BucketName  string            `json:"bucket_name"`
	Credentials CredentialsStruct `json:"credentials"`
}

type UploadFileStruct struct {
	FilePath       string            `json:"file_path"`
	ServerFileName string            `json:"server_file_name"`
	BucketName     string            `json:"bucket_name"`
	Credentials    CredentialsStruct `json:"credentials"`
}

type ListFileStruct struct {
	Prefix      string            `json:"prefix"`
	BucketName  string            `json:"bucket_name"`
	Credentials CredentialsStruct `json:"credentials"`
}

type FileOperationsResponseStruct struct {
	Succesful   bool     `json:"successful"`
	Description string   `json:"description"`
	FileName    string   `json:"file_name"`
	ItemsList   []string `json:"items_list"`
}

type CredentialsStruct struct {
	AWS_ACCESS_KEY_ID     string `json:"aws_access_key"`
	AWS_SECRET_ACCESS_KEY string `json:"aws_secret_key"`
	AWS_S3_REGION         string `json:"region"`
}
