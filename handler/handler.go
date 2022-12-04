package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shant3r/lessons-be/db"
)

type Handler struct {
	r *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{r: repository}
}

func (h *Handler) AddProduct(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	product := new(Product)
	err = json.Unmarshal(jsonData, product)
	if err != nil {
		internalError(c, err)
		return
	}
	if product.Name == "" {
		badRequst(c)
		return
	}
	if product.Price <= 0 {
		badRequst(c)
		return
	}
	err = h.r.AddProduct(convertToDBProduct(product))
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) GetProducts(c *gin.Context) {
	idString := c.Request.URL.Query().Get("id")
	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequst(c)
			return
		}
		product, ok := h.getProduct(id)
		if ok {
			statusOk(c, product)
		} else {
			notFound(c)
		}
		return
	}
	products := h.r.GetProducts()

	c.JSON(http.StatusOK, convertToProducts(products))

}

func (h *Handler) UpdateProduct(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	product := new(Product)
	err = json.Unmarshal(jsonData, product)
	if err != nil {
		internalError(c, err)
		return
	}
	if product.Name == "" && product.Price <= 0 {
		badRequst(c)
		return
	}
	if product.Identity <= 0 {
		badRequst(c)
		return
	}

	ok, err := h.r.UpdateProduct(convertToDBProduct(product))
	if err != nil {
		internalError(c, err)
		return
	}
	if !ok {
		notFound(c)
		return
	}

}

func (h *Handler) getProduct(id int64) (*Product, bool) {
	product, ok := h.r.GetProduct(id)
	if ok {
		return convertToProduct(product), true
	}
	return nil, false
}

func convertToProduct(p *db.Product) *Product {
	return &Product{
		Identity: p.ID,
		Name:     p.Title,
		Price:    p.Price,
	}
}

func convertToDBProduct(p *Product) *db.Product {
	return &db.Product{
		Title: p.Name,
		ID:    p.Identity,
		Price: p.Price,
	}
}

func convertToProducts(products []*db.Product) []*Product {
	res := make([]*Product, 0, len(products))
	for _, p := range products {
		res = append(res, convertToProduct(p))
	}
	return res

}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequst(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, "not found")
}

func statusOk(c *gin.Context, val any) {
	c.JSON(http.StatusOK, val)
}

func (h *Handler) AddUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	user := new(User)
	err = json.Unmarshal(jsonData, user)
	if err != nil {
		internalError(c, err)
		return
	}
	if user.UserName == "" {
		badRequst(c)
		return
	}
	if user.UserSurname == "" {
		badRequst(c)
		return
	}
	err = h.r.AddUser(convertToDBUser(user))
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) GetUsers(c *gin.Context) {
	idString := c.Request.URL.Query().Get("id")
	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequst(c)
			return
		}
		user, ok := h.getUser(id)
		if ok {
			statusOk(c, user)
		} else {
			notFound(c)
		}
		return
	}
	users := h.r.GetUsers()

	c.JSON(http.StatusOK, convertToUsers(users))
}

func (h *Handler) getUser(id int64) (*User, bool) {
	user, ok := h.r.GetUser(id)
	if ok {
		return convertToUser(user), true
	}
	return nil, false
}

func convertToDBUser(p *User) *db.User {
	return &db.User{
		ID:      p.UserID,
		Name:    p.UserName,
		Surname: p.UserSurname,
	}
}

func convertToUsers(users []*db.User) []*User {
	res := make([]*User, 0, len(users))
	for _, p := range users {
		res = append(res, convertToUser(p))
	}
	return res
}

func convertToUser(p *db.User) *User {
	return &User{
		UserID:      p.ID,
		UserName:    p.Name,
		UserSurname: p.Surname,
	}
}
