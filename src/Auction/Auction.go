//2021 Kevin Szmyd | zalatwic

package Auction

import (
	"fmt"
	"time"
)

//remember to make the HoldBook in the main function
type Auction struct {
	BuyBook		[]Order
	SellBook	[]Order
	HoldBook	map[int]float32
	History		[]Record
	Price		int
	SQL		bool
}

//records a trade of assets in a book
//-> /x/ itself
//-> (n) number of shares
//-> (p) price of shares
//-> (a) trader id of seller
//-> (b) trader id of buyer
func (x *Auction) recTrade(n float32, p float32, a int, b int) {
	newRecord := Record{n, p, a, b}
	x.History = append(x.History, newRecord)
	x.HoldBook[a] -= n
	x.HoldBook[b] += n

	//record sale to sql server
	if x.SQL {
	}
}

//fulfill buy orders
//might need to reference order by pointer
//-> /x/ itself
//-> (y) order to be placed
//<- record of exchanges
func (x *Auction) buyOrder(y Order) ORM {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen := len(x.SellBook)
	cTime :=
	numShares = y.NumShares
	rec := ORM{POrder: y}

	for co := 0; co < sLen; co++ {
		//break if the sell price ever exceeds the buy price
		if x.SellBook[col].Price > y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if numShares <= 0 {
			break
		}

		//ignore old orders
		if x.SellBook[col].Timeout > time.Now().Unix() {
			//if the first order is larger and not AON, fill the submitted order
			if x.SellBook[col].NumShares >= numShares && !x.SellBook[col].PFill {
				//record the sale for the buyer and seller
				tX := x.SellBook[col]
				tX.NumShares -= numShares
				rec.FOrder = append(rec.FOrder, tX)
				recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

				//actually fill the order here
				x.SellBook[col].NumShares -= numShares
				break
			}

			//if the first order is smaller, fill it and kill it
			if x.SellBook[col].NumShares < numShares && y.PFill {
				//record the sale for the buyer and seller
				rec.FOrder = append(rec.FOrder, x.SellBook[col])
				recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

				//actually fill the order here
				numShares -= x.SellBook[col].NumShares
				x.SellBook = append(x.SellBook[:col], x.SellBook[col + 1])
				col--
				sLen--
			}
		}
	}

	//put the order on the books if there are more shares remaining
	if numShares > 0 {
		//return unfulfilled if the order timeout is set to 0
		if y.Timeout == 0 {
			//actually do that here
			rec.Status = 0

		}

		//search through the buy book and place the order where appropriate
		else {
			rec.Status = 1
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
			rec.COrder = newOrder
		}
	}

	//if no shares remain, pass a flag saying so
	else {
		rec.Status = 2
	}

	//pass info on orders created and filled back
	return rec
}

//fulfill sell orders
//might need to reference order by pointer
//-> /x/ itself
//-> (y) order to be placed
//<- record of exchanges
func (x *Auction) sellOrder(y Order) ORM {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen := len(x.BuyBook)
	numShares = y.NumShares
	rec := ORM{POrder: y}

	//verify that the seller has enough shares
	//actually do that here

	for co := 0; co < sLen; co++ {
		//break if the sell price ever exceeds the buy price
		if x.BuyBook[col].Price < y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if numShares <= 0 {
			break
		}

		//ignore old orders
		if x.BuyBook[col].Timeout > time.Now().Unix() {
			//if the first order is larger and not AON, fill the submitted order
			if x.BuyBook[col].NumShares >= numShares && !x.BuyBook[col].PFill {
				//record the sale for the buyer and seller
				tX := x.BuyBook[col]
				tX.NumShares -= numShares
				rec.FOrder = append(rec.FOrder, tX)
				recTrade(numShares, x.BuyBook[col].Price, y.TID, x.BuyBook[col].TID)

				//actually fill the order here
				x.BuyBook[col].NumShares -= numShares
				break
			}

			//if the first order is smaller, fill it and kill it
			if x.BuyBook[col].NumShares < numShares && y.PFill {
				//record the sale for the buyer and seller
				rec.FOrder = append(rec.FOrder, x.BuyBook[col])
				recTrade(numShares, x.BuyBook[col].Price, y.TID, x.BuyBook[col].TID)

				//actually fill the order here
				numShares -= x.BuyBook[col].NumShares
				x.SellBook = append(x.BuyBook[:col], x.BuyBook[col + 1])
				col--
				sLen--
			}
		}
	}

	//put the order on the books if there are more shares remaining
	if numShares > 0 {
		//return unfulfilled if the order timeout is set to 0
		if y.Timeout == 0 {
			//actually do that here
			rec.Status = 0

		}

		//search through the buy book and place the order where appropriate
		else {
			rec.Status = 1
			newOrder := y
			newOrder.NumShares = numShares
			bLen := len(x.SellBook)

			for co := 0; co <= bLen; co++ {
				if co == bLen {
					x.SellBook = append(x.SellBook, newOrder)
				}

				if x.SellBook[co].Price > newOrder.Price {
					x.SellBook = append(x.SellBook[:co], append(newOrder, x.SellBook[co:]))
				}
			}

			//add this order to the record
			rec.COrder = newOrder
		}
	}

	//if no shares remain, pass a flag saying so
	else {
		rec.Status = 2
	}

	//pass info on orders created and filled back
	return rec
}

func (x *Auction) open(com chan []) {
	//create map of current holders

	//enter loop, take commands from brokers
	close := false
	for !close {
		//take a list of commands from the server specified using the passed 
	}
}
