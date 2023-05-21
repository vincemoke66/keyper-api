package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Building struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"`
	Name  string
	Abbrv string
}

type Room struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string
	Building
	Floor int
}

type Key struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	RFID   string
	Status string
	Room
}

type Student struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string
	LastName  string
	SchoolID  string
	College   string
	Course    string
}

type Instructor struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string
	LastName  string
	SchoolID  string
}

type Record struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	Key
	CreatedAt time.Time
	Type      string
	Student
	Instructor
}
