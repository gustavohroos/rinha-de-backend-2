package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", dbinfo)
	println("dbinfo: ", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

func setTransactionByClientId(w http.ResponseWriter, r *http.Request) {

}

func getClientStatementById(w http.ResponseWriter, r *http.Request) {

}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/clientes/{id}/transacoes", setTransactionByClientId).Methods("POST")
	r.HandleFunc("/clientes/{id}/extrato", getClientStatementById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
