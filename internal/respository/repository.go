package repository

import (
	"time"

	"github.com/Hikari1149/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r models.RommRestriction) error

	SearchAvailabilityByDatesByRoomID(start, end time.Time, rooID int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)

	GetRoomByID(id int) (models.Room, error)
}
