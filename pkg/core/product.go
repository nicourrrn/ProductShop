package core

type Product struct {
	id int
	Cost int
	Name string
	Description string
	Supplier *Supplier
	Categories []string
}
