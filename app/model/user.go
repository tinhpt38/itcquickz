package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model 
	idNumber   string
	fullName   string
	sex string
	bod time.Time
	pod string
	room string
}