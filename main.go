package main

import (
	requesthendlers "aws_s3_api/request_hendlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "/",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "OK",
		})
	})

	r.POST("/downloadFile", func(c *gin.Context) {
		var json requesthendlers.DownloadFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer requesthendlers.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.DownloadHendler(json.BucketName, json.FileName, json.FileNameOut)
		c.JSON(200, responce)
	})

	r.POST("/uploadFile", func(c *gin.Context) {
		var json requesthendlers.UploadFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer requesthendlers.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.UploadHendler(json.BucketName, json.ServerFileName, json.FilePath)
		c.JSON(200, responce)
	})

	r.POST("/objectslist", func(c *gin.Context) {
		var json requesthendlers.ListFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer requesthendlers.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}

		responce := requesthendlers.ListHendler(json.BucketName, json.Prefix)
		c.JSON(200, responce)
	})

	port := "3050"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)
}
