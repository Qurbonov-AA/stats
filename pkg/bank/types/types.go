package types

type Money int64

type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
