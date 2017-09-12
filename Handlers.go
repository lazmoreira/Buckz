package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"
)

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := leveldb.OpenFile("./db/buckzdb", nil)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = db.Put([]byte(vars["name"]), []byte(vars["name"]), nil)
}

func CategoryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := leveldb.OpenFile("./db/buckzdb", nil)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	data, err := db.Get([]byte(vars["name"]), nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Data: %s", data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	db, err := leveldb.OpenFile("./db/buckzdb", nil)

	if err != nil {
		return
	}

	defer db.Close()
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
