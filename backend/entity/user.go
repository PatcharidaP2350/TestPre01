package entity

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string
	LastName       string
	Age            int
	BirthDate      time.Time
	PhoneNumber    string
	Weight         float32
	Height         float32
	Email     string	

}