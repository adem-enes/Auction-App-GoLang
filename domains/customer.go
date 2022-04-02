package domains

import "errors"

type Customer struct {
	idNumber    int
	Name        string
	LastName    string
	phoneNumber int
	wallet      int
}

var Customers map[int]*Customer

func init() {
	Customers = make(map[int]*Customer)
}

func NewCustomer(IdNumber int, Name, LastName string, PhoneNumber int) (*Customer, error) {
	customer := &Customer{
		idNumber:    IdNumber,
		Name:        Name,
		LastName:    LastName,
		phoneNumber: PhoneNumber,
		wallet:      1000,
	}

	if err := customer.validate(); err != nil {
		return nil, err
	}

	Customers[customer.idNumber] = customer
	return customer, nil
}

func (customer *Customer) validate() error {
	if customer.wallet < 1000 {
		return errors.New("Wallet must be bigger than 1000..")
	}
	if _, ok := Customers[customer.idNumber]; ok {
		return errors.New("The Customer already exist..")
	}
	return nil
}

func CheckCustomer(customerId int) bool {
	if _, ok := Customers[customerId]; ok {
		return true
	} else {
		return false
	}
}

func (customer *Customer) GetId() int {
	return customer.idNumber
}

func (customer *Customer) GetWallet() int {
	return customer.wallet
}
func (customer *Customer) SetWallet(wallet int) {
	customer.wallet = wallet
}

func (customer *Customer) GetPhone() int {
	return customer.phoneNumber
}

func (customer *Customer) SetPhone(phone int) {
	customer.phoneNumber = phone
}
