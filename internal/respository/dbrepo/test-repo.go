package dbrepo

import (
	"errors"
	"time"

	"github.com/Hikari1149/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	if res.RoomID == 2 {
		return 0, errors.New("test error")
	}
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
	if id > 2 {
		return room, errors.New("error : get room by id > 2")
	}
	return room, nil

}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {

	var u models.User
	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 0, "", nil
}
