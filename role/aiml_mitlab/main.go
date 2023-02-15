package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const IP = "172.19.0.1"

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/aiml_mitlab/model/update/:xAppId", func(c *gin.Context) {
		id := c.Param("xAppId")
		if id == "1" {
			url := "http://" + IP + ":4501/inference_node/model/reload"
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
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}

	})

	router.POST("/aiml_mitlab/data/upload", func(c *gin.Context) {
		payload := make(map[string]interface{})
		if err := c.BindJSON(&payload); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error"})
		}
		data, _ := json.Marshal(payload)
		fmt.Println(string(data))
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.Run(":4503")
}
