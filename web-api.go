package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//User : Kullanıcı Listesi
type User struct {
	//kullanıcı id
	ID int `json:"id,omitempty"`
	//kullanıcı firstname
	FirstName string `json:"firstname,omitempty"`
	//kullanıcı lastname
	LastName string `json:"lastname,omitempty"`
	//kullanıcı email
	Email string `json:"email,omitempty"`
}

var people []User

func getAllUser(w http.ResponseWriter, req *http.Request) {
	people = nil
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := sql.Open("mysql", "kullanici_adi:sifre@tcp(mysql_sunucu_adresi:port_numarasi)/veritabani_adi?charset=utf8")
	rows, err := db.Query("CALL userlist")
	checkErr(err)

	for rows.Next() {
		var uid int
		var fname string
		var lname string
		var uemail string
		err = rows.Scan(&uid, &fname, &lname, &uemail)
		checkErr(err)
		people = append(people, User{ID: uid, FirstName: fname, LastName: lname, Email: uemail})
	}

	db.Close()
	json.NewEncoder(w).Encode(people)
}

func insertNewUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)

	db, err := sql.Open("mysql", "kullanici_adi:sifre@tcp(mysql_sunucu_adresi:port_numarasi)/veritabani_adi?charset=utf8")
	stmt, err := db.Prepare("INSERT users SET firstname=?,lastname=?,email=?")
	checkErr(err)

	res, err := stmt.Exec(person.FirstName, person.LastName, person.Email)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	db.Close()
	json.NewEncoder(w).Encode(id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getAllUser", getAllUser).Methods("GET")
	router.HandleFunc("/insertUser", insertNewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8035", router))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
