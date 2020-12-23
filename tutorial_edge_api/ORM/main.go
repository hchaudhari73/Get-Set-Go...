/*************************************************************
GORM
- Go-ORM is to interact with sqlite3 database
- ORM (object relationship managers) act almost as brokers
between us developers and our underlying database technology
*************************************************************/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage")
}

func handle_request() {
	new_route := mux.NewRouter().StrictSlash(true)
	new_route.HandleFunc("/", homepage).Methods("GET")
	new_route.HandleFunc("/users", all_users).Methods("GET")
	new_route.HandleFunc("/user/{name}", delete_user).Methods("DELETE")
	new_route.HandleFunc("/user/{name}/{email}", update_user).Methods("PUT")
	new_route.HandleFunc("/user/{name}/{email}", new_user).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", new_route))
}

//initial migration
func initial_migration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	defer db.Close()
	//
	db.AutoMigrate(&User{})
}

func main() {
	initial_migration()

	handle_request()
}
