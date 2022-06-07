package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Insert new user at database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("No request body data"))
		return
	}

	var user user

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("Erro when user is converted to struct"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		w.Write([]byte("Database connection error"))
		return
	}

	defer db.Close()

	// PREPARE STATEMENT
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Create statement error"))
		return
	}

	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Error when executing statement"))
		return
	}

	lastIdInserted, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Error when get inserted id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User has created successfully! ID: %d", lastIdInserted)))
}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connection()
	if erro != nil {
		w.Write([]byte("Error when connect to database"))
		return
	}
	defer db.Close()

	rows, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Error when searching users"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		if erro := rows.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Erro when scaning data"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro when converting to JSON"))
	}
}

// Get specific user
func GetUser(w http.ResponseWriter, r *http.Request) {

}
