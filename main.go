package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { 
		log.Fatalf("fatal error config file: %v", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"code": 200,
			"status": "success",
			"data": "welcome to demo api",
		})
	})
	
	r.GET("/secret", simpleAuthMiddleware(), func(c *gin.Context) {
		secretInformation := viper.GetString("SECRET_INFORMATION")
		c.JSON(http.StatusOK, gin.H {
			"code": http.StatusOK,
			"status": "success",
			"data": secretInformation,
		})
	})

	port := fmt.Sprintf(":%d", viper.GetInt("APP_PORT"))
	r.Run(port)

}

func simpleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("x-api-key")
		if apiKey != "super-secret-key" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"code": http.StatusUnauthorized,
				"status": "error",
				"data": gin.H {
					"error": "unauthorized",
				},
			})
			return
		}
		c.Next()
	}
}