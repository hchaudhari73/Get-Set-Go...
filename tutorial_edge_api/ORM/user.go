//user API

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func all_users(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: all_users")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)
	json.NewEncoder(w).Encode(users)
}

func new_user(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: new_user")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{
		Name:  name,
		Email: email,
	})
	fmt.Fprintf(w, "New User Successfully Created")
}

func delete_user(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: delete_user")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Successfully deleted user")
}

func update_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: update_user")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	fmt.Fprintf(w, "Successfully updated user")
}
