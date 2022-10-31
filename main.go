package main

import (
	data_types "aws_s3_api/datatypes"
	requesthendlers "aws_s3_api/request_hendlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	debugMode := false

	var r *gin.Engine

	if debugMode {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.GET("api/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "/",
		})
	})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "OK",
		})
	})

	r.POST("/api/downloadFile", func(c *gin.Context) {
		var json data_types.DownloadFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer data_types.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.DownloadHandler(json)
		c.JSON(200, responce)
	})

	r.POST("/api/uploadFile", func(c *gin.Context) {
		var json data_types.UploadFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer data_types.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.UploadHandler(json)
		c.JSON(200, responce)
	})

	r.POST("/api/objectslist", func(c *gin.Context) {
		var json data_types.ListFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer data_types.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}

		responce := requesthendlers.ListHandler(json)
		c.JSON(200, responce)
	})

	r.POST("/api/replaceFile", func(c *gin.Context) {
		var json data_types.ReplaceFileStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer data_types.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.ReplaceHandler(json)
		c.JSON(200, responce)
	})

	r.POST("/api/deleteObject", func(c *gin.Context) {
		var json data_types.DeleteObjStruct
		if err := c.ShouldBindJSON(&json); err != nil {
			var jsonAnswer data_types.FileOperationsResponseStruct
			jsonAnswer.Succesful = false
			jsonAnswer.Description = err.Error()
			c.JSON(http.StatusBadRequest, jsonAnswer)
			return
		}
		responce := requesthendlers.DeleteHandler(json)
		c.JSON(200, responce)
	})

	port := "3050"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)
}
