package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/policy", policy)
		v1.GET("/terms", terms)
	}

	router.Run(":3007")

}

func policy(c *gin.Context) {
	htmlContent, err := os.ReadFile("policy.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao carregar o arquivo HTML")
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(htmlContent))
}

func terms(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Terms of use"})
}
