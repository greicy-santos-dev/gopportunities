package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rodando a api em localhost:8080 (padrão)
	router.Run(":8080")
}
