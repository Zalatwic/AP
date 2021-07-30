//2021 Kevin Szmyd | zalatwic

package auction

import (
	// "fmt"
	"time"
	hk	"../HK"
)

//remember to make the HoldBook in the main function
type Auction struct {
	BuyBook		[]hk.Order
	SellBook	[]hk.Order
	HoldBook	map[int]float32
	History		[]hk.Record
	Price		float32
	Time		int64
	SQL		bool
}

//records a trade of assets in a book
//-> /x/ itself
//-> (n) number of shares
//-> (p) price of shares
//-> (a) trader id of seller
//-> (b) trader id of buyer
func (x *Auction) recTrade(n float32, p float32, a int, b int) {
	newRecord := hk.Record{n, p, a, b}
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
func (x *Auction) buyOrder(y hk.Order) hk.ORM {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen := len(x.SellBook)
	numShares := y.NumShares
	rec := hk.ORM{POrder: y}

	for col := 0; col < sLen; col++ {
		//break if the sell price ever exceeds the buy price
		if x.SellBook[col].Price > y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if numShares <= 0 {
			break
		}

		//ignore old orders
		if x.SellBook[col].Timeout > x.Time {
			//if the first order is larger and not AON, fill the submitted order
			if x.SellBook[col].NumShares >= numShares && !x.SellBook[col].PFill {
				//record the sale for the buyer and seller
				x.Price = x.SellBook[col].Price
				tX := x.SellBook[col]
				tX.NumShares -= numShares
				rec.FOrder = append(rec.FOrder, tX)
				x.recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

				//actually fill the order here
				x.SellBook[col].NumShares -= numShares
				break
			}

			//if the first order is smaller, fill it and kill it
			if x.SellBook[col].NumShares < numShares && y.PFill {
				//record the sale for the buyer and seller
				x.Price = x.SellBook[col].Price
				rec.FOrder = append(rec.FOrder, x.SellBook[col])
				x.recTrade(numShares, x.SellBook[col].Price, x.SellBook[col].TID, y.TID)

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

		} else {
			//search through the buy book and place the order where appropriate
			rec.Status = 1
			newOrder := y
			newOrder.NumShares = numShares
			bLen := len(x.BuyBook)

			for co := 0; co <= bLen; co++ {
				x.BuyBook = append(x.BuyBook, newOrder)
				if x.BuyBook[co].Price > newOrder.Price && co != bLen{
					copy(x.BuyBook[(co + 1):], x.BuyBook[co:])
					x.BuyBook[co] = newOrder
				}
			}

			//add this order to the record
			rec.COrder = newOrder
		}
	} else {
		//if no shares remain, pass a flag saying so
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
func (x *Auction) sellOrder(y hk.Order) hk.ORM {
	//run down the order sheet, write an order if broken with positive shares outstanding
	sLen := len(x.BuyBook)
	numShares := y.NumShares
	rec := hk.ORM{POrder: y}

	//verify that the seller has enough shares
	//actually do that here

	for col := 0; col < sLen; col++ {
		//break if the sell price ever exceeds the buy price
		if x.BuyBook[col].Price < y.Price {
			break
		}

		//break when the number of shares to be purchased hits zero
		if numShares <= 0 {
			break
		}

		//ignore old orders
		if x.BuyBook[col].Timeout > x.Time {
			//if the first order is larger and not AON, fill the submitted order
			if x.BuyBook[col].NumShares >= numShares && !x.BuyBook[col].PFill {
				//record the sale for the buyer and seller
				x.Price = x.BuyBook[col].Price
				tX := x.BuyBook[col]
				tX.NumShares -= numShares
				rec.FOrder = append(rec.FOrder, tX)
				x.recTrade(numShares, x.BuyBook[col].Price, y.TID, x.BuyBook[col].TID)

				//actually fill the order here
				x.BuyBook[col].NumShares -= numShares
				break
			}

			//if the first order is smaller, fill it and kill it
			if x.BuyBook[col].NumShares < numShares && y.PFill {
				//record the sale for the buyer and seller
				x.Price = x.BuyBook[col].Price
				rec.FOrder = append(rec.FOrder, x.BuyBook[col])
				x.recTrade(numShares, x.BuyBook[col].Price, y.TID, x.BuyBook[col].TID)

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

		} else {
			//search through the buy book and place the order where appropriate
			rec.Status = 1
			newOrder := y
			newOrder.NumShares = numShares
			bLen := len(x.SellBook)

			for co := 0; co <= bLen; co++ {
				x.SellBook = append(x.SellBook, newOrder)
				if x.SellBook[co].Price < newOrder.Price && co != bLen{
					copy(x.SellBook[(co + 1):], x.SellBook[co:])
					x.SellBook[co] = newOrder
				}
			}

			//add this order to the record
			rec.COrder = newOrder
		}
	} else {
		//if no shares remain, pass a flag saying so
		rec.Status = 2
	}

	//pass info on orders created and filled back
	return rec
}

//open the auction
//<- /x/ itself
//<- (com) channel open to brokers and controller
func (x *Auction) Open(com chan hk.BAC) {
	sUTime := time.Now().Unix()

	//enter loop, take commands from brokers
	open := true
	for open {
		//get the current local time
		x.Time += (time.Now().Unix() - sUTime)

		//take a list of commands from the server specified using the passed
		event := <-com
		message := hk.BAR{}

		//see hk for type definitions (found under BAP)

		if event.Type == 1 {
			message.Wine = x.sellOrder(event.Blood)
		} else if event.Type == 2 {
			message.Wine = x.buyOrder(event.Blood)
		} else if event.Type == 3 {
			message.History = x.History
		} else if event.Type == 4 {
			message.NavyBook = x.BuyBook
		} else if event.Type == 5 {
			message.NavyBook = x.SellBook
		} else if event.Type == 6 {
			message.NavyBook = append(x.BuyBook, x.SellBook...)
		} else if event.Type == 7 {
			message.History = x.History
			message.NavyBook = append(x.BuyBook, x.SellBook...)
			message.CyanBook = x.HoldBook
			open = false
		} else {
			message.CyanBook = x.HoldBook
		}


		//send a message back with the price
		message.Price = x.Price
		event.Pike <- message
	}
}
