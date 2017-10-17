package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"
)

func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	categories := Categories{
		Category{CategoryID: "0001", Name: "Category 1"},
		Category{CategoryID: "0002", Name: "Category 2"},
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		panic(err)
	}
}

//CategoryCreate comment
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	db, err := leveldb.OpenFile(c.Database, nil)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	category := Category{CategoryID: "0002", Name: "Categoria 2", Uri: c.baseUrl() + "category/0002"}

	cat, err := json.Marshal(category)

	if err != nil {
		panic(err)
	}

	if err := db.Put([]byte("cat-"+category.CategoryID), cat, nil); err != nil {
		panic(err)
	}

}

//CategoryShow comment
func CategoryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := leveldb.OpenFile(c.Database, nil)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	categoryName := []byte(vars["name"])

	data, err := db.Get(categoryName, nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Data: %s", data)
}

//Index comment
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Fprintln(w, c.baseUrl())
}
