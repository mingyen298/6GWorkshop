package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const IP = "172.19.0.1"

// tail -f /proc/7/fd/1
func Unzip(fullpath string) {
	dst := "models"
	archive, err := zip.OpenReader(fullpath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func DownloadNewModel(url string) {
	if _, err := os.Stat("models/model"); os.IsExist(err) {
		os.RemoveAll("models/model")
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while downloading %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	if _, err := os.Stat("models/model.zip"); os.IsExist(err) {
		os.Remove("models/model.zip")
	}
	file, err := os.Create("models/model.zip")
	if err != nil {
		fmt.Printf("Error while creating file: %s\n", err)
		return
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error while downloading %s: %s\n", url, err)
		return
	}
	file.Close()

	Unzip("models/model.zip")

	fmt.Println("Downloaded successfully")
}
func main() {
	// DownloadNewModel("http://ws-model_storage:4504/model_storage/download/1")
	// DownloadNewModel("http://192.168.2.129:4504/model_storage/download/1")
	// exec.Command("/bin/bash", "python", `main.py`, "&").Output()
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/inference_node/model/reload", func(c *gin.Context) {
		DownloadNewModel("http://" + IP + ":4504/model_storage/download/2")

		url := "http://" + IP + ":4502/inference_node/model/reload"
		data := []byte(`{"key1":"value1"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error while sending request: %s\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		defer resp.Body.Close()
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.POST("/inference_node/data/upload", func(c *gin.Context) {

		payload := make(map[string]interface{})
		if err := c.BindJSON(&payload); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		url := "http://" + IP + ":4503/aiml_mitlab/data/upload"

		data, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error while sending request: %s\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		defer resp.Body.Close()
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.Run(":4501")
}
