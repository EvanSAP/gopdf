package main

import (
    "github.com/gin-gonic/gin"
)
func main() {
router := gin.Default()
v1 := router.Group("/api/v1/pdfs")
 {
  v1.GET("/", getPDF)
 }
 router.Run()
}