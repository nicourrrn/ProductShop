package core

type Basket struct {
	id int
	Address string
	Paid, Close bool
	Total int
	Products interface{}
}
