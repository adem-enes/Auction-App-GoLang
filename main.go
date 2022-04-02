package main

import (
	"auction-hm2/domains"
	"fmt"
	"math/rand"
)

func init() {
	adem, error := domains.NewCustomer(1, "Adem", "Polat", 5554443322, 1500)
	enes, error := domains.NewCustomer(2, "Enes", "Dinc", 5556663322, 2000)
	sevket, error := domains.NewCustomer(3, "Sevket", "Yılmaz", 5557773322, 1000)

	urun1, error1 := domains.NewProduct("Vazo", "", 100, adem)
	urun2, error1 := domains.NewProduct("Araba", "", 500, sevket)
	urun3, error1 := domains.NewProduct("Telefon", "", 300, enes)

	if error != nil {
		fmt.Println(error)
	} else if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println(urun1)
		fmt.Println(urun2)
		fmt.Println(urun3)
	}
}

func main() {
	createCustomer()
	createProduct()
	var productId int
	fmt.Print("ProductId: ")
	fmt.Scan(&productId)

	auction(productId)
}

func createCustomer() {

	fmt.Println("\nCustomers.................")
	for k, v := range domains.Customers {
		fmt.Print(k, "\t")
		fmt.Println(v)
	}

}

func createProduct() {
	fmt.Println("\nProducts.................")
	for k, v := range domains.Products {
		fmt.Print(k, "\t")
		fmt.Println(v)
	}
}
func reports() {
	fmt.Println("\nReports.................")
	for k, v := range domains.AuctionReports {
		fmt.Print(k, "\t")
		fmt.Println(v)
	}
}

func auction(productId int) {
	product, error := domains.CheckProduct(productId)
	auctionCount := rand.Intn(3) + 1
	fmt.Println(auctionCount)
	var lastGivenPrice int
	lastCustomerId := -1

	if error == nil {
		for auctionCount > 0 {
			fmt.Println(product.Name, " ", product.Price, " ", product.GetOwner())
			if lastGivenPrice == 0 {
				fmt.Println("The Auction Has Started")
				lastGivenPrice = product.Price
			} else {
				fmt.Println("The Last Price: ", lastGivenPrice)
			}

			var customerId int
			fmt.Print("Customer Id: ")
			fmt.Scan(&customerId)

			if customer, ok := domains.Customers[customerId]; ok && customer.GetId() != product.GetOwner().GetId() &&
				customer.GetId() != lastCustomerId {
				var offer int
				fmt.Print("Cutomer: ")
				fmt.Println(customer.Name, " ", customer.LastName)
				fmt.Print("OfferedPrice: ")
				fmt.Scan(&offer)

				if offer > lastGivenPrice && customer.GetWallet() > offer {
					lastGivenPrice = offer
					lastCustomerId = customer.GetId()

					if auctionCount == 1 {
						product.Sell(lastGivenPrice, customer)

						fmt.Println("Satış Gerçekleşti.. Hayırlı Olsun..\n\n ")
						createCustomer()
						fmt.Println("")
						createProduct()
						fmt.Println("")
						reports()
					}
				} else {
					fmt.Println("Fiyatınız geçerli değildir..")
					auctionCount++
				}
			} else {
				fmt.Println("Geçersiz Müşteri..")
				auctionCount++
			}
			auctionCount--
		}
	} else {
		fmt.Println(error)
	}
}
