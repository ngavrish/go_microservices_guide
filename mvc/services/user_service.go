package services

import (
	"github.com/ngavrish/go_microservices_guide/mvc/domain"
	"github.com/ngavrish/go_microservices_guide/mvc/utils"
)

// GetUser returns pointer of User object
func GetUser(userID int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userID)
}
