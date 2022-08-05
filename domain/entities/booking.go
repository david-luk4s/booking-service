package entities

import (
	"sync"
	"time"

	"github.com/david-luk4s/booking_service/domain/enums"
	"gorm.io/gorm"
)

type Booking struct {
	sync.Mutex
	gorm.Model
	ID      int32
	PlaceId time.Time
	Start   time.Time
	End     time.Time
	status  enums.StatusPayment `default: "enums.Created"`
}

func (b *Booking) CurrentStatus() enums.StatusPayment {
	return b.status
}

func (b *Booking) ChangeState(action enums.ActionPayment) {
	switch action {
	case enums.Pay:
		if b.status == enums.Created {
			b.Lock()
			b.status = enums.Paid
			b.Unlock()
		}
	case enums.Cancel:
		if b.status == enums.Created {
			b.Lock()
			b.status = enums.Canceled
			b.Unlock()
		}
	case enums.Finish:
		if b.status == enums.Paid {
			b.Lock()
			b.status = enums.Finished
			b.Unlock()
		}
	case enums.Refound:
		if b.status == enums.Paid {
			b.Lock()
			b.status = enums.Refounded
			b.Unlock()
		}
	case enums.Reopen:
		if b.status == enums.Canceled {
			b.Lock()
			b.status = enums.Created
			b.Unlock()
		}
	}
}
