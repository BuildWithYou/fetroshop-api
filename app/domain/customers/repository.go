package customers

type CustomerRepository interface {
	Create(cst *Customer) error
}
