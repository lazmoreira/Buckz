package main

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type User struct {
	UserID    int32     `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

type Category struct {
	CategoryID     string `json:"categoryId"`
	Name           string `json:"name"`
	ParentCategory int32  `json:"parentCategory"`
	Uri            string `json:"uri"`
}

type Todos []Todo
type Categories []Category
