package repos

import "github.com/david-luk4s/booking_service/domain/entities"

type GuestRepository interface {
	Get(int64) (*entities.Guest, error)
	Save(*entities.Guest) error
}

type GuestRepos struct {
	Repo GuestRepository
}

func (s *GuestRepos) ServiceGet(id int64) (*entities.Guest, error) {
	return s.Repo.Get(id)
}

func (s *GuestRepos) ServiceSave(guest *entities.Guest) error {
	return s.Repo.Save(guest)
}
