package repository

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserCreate struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	CreateDB() error
	SignUp(UserCreate) (*User, error)
	GetAll() ([]User, error)
	GetId(string) (*User, error)
}
