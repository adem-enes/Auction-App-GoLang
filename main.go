/*
- Struct
	=> Customer
		- Going have idNumber, name, last name, phone, money and products
		- Might have several Products or not even one.
		- Should have at least 1000₺ to attend the auctions..
	>

	=> Product
		- Going to have id, name, type, price, soldPrice, customer
		- Going to have a customer
	>

- Selling a Product
	- We'll start from it's price
	- Next price has to be more then (last given price + 10)₺
	- If customer doesn't have enough money for next price we won't let that customer to make it
	- The auction lasts for the random number between 1 - 7. Last Price wins
	-  We can make the auction like this:
		auction (product, offeredPrice) => {
			- Show The Product's -> Name - Type - Owner - Starting Price
			- Show -> Last Offered Price (Say SOMETHING in the begining)
				(The Auction Has Started)


			- Who's going to make next offer (idNumber)
				- if idNumber exist && have enough money to make next offer
					- Make him/her next offer
					- call auction(product, lastOffer)
				- else
					- This customer not exist || You don't have enough money
		}
		or it might be in a for loop instead of a function..
	- When a customer wins remove this product from seller, add to new owner,
		add sold price to product (last offer)
		and reduce the money of buyer(amount of his/her's last offer)




*/

package main

import (
	"auction-hm2/domains"
	"fmt"
)

var customers = map[int]*domains.Customer{}
var products = map[int]*domains.Product{}

func main() {
	fmt.Println("Hello")
}
