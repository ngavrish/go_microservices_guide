package domain

import (
	"fmt"
	"net/http"

	"github.com/ngavrish/go_microservices_guide/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Name1", LastName: "LastName1", Address: "addr1"},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.AppError{
			Message:    fmt.Sprintf("No user found by ID = %v", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}
	return user, nil
}
