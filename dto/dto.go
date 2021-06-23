package dto

import (
	"github.com/profile/models"
)

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"` // use another lib/service for unique name conversion
}

func (obj CreateUser) ToDatabase() models.User {
	return models.User{
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Nickname:  obj.Nickname,
		Email:     obj.Email,
		Password:  obj.Password,
		Country:   obj.Country,
	}
}

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func UserFromDatabase(user models.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func (obj UpdateUser) ToDatabase(id uint64) models.User {
	return models.User{
		ID:        id,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Nickname:  obj.Nickname,
		Email:     obj.Email,
		Password:  obj.Password,
		Country:   obj.Country,
	}
}

type Filter struct {
	Country string
}

func (obj Filter) ToDatabase() models.Filter {
	return models.Filter{
		Country: obj.Country,
	}
}
