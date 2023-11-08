package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		file := "web/index.html"

		d, err := os.ReadFile(file)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Data(http.StatusOK, "text/html", d)
	})

	r.GET("/linux/speedtest", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/speedtest/speedtest-go")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "application/octet-stream", d)
	})

	r.GET("/linux/speedtest.sh", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/speedtest/speedtest.sh")

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Data(http.StatusOK, "text/plain", d)
	})

	r.Run(":80")
}
