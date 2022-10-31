
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

#### Replace object

```http
  POST /api/replaceFile
```
request body : json 

| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `input_file` | `string` | **Required**. file`s name which we delete |
| `output_file` | `string` | **Required**. file`s name which we create |
| `bucket_name` | `string` | **Required**. Bucket name |
| `credentials` | `CredentialStruct` | **Required**. Credential struct |

#### Delete object

```http
  POST /api/deleteObject
```
request body : json 

| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `object_key` | `string` | **Required**. file`s name which we delete |
| `bucket_name` | `string` | **Required**. Bucket name |
| `credentials` | `CredentialStruct` | **Required**. Credential struct |


## CredentialStruct
| Parameter | Type     | Description                  |
| :-------- | :------- | :-------------------------   |
| `aws_access_key` | `string` | **Required**. AWS access key |
| `aws_secret_key` | `string` | **Required**. AWS secret key  |
| `region` | `string` | **Optional**. AWS region |

## Deploy in IIS
All files are by the path `/files for IIS`

First of all, you must install standard IIS module, for example as for 1C publications.

After that create a folder with your .exe file and file "web.config" (example of the file`s content is at the end of this topic). 

Then you can adjust IIS. For that run IIS manager, add a new pool (name may be anyone), parameter ".Net CLR version" change to "No Managed Code" else parameters stay default.

Add a new website, the name may be anyone, the port is the port on which will start your service, choose your new pool, and indicate the path to the folder with your files.

#### Example of web.config
```
<?xml version="1.0" encoding="utf-8"?>
<configuration>
    <system.webServer>
        <handlers>
            <add name="aspNetCore" path="*" verb="*" modules="AspNetCoreModule" resourceType="Unspecified" />
        </handlers>
        <aspNetCore processPath=".\aws_api.exe" />
    </system.webServer>
</configuration>
```
 

## Used By

This project is used by the following companies:

- MediaService
- Gorgany

## Authors

- [@vov_ms](https://github.com/vovms)



