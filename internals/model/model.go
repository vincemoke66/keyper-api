package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Building struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"`
	Name  string    `json:"name"`
	Abbrv string    `json:"abbrv"`
}

type Room struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"`
	Name  string    `json:"name"`
	Floor int       `json:"floor"`
	Building
}

type Key struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	RFID   string    `json:"rfid"`
	Status string    `json:"status"`
	Room
}

type Student struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string    `json:"fist_name"`
	LastName  string    `json:"last_name"`
	SchoolID  string    `json:"school_id"`
	College   string    `json:"college"`
	Course    string    `json:"course"`
}

type Instructor struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	SchoolID  string    `json:"school_id"`
}

type Record struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	Type      string    `json:"type"`
	CreatedAt time.Time
	Key
	Student
}
