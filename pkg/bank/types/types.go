package types

type Money int64

type Category string

type Status string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
