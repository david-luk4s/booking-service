package enums

type StatusPayment int16

const (
	Created StatusPayment = iota
	Paid
	Finished
	Canceled
	Refounded
)
