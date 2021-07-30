//2021 Kevin Szmyd | zalatwic

package Auction

import (
	"fmt"
	"time"
)

//internal structs

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

//STID		<- seller's TID
//BTID		<- buyer's TID
//NumShares	<- number of shares
type Record struct {
	NumShares	float32
	Price		float32
	STID		int
	BTID		int
}

//communication structs

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
	newRecord := Record{n, p, a, b}
	x.History = append(x.History, newRecord)
	x.HoldBook[a] -= n
	x.HoldBook[b] += n

	//record sale to sql server
	if x.SQL {
	}
}