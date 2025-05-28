package main

import (
	"api/usecase"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Printf("\nIniciando API de contas...\n")

	router := gin.Default()
	router.GET("/contas", func(c *gin.Context) {
		contas := usecase.ObterContas()
		c.JSON(http.StatusOK, contas)
	})

	router.Run("0.0.0.0:9090")
}
