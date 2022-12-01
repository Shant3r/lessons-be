package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shant3r/lessons-be/db"
	"github.com/shant3r/lessons-be/handler"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	h := handler.New(db.New())

	r.GET("/products", h.GetProducts)
	r.POST("/products", h.AddProduct)
	r.PUT("/products", h.UpdateProduct)

	r.Run()
}
