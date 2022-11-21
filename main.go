package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
			id, err := strconv.ParseInt(idString, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, "bad request")
				return
			}
			product, ok := h.GetProduct(id)
			if ok {
				c.JSON(http.StatusOK, product)
			} else {
				c.JSON(http.StatusNotFound, "not found")
			}
			return
		}
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
