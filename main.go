package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/shant3r/lessons-be/db"
	"github.com/shant3r/lessons-be/handler"
	"github.com/shant3r/lessons-be/handlerUsers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pwd"
	dbname   = "products"
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

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	repository := db.New(database)

	h := handler.New(repository)
	u := handlerUsers.New(repository)

	r.GET("/products", h.GetProducts)
	r.POST("/products", h.AddProduct)
	r.PUT("/products", h.UpdateProduct)
	r.GET("/users", u.GetUsers)
	r.POST("/users", u.AddUser)

	r.Run()
}
