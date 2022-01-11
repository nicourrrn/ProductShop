package core

type Basket struct {
	Id          int
	Address     string
	Paid, Close bool
	Total       int
	Products    interface{}
}
