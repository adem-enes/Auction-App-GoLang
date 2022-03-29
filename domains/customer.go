package domains

/*
- Might add the login logout..

*/

type Customer struct {
	idNumber    int
	Name        string
	LastName    string
	phoneNumber int
	Money       float64
	products    []*Product
}

func NewCustomer(IdNumber int, Name, LastName string, PhoneNumber int, Money float64) Customer {
	return Customer{
		idNumber:    IdNumber,
		Name:        Name,
		LastName:    LastName,
		phoneNumber: PhoneNumber,
		Money:       Money,
	}
}

func (customer *Customer) getProducts() []*Product {
	return customer.products
}
