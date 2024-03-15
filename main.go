package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {

	router := gin.Default()

	// Middleware para permitir CORS
	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.GET("/privacy", privacy)
		v1.GET("/delete-account", deleteAccount)
		v1.GET("/support_dlist", support_dlist)
	}

	router.Run(":3007")

	//// Define o caminho para o certificado e a chave
	//certFile := "/etc/letsencrypt/live/dlist.com.br/fullchain.pem"
	//keyFile := "/etc/letsencrypt/live/dlist.com.br/privkey.pem"
//
	//// Roda o servidor usando TLS
	//if err := router.RunTLS(":3007", certFile, keyFile); err != nil {
	//	panic(err)
	//}

}

func privacy(c *gin.Context) {
	htmlContent, err := os.ReadFile("privacy.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao carregar o arquivo HTML")
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(htmlContent))
}

func deleteAccount(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": "Terms of use"})
	htmlContent, err := os.ReadFile("deleteAccount.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao carregar o arquivo HTML")
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(htmlContent))
}

func support_dlist(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": "Terms of use"})
	htmlContent, err := os.ReadFile("support_dlist.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao carregar o arquivo HTML")
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(htmlContent))
}
