package handlerUsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shant3r/lessons-be/db"
)

type HandlerUsers struct {
	r *db.Repository
}

func New(repository *db.Repository) *HandlerUsers {
	return &HandlerUsers{r: repository}
}

func (h *HandlerUsers) AddUser(c *gin.Context) {
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

func (h *HandlerUsers) GetUsers(c *gin.Context) {
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

func (h *HandlerUsers) getUser(id int64) (*User, bool) {
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
