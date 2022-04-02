package domains

import (
	"errors"
	"math/rand"
)

type Product struct {
	id        int
	Name      string
	Type      string
	Price     int
	SoldPrice int
	owner     *Customer
}

var Products map[int]*Product

func init() {
	Products = make(map[int]*Product)
}

func NewProduct(Name, Type string, Price int, Owner *Customer) (*Product, error) {
	product := &Product{
		id:        rand.Int(),
		Name:      Name,
		Type:      Type,
		Price:     Price,
		owner:     Owner,
		SoldPrice: 0, //If haven't been sold
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	Products[product.id] = product

	return product, nil
}

func (product *Product) validate() error {
	if _, ok := Products[product.id]; ok {
		return errors.New("Product couldn't created. Please Try Again..")
	}
	if product.Price <= 0 {
		return errors.New("Price must be valid..")
	}
	if product.Name == "" {
		return errors.New("Name space can't be empty..")
	}

	return nil
}

func (product *Product) GetOwner() *Customer {
	return product.owner
}

func CheckProduct(productId int) (*Product, error) {

	product, ok := Products[productId]
	if !ok {
		err := errors.New("This product is not exist..")

		return nil, err
	}

	return product, nil

}

func (product *Product) Sell(Price int, newOwner *Customer) {
	NewReport(Price, newOwner, product.owner)

	product.owner.SetWallet(product.owner.GetWallet() + Price)
	newOwner.SetWallet((newOwner.GetWallet() - Price))

	product.owner = newOwner
	product.SoldPrice = Price
}
