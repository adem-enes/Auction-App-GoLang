package domains

import (
	"errors"
	"math/rand"
)

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

	if err := report.validate(); err != nil {
		NewReport(report.Price, report.Buyer, report.Seller)
	}

	AuctionReports[report.id] = report
}

func (auctions *Auctions) validate() error {
	if _, ok := AuctionReports[auctions.id]; ok {
		return errors.New("This report already exist")
	}
	return nil
}
