package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("favicon.ico", func(c *gin.Context) {
		data, err := os.ReadFile("web/favicon.ico")

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "image/x-icon", data)
	})

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

	r.GET("/speedtest", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/speedtest/speedtest.sh")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/plain", d)
	})

	r.GET("/weather", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/weather/weather.sh")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/plain", d)
	})

	r.GET("/track", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/track/track.sh")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Data(http.StatusOK, "text/plain", d)
	})

	r.GET("/linux/track", func(c *gin.Context) {
		d, err := os.ReadFile("tools/linux/track/track")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Data(http.StatusOK, "application/octet-stream", d)
	})

	r.Run(":80")
}
