
# AWS API S3

Opportunities :
 - Download file from bucket
 - Upload file from bucket
 - Get list of files 

## API Reference

#### Download file

```http
  POST /api/downloadFile
```
request body : json 

| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `file_name` | `string` | **Required**. File`s name for server |
| `file_name_out` | `string` | **Required**. Full path to file in the local system |
| `bucket_name` | `string` | **Required**. Bucket name |
| `credentials` | `CredentialStruct` | **Required**. Credential struct |

#### Upload file

```http
  POST /api/uploadFile
```
request body : json 

| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `file_path` | `string` | **Required**. Full path to file in local system |
| `server_file_name` | `string` | **Required**. File`s name in server |
| `bucket_name` | `string` | **Required**. Bucket name |
| `credentials` | `CredentialStruct` | **Required**. Credential struct |

#### Get objects list

```http
  POST /api/objectslist
```
request body : json 

| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `prefix` | `string` | **Required**. Prefix for data filter |
| `bucket_name` | `string` | **Required**. Bucket name |
| `credentials` | `CredentialStruct` | **Required**. Credential struct |


## CredentialStruct
| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `aws_access_key` | `string` | **Required**. AWS access key |
| `aws_secret_key` | `string` | **Required**. AWS secret key  |
| `region` | `string` | **Optional**. AWS region |


## Used By

This project is used by the following companies:

- MediaService
- Gorgany

## Authors

- [@vov_ms](https://github.com/vovms)



