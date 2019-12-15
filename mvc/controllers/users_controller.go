package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ngavrish/go_microservices_guide/mvc/services"
)

// GetUser - get single user details
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	log.Printf("User ID = %v", userId)
	if err != nil {
		// return bad request
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("userID must be number"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// handle the error
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
