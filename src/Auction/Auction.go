package Auction

import "fmt"

//price		<- price to execute at
//numShares 	<- number of shares to purchase
//timeout	<- number of ticks before the order expires, set to 0 for 
//pFill		<- partial fills of orders, set to false for AON/FOK
type Order struct {
	Price		float
	NumShares	float
	Timeout		int
	PFill		bool
}

type Auction struct {
	BuyBook		[]Order
	SellBook	[]Order
}

//fulfill limit buy orders by placing them on the book
//might need to reference order by pointer
func (x *Auction) partialBuyOrder(y Order) {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen =: len(x.SellBook)

	for co := 0; co < sLen; co++ {
		//break if the sell price ever exceeds the buy price
		if x.SellBook[col].Price > y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if y.NumShares <= 0 {
			break
		}

		//if the first order is smaller, fill it and kill it
		if x.SellBook[col].NumShares < y.NumShares && y.PFill {
			//actually fill the order here
			y.NumShares -= x.SellBook[col].NumShares
			x.SellBook = append(x.SellBook[:col], x.SellBook[col + 1])
			col--
			sLen--
		}

		if !x.PFill {
			//actually fill the order here
			x.SellBook[col].NumShares -= y.NumShares
			break
		}
	}

	if y.NumShares > 0 {
		//search through the buy book and place the order where appropriate
	}
}

func (x *Auction) fullBuyOrder(price int, numShares int, timeout int) {
