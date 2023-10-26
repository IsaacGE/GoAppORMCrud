package types

import (
	"time"
)

type User struct {
	Name         string
	Age          int
	Id           int
	Email        string
	Password     string
	RegisterDate time.Time
}
