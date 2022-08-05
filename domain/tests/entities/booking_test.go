package tests

import (
	"testing"
	"time"

	"github.com/david-luk4s/booking_service/domain/entities"
	"github.com/david-luk4s/booking_service/domain/enums"
)

type TestActionChangeStatusBooking struct {
	booking        *entities.Booking
	action         enums.ActionPayment
	expectedStatus enums.StatusPayment
	name           string
}

func TestActionForChangeStatusBooking(t *testing.T) {
	cases := []TestActionChangeStatusBooking{
		{
			name:           "Test action for Pay",
			booking:        &entities.Booking{ID: 1, PlaceId: time.Now(), Start: time.Now(), End: time.Now()},
			action:         enums.Pay,
			expectedStatus: enums.Paid,
		},
		{
			name:           "Test action with status Created for Finished",
			booking:        &entities.Booking{ID: 2, PlaceId: time.Now(), Start: time.Now(), End: time.Now()},
			action:         enums.Finish,
			expectedStatus: enums.Created, //can not Finish if not Paid
		},
		{
			name:           "Test action with status Created for Canceled",
			booking:        &entities.Booking{ID: 3, PlaceId: time.Now(), Start: time.Now(), End: time.Now()},
			action:         enums.Cancel,
			expectedStatus: enums.Canceled,
		},
		{
			name:           "Test action with status Created for Reopen",
			booking:        &entities.Booking{ID: 4, PlaceId: time.Now(), Start: time.Now(), End: time.Now()},
			action:         enums.Reopen,
			expectedStatus: enums.Created, //can not Reopen if not finish
		},
	}

	for i, ct := range cases {
		t.Run(ct.name, func(t *testing.T) {
			if ct.booking.CurrentStatus() != enums.Created {
				t.Errorf("%d° - CurrentStatus FAILED: expected %d, got %d", i+1, enums.Created, ct.booking.CurrentStatus())
			}

			ct.booking.ChangeState(ct.action)
			if ct.booking.CurrentStatus() != ct.expectedStatus {
				t.Errorf("%d° - ChangeState(%d) FAILED: expected %d, got %d", i+1, ct.action, ct.expectedStatus, ct.booking.CurrentStatus())
			}

		})
	}
}
