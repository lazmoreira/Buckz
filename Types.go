package main

import "time"

type User struct {
	UserID    int       `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

type Category struct {
	CategoryID     string `json:"categoryId"`
	Name           string `json:"name"`
	ParentCategory int    `json:"parentCategory"`
	URI            string `json:"uri"`
}

//Categories documentation
type Categories []Category
