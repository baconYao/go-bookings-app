package repository

import "github.com/baconYao/bookings-app/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestrictions(r models.RoomRestriction) error
}
