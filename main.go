package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shant3r/lessons-be/db"
	"github.com/shant3r/lessons-be/handler"
	"github.com/shant3r/lessons-be/handlerUsers"
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
	u := handlerUsers.New(db.New())

	r.GET("/products", h.GetProducts)
	r.POST("/products", h.AddProduct)
	r.PUT("/products", h.UpdateProduct)
	r.GET("/users", u.GetUsers)
	r.POST("/users", u.AddUser)

	r.Run()
}
