package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "todolist-go/models"
)

func createUser(c *gin.Context) {

}

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}