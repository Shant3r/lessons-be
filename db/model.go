package db

type Product struct {
	ID    int64
	Title string
	Price float64
}

type User struct {
	ID      int64
	Name    string
	Surname string
}
