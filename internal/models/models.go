package models

import "time"

// holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the user model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is the room model
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is the restriction model
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the reservation model
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int

	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RommRestrictions is the Roomrestriction model
type RommRestrictions struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StartDate     time.Time
	EndDate       time.Time
	RoomID        string
	Room          Rooms
	ReservationID string
	Reservation   Reservations
	RestrictID    string
	Restriction   Restrictions
}
