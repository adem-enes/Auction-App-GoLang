package main

import (
	"auction-hm2/domains"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

func init() {
	customer1, error := domains.NewCustomer(1, "Adem", "Polat", 5554443322)
	customer2, error := domains.NewCustomer(2, "Enes", "Dinc", 5556663322)
	customer3, error := domains.NewCustomer(3, "Sevket", "Yılmaz", 5557773322)

	_, error1 := domains.NewProduct("Vazo", 100, customer1)
	_, error1 = domains.NewProduct("Araba", 500, customer2)
	_, error1 = domains.NewProduct("Telefon", 300, customer3)

	domains.NewReport(300, customer1, customer2)

	if error != nil {
		fmt.Println(error)
	} else if error1 != nil {
		fmt.Println(error1)
	}
}

func main() {
	menu()
}

func menu() {
	clearScreen()
	fmt.Println("\t\t\t..::Welcome Auction App::..")
	fmt.Println("[1] - See Products \n[2] - See Customers \n[3] - Create New Product")
	fmt.Println("[4] - Create New Customer \n[5] - Auction\n[6] - Auctions Report")
	fmt.Println("[7] - Exit")
	fmt.Print("Your Choise: ")
	var choise int
	fmt.Scan(&choise)

	clearScreen()
	switch choise {
	case 1: //See Products
		seeProducts()
	case 2: //See Customers
		seeCustomers()
	case 3: //Create New Product
		createProduct()
	case 4: //Create New Customer
		createCustomer()
	case 5: // Auction
		productId := getProductFromOwner()
		if _, err := uuid.Parse(productId); err == nil {
			auction(uuid.MustParse(productId))
		} else {
			countDown(2)
			clearScreen()
		}
	case 6: //Auction Reports
		reports()
	case 7: //Exit
		return
	default:
		fmt.Println("Wrong Choise")
	}

	returnMenu()
}
func returnMenu() {
	fmt.Println("\n---------------------------------------------------------------------")

	fmt.Println("\nTo Return Menu Please Press 'Y'")
	var choise string
	fmt.Scan(&choise)
	if choise == "Y" || choise == "y" {
		clearScreen()
		menu()
	}
}
func clearScreen() {
	//To clear console in windows..
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	// To clear console in Mac or Linux
	// fmt.Println("\033[2J")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func seeCustomers() {
	fmt.Println("\nCustomers: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Customer ID \t\t Name LastName\t\t\t Wallet")
	fmt.Println("--------------- \t -----------------\t\t-----------")

	for k, v := range domains.Customers {
		fmt.Print(k)
		fmt.Print("\t\t\t  ", v.Name, " ", v.LastName)
		fmt.Println("\t\t\t ", v.GetWallet())
	}
}
func seeProducts() {
	fmt.Println("\nProducts: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Products ID \t\t\t\t  Name\t\t Price \t\t Owner")
	fmt.Print("-------------------------------------- \t -----------")
	fmt.Println("\t-----------\t---------------")
	for k, v := range domains.Products {
		fmt.Print(k)
		fmt.Print("\t  ", v.Name)
		fmt.Print("   \t  ", v.Price, " ₺")
		fmt.Println("\t\t", v.GetOwner().Name, " ", v.GetOwner().LastName)
	}
}
func reports() {
	fmt.Println("\nAuction Reports: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Report ID \t\t\t  Sold Price\t\t Buyer \t\t\t Seller")
	fmt.Print("----------------------- \t -------------")
	fmt.Println("\t\t-------------\t\t---------------")
	for k, v := range domains.AuctionReports {
		fmt.Print(k)
		fmt.Print("\t\t   ", v.Price, "\t\t\t ", v.Buyer.Name, " ", v.Buyer.LastName)
		fmt.Println("\t\t ", v.Seller.Name, " ", v.Seller.LastName)
	}
}

func createProduct() {
	var name string
	var price int

	fmt.Println("Product Values:")
	fmt.Println("---------------------------------------------------------------------")

	fmt.Print("Product Name: ")
	fmt.Scan(&name)
	fmt.Print("Product Price: ")
	fmt.Scan(&price)

	var ownerId int
	fmt.Println("This Prduct Belonges To:")
	fmt.Print("Customer Id: ")
	fmt.Scan(&ownerId)

	if domains.CheckCustomer(ownerId) {
		domains.NewProduct(name, price, domains.Customers[ownerId])

		fmt.Println("Product Created Successfully..")
	} else {
		fmt.Println("This User Not exist.. Please try again..")
		countDown(2)
		clearScreen()
		createProduct()
	}

}
func createCustomer() {
	var name string
	var lastName string

	fmt.Println("Customer Values:")
	fmt.Println("---------------------------------------------------------------------")

	var idNumber int
	fmt.Print("Customer Id Number:")
	fmt.Scan(&idNumber)

	if !domains.CheckCustomer(idNumber) {
		fmt.Print("Customer Name: ")
		fmt.Scan(&name)
		fmt.Print("Customer LastName: ")
		fmt.Scan(&lastName)

		var phone int
		fmt.Print("Customer Phone: ")
		fmt.Scan(&phone)

		domains.NewCustomer(idNumber, name, lastName, phone)
		fmt.Println("Customer Created Successfully..")
	} else {
		fmt.Println("This User Already exist.. Please try again..")
		countDown(2)
		clearScreen()
		createCustomer()
	}
}
func getProductFromOwner() string {
	var idNumber int
	fmt.Print("Customer Id Number: ")
	fmt.Scan(&idNumber)

	if domains.CheckCustomer(idNumber) {
		showMyProducts(idNumber)
		fmt.Println("\n---------------------------------------------------------------------")
		fmt.Println("Choose Product You Want to Sell..")

		var productId string
		fmt.Print("ProductId: ")
		fmt.Scan(&productId)
		if _, err := uuid.Parse(productId); err == nil {
			if product, ok := domains.Products[uuid.MustParse(productId)]; ok && product.GetOwner().GetId() == idNumber {
				return productId
			} else {
				fmt.Println("This Product doesn't belong to",
					domains.Customers[idNumber].Name, " ", domains.Customers[idNumber].LastName)
				return ""
			}
		} else {
			fmt.Println("\n---------------------------------------------------------------------")
			fmt.Println("Invalid Product Id.. Try Again..")
		}
	} else {
		fmt.Println("\n---------------------------------------------------------------------")
		fmt.Println("Invalid Customer Id.. Try Again..")
	}
	return ""
}

func auction(productId uuid.UUID) {
	product, error := domains.CheckProduct(productId)
	auctionCount := rand.Intn(3) + 1
	// fmt.Println(auctionCount)
	var lastGivenPrice int
	lastCustomerId := -1

	if error == nil {
		if product.SoldPrice != 0 {
			product.Price = product.SoldPrice
		}
		fmt.Println("Product: ", product.Name)
		fmt.Println("Owner: ", product.GetOwner().Name, " ", product.GetOwner().LastName)
		fmt.Println("Price: ", product.Price, " ")
		if lastGivenPrice == 0 {
			fmt.Println("The Auction Has Started")
			lastGivenPrice = product.Price
		} else {
			fmt.Println("The Last Price: ", lastGivenPrice)
		}
		fmt.Println("")

		for auctionCount > 0 {
			fmt.Println("---------------------------------------")
			var customerId int
			fmt.Print("\nCustomer Id: ")
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

						fmt.Println("\nSatış Gerçekleşti.. Hayırlı Olsun..\n\n ")
					}
				} else {
					fmt.Println("Invalid Price..")
					auctionCount++
				}
			} else {
				fmt.Println("Invalid Customer..")
				auctionCount++
			}
			auctionCount--
		}
	} else {
		fmt.Println(error)
	}
}

func showMyProducts(customerId int) {
	customer := domains.Customers[customerId]
	fmt.Println("\n", customer.Name, customer.LastName, "'s Products: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Products ID \t\t\t\t  Name\t\t Price \t\t SoldPrice")
	fmt.Print("-------------------------------------- \t -----------")
	fmt.Println("\t-----------\t---------------")
	for k, v := range domains.Products {
		if v.GetOwner().GetId() == customerId {
			fmt.Print(k)
			fmt.Println("\t ", v.Name, "\t\t", v.Price, "\t\t", v.SoldPrice)
		}
	}
}

func countDown(second int) {
	for second > 0 {
		if rand.Intn(100) == 1 {
			break
		}
		fmt.Println("Refreshing in ", second, "s")
		time.Sleep(time.Second)
		second--
	}
	return
}
