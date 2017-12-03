package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rymccue/golang-standard-lib-rest-api/repositories"
	"github.com/rymccue/golang-standard-lib-rest-api/requests"
	"github.com/rymccue/golang-standard-lib-rest-api/utils/caching"
	"github.com/rymccue/golang-standard-lib-rest-api/utils/crypto"
)

type UserController struct {
	DB    *sql.DB
	Cache caching.Cache
}

func NewUserController(db *sql.DB, c caching.Cache) *UserController {
	return &UserController{
		DB:    db,
		Cache: c,
	}
}

func (jc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var rr requests.RegisterRequest
	err := decoder.Decode(&rr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := repositories.CreateUser(jc.DB, rr.Email, rr.Name, rr.Password)
	if err != nil {
		log.Fatalf("Add user to database error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Generate token Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = jc.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(id), oneMonth)
	if err != nil {
		log.Fatalf("Add token to redis Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (jc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var lr requests.LoginRequest
	err := decoder.Decode(&lr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := repositories.GetPrivateUserDetailsByEmail(jc.DB, lr.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	password := crypto.HashPassword(lr.Password, user.Salt)
	if user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}
	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = jc.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(user.ID), oneMonth)
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
