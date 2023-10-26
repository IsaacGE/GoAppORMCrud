package models

import "time"

type User struct {
	Id           uint
	Name         string
	Email        string
	Password     string
	PhoneNumber  string
	RegisterDate time.Time
}
