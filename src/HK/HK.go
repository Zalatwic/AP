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

//Status	<- 0 unplaced | 1 limit placed | 2 filled
//POrder	<- primary order this has to do with
//COrder	<- order put on book as a result
//FOrder	<- list of orders filled as a result
type ORM struct {
	Status		int
	POrder		Order
	COrder		Order
	FOrder		[]Order
}

