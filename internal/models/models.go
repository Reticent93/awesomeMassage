package models

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type Therapist struct {
	ID            int
	TherapistName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Reservation struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	TherapistID int
	Therapist   Therapist
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Restriction struct {
	ID              int       `json:"id"`
	RestrictionName string    `json:"restriction_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RoomRestriction struct {
	ID            int         `json:"id"`
	StartDate     time.Time   `json:"start_date"`
	EndDate       time.Time   `json:"end_date"`
	TherapistID   int         `json:"therapist_id"`
	ReservationID int         `json:"reservation_id"`
	RestrictionID int         `json:"restriction_id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Therapist     Therapist   `json:"therapist"`
	Reservation   Reservation `json:"reservation"`
	Restriction   Restriction `json:"restriction"`
}

type MassageType struct {
	ID        int       `json:"id"`
	Therapist string    `json:"therapist"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
