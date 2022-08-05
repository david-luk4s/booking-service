package enums

type ActionPayment int16

const (
	Pay ActionPayment = iota
	Finish
	Cancel
	Refound
	Reopen
)
