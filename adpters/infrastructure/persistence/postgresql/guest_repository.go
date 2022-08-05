package postgresql

import (
	"errors"

	"github.com/david-luk4s/booking_service/domain/entities"
	"gorm.io/gorm"
)

var (
	ErrorGuestNotFound  = errors.New("Guest Not Found")
	ErrorGuestNotCreate = errors.New("Guest Not Create")
)

type GuestRepositoryImpl struct {
	Conn *gorm.DB
}

func NewGuestRepository(conn *gorm.DB) *GuestRepositoryImpl {
	return &GuestRepositoryImpl{Conn: conn}
}

func (g *GuestRepositoryImpl) Get(id int64) (*entities.Guest, error) {
	guest := entities.NewGuest()
	g.Conn.First(&guest, id)

	if guest.Model.ID == 0 {
		return guest, ErrorGuestNotFound
	}

	return guest, nil
}

func (g *GuestRepositoryImpl) Save(guest *entities.Guest) error {
	g.Conn.Create(guest)

	if guest.Model.ID == 0 {
		return ErrorGuestNotCreate
	}

	return nil
}
