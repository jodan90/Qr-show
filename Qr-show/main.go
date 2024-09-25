package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/generate", func(c *gin.Context) {
		url := c.PostForm("url")
		png, err := qrcode.Encode(url, qrcode.Medium, 256)
		if err != nil {
			c.String(http.StatusInternalServerError, "QR 코드 생성 중 오류 발생")
			return
		}
		c.Data(http.StatusOK, "image/png", png)
	})

	router.Run(":8080")
}
