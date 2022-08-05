package application

import (
	"github.com/david-luk4s/booking_service/adpters/infrastructure/persistence/postgresql"
	"github.com/david-luk4s/booking_service/config/database"
	"github.com/david-luk4s/booking_service/domain/entities"
	"github.com/david-luk4s/booking_service/domain/ports/repos"
)

func GetGuest(id int64) (*entities.Guest, error) {
	pgr := postgresql.NewGuestRepository(database.ConnectionDB())

	rep := &repos.GuestRepos{Repo: pgr}
	guest, err := rep.ServiceGet(id)

	if err != nil {
		return guest, postgresql.ErrorGuestNotFound
	}

	return guest, nil
}

func CreateGuest(id int64) (*entities.Guest, error) {
	pgr := postgresql.NewGuestRepository(database.ConnectionDB())
	repo := &repos.GuestRepos{Repo: pgr}

	guest := entities.NewGuest()
	err := repo.ServiceSave(guest)

	return guest, err
}
