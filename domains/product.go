package domains

type Product struct {
	id        int
	Name      string
	Type      string
	Price     int
	SoldPrice int
	customer  *Customer
}

func NewProduct(id int, Name, Type string, Price int, Customer *Customer) Product {
	return Product{
		id:        id,
		Name:      Name,
		Type:      Type,
		Price:     Price,
		customer:  Customer,
		SoldPrice: 0, //If haven't been sold
	}
}

func (product *Product) getCustomer() *Customer {
	return product.customer
}
