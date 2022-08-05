package repos

import "github.com/david-luk4s/booking_service/domain/entities"

type RoomRepository interface {
	Get(int32) (*entities.Room, error)
	Save(*entities.Room) error
}

type RoomRepos struct {
	repo RoomRepository
}

func (r *RoomRepos) ServiceGet(id int32) (*entities.Room, error) {
	return r.repo.Get(id)
}

func (r *RoomRepos) ServiceSave(room *entities.Room) error {
	return r.repo.Save(room)
}
