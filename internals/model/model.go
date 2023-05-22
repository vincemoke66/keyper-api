package model

import (
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
	ID         uuid.UUID `gorm:"type:uuid"`
	Name       string    `json:"name"`
	Floor      int       `json:"floor"`
	BuildingID uuid.UUID `gorm:"foreignkey:BuildingID"`
}

type Key struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	RFID       string    `json:"rfid" gorm:"column:rfid"`
	Status     KeyStatus `json:"status"`
	BuildingID uuid.UUID `gorm:"foreignkey:BuildingID"`
	RoomID     uuid.UUID `gorm:"foreignkey:RoomID"`
}

type KeyStatus string

const (
	KeyStatusAvailable   KeyStatus = "available"
	KeyStatusBorrowed    KeyStatus = "borrowed"
	KeyStatusLost        KeyStatus = "lost"
	KeyStatusUnavailable KeyStatus = "unavailable"
)

type Student struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string    `json:"first_name"`
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
	ID        uuid.UUID  `gorm:"type:uuid"`
	Type      RecordType `json:"type"`
	StudentID uuid.UUID  `gorm:"foreignkey:StudentID"`
	KeyID     uuid.UUID  `gorm:"foreignkey:KeyID"`
	RoomID    uuid.UUID  `gorm:"foreignkey:RoomID"`
}

type RecordType string

const (
	RecordTypeBorrow RecordType = "borrow"
	RecordTypeReturn RecordType = "return"
)
