package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//CategoryIndex documentation
func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	db, err := leveldb.OpenFile(c.Database, nil)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	iterator := db.NewIterator(util.BytesPrefix([]byte("cat-")), nil)

	categories := Categories{}
	category := Category{}

	for iterator.Next() {
		err := json.Unmarshal(iterator.Value(), &category)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
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

	var category Category

	category.CategoryID = r.FormValue("categoryId")
	category.Name = r.FormValue("name")
	category.ParentCategory, _ = strconv.Atoi(r.FormValue("parentCategory"))
	category.URI = c.BaseURL + "category/cat-" + category.CategoryID

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

	categoryID := []byte(vars["categoryId"])

	data, err := db.Get(categoryID, nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Data: %s", data)
}

//CategoryDelete comments
func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := leveldb.OpenFile(c.Database, nil)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	categoryID := []byte(vars["categoryId"])

	err = db.Delete(categoryID, nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Category %s deleted successfully.", categoryID)
}

//Index comment
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Fprintln(w, c.BaseURL)
}
