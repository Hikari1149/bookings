package dbrepo

import (
	"time"

	"github.com/Hikari1149/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}
func (m *testDBRepo) InsertRoomRestriction(r models.RommRestriction) error {

	return nil
}

// SearchAvailabilityByDates return true if availability exists for roomID
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, rooID int) (bool, error) {
	return false, nil

}

// SearchAvailabilityForAllRooms return slice of room
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil

}

func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	return room, nil

}
