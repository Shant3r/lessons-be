package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	r.GET("/products", func(c *gin.Context) {
		idString := c.Request.URL.Query().Get("id")
		if idString != "" {
			//Конвертировать idString  из строки в int64 (записать в переменную id)
			//Если параметр не число то нужно вернуть статус badrequest (400)
			product, ok := h.GetProduct(id)
			if ok {
				//Вернуть как результат продукт со статусом ок (200)
			}
			//Иначе вернуть notfound(404)
			return
		}
		c.JSON(http.StatusOK, h.GetProducts())
	})
	r.POST("/products", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
		product := new(handler.Product)
		err = json.Unmarshal(jsonData, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
		err = h.AddProduct(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "internal error")
			return
		}
	})

	r.Run()
}
