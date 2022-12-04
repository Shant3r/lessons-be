package handlerUsers

type User struct {
	UserID      int64  `json:"id"`
	UserName    string `json:"name"`
	UserSurname string `json:"surname"`
}
