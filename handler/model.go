package handler

type Product struct {
	Identity int64   `json:"id"`
	Name     string  `json:"title"`
	Price    float64 `json:"price"`
}

type User struct {
	UserID      int64  `json:"id"`
	UserName    string `json:"name"`
	UserSurname string `json:"surname"`
}
