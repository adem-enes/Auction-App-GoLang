package domains

import "math/rand"

type Auctions struct {
	id     int
	Price  int
	Buyer  *Customer
	Seller *Customer
}

var AuctionReports map[int]*Auctions

func init() {
	AuctionReports = make(map[int]*Auctions)
}
func NewReport(Price int, Buyer, Seller *Customer) {
	report := &Auctions{
		id:     rand.Int(),
		Price:  Price,
		Buyer:  Buyer,
		Seller: Seller,
	}

	AuctionReports[report.id] = report
}
