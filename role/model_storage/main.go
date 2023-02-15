package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/model_storage/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		fileType := ".zip"
		_, errByOpenFile := os.Open("models/" + id + fileType)

		if errByOpenFile != nil {
			c.Redirect(http.StatusFound, "/404")
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+id+fileType)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File("models" + "/" + id + fileType)
		// return
	})

	router.Run(":4504")
}
