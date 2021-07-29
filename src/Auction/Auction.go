package Auction

import "fmt"

//price		<- price to execute at
//numShares 	<- number of shares to purchase
//timeout	<- number of ticks before the order expires, set to 0 for 
//pFill		<- partial fills of orders, set to false for AON/FOK
//BID		<- broker id
//TID		<- trader id
type Order struct {
	Price		float32
	NumShares	float32
	Timeout		int
	PFill		bool
	BID		int
	TID		int
}

//remember to make the HoldBook in the main function
type Auction struct {
	BuyBook		[]Order
	SellBook	[]Order
	HoldBook	map[int]float32
	Price		int
	SQL		bool
}

//status	<- 0 unplaced | 1 limit placed | 2 filled
//pOrder	<- primary order this has to do with
//fOrder	<- list of orders filled as a result
//cOrder	<- list of orders put on book as a result
type ORM struct {
	status		int
	pOrder		Order
	cOrder		Order
	cOrder		[]Order
}

//records a trade of assets in a book
//-> /x/ itself
//-> (n) number of shares
//-> (p) price of shares
//-> (a) trader id of seller
//-> (b) trader id of buyer
func (x *Auction) recTrade(n float32, p float32, a int, b int) {

}

//fulfill limit buy orders by placing them on the book
//might need to reference order by pointer
//-> /x/ itself
//-> (y) order to be placed
//<- record of exchanges
func (x *Auction) buyOrder(y Order) ORM {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen := len(x.SellBook)
	numShares = y.NumShares
	rec := ORM{pOrder: y}

	for co := 0; co < sLen; co++ {
		//break if the sell price ever exceeds the buy price
		if x.SellBook[col].Price > y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if numShares <= 0 {
			break
		}

		//if the first order is larger and not AON, fill the submitted order
		if x.SellBook[col].NumShares >= numShares && !x.SellBook[col].PFill {
			//record the sale for the buyer and seller
			tX := x.SellBook[col]
			tX.NumShares -= numShares
			rec.fOrder = append(rec.fOrder, tX)
			recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

			//actually fill the order here
			x.SellBook[col].NumShares -= numShares
			break
		}

		//if the first order is smaller, fill it and kill it
		if x.SellBook[col].NumShares < numShares && y.PFill {
			//record the sale for the buyer and seller
			rec.fOrder = append(rec.fOrder, x.SellBook[col])
			recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

			//actually fill the order here
			numShares -= x.SellBook[col].NumShares
			x.SellBook = append(x.SellBook[:col], x.SellBook[col + 1])
			col--
			sLen--
		}
	}

	//put the order on the books if there are more shares remaining
	if numShares > 0 {
		//return unfulfilled if the order timeout is set to 0
		if y.Timeout == 0 {
			//actually do that here
			rec.status = 0

		}

		//search through the buy book and place the order where appropriate
		else {
			rec.status = 1
			newOrder := y
			newOrder.NumShares = numShares
			bLen := len(x.BuyBook)

			for co := 0; co <= bLen; co++ {
				if co == bLen {
					x.BuyBook = append(x.BuyBook, newOrder)
				}

				if x.BuyBook[co].Price > newOrder.Price {
					x.BuyBook = append(x.BuyBook[:co], append(newOrder, x.BuyBook[co:]))
				}
			}

			//add this order to the record
			rec.cOrder = newOrder
		}
	}

	//if no shares remain, pass a flag saying so
	else {
		rec.status = 2
	}

	//pass info on orders created and filled back
	return rec
}

