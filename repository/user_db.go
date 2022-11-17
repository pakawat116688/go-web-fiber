package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserdb(db *sqlx.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateDB() error  {
	createdb := `CREATE TABLE "user" (
		"id"	INTEGER,
		"username"	TEXT,
		"password"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`
	
	_, err := r.db.Exec(createdb)
	if err != nil {
		return err
	}

	fmt.Println("Create table success...")
	
	return nil
}

func (r userRepositoryDB) SignUp(user UserCreate) (*User, error) {

	query := `insert into user(username, password) values (?, ?)`

	result, err := r.db.Exec(
		query, 
		user.Username, 
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	users := User{
		Id: int(id),
		Username: user.Username,
		Password: user.Password,
	}

	fmt.Println("Signup success...")

	return &users, nil
}

func (r userRepositoryDB) GetAll() ([]User, error)  {

	query := `select * from user`
	data := []User{}
	err := r.db.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r userRepositoryDB) GetId(uname string) (*User, error)  {
	query := `select * from user where username=?`
	user := User{}
	err := r.db.Get(&user, query, uname)
	if err != nil {
		return nil, err
	}
	fmt.Println("SginIt success...")
	return &user, nil
}