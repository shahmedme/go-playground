package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/generate", GenerateHandler)

	db, err := gorm.Open(sqlite.Open("test.sqlite3"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting database")
	}

	db.AutoMigrate(&Product{})

	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.sqlite3"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting database")
	}

	var product Product

	db.Take(&product)

	result, err := json.Marshal(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.sqlite3"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting database")
	}

	db.Create(&Product{Code: "D15", Price: 100})

	fmt.Fprint(w, "Product generated successfully")
}
