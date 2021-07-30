//2021 Kevin Szmyd | zalatwic

package HK

//internal structs

//Price		<- price to execute at
//NumShares 	<- number of shares to purchase
//Timeout	<- number of ticks before the order expires, set to 0 for 
//PFill		<- partial fills of orders, set to false for AON/FOK
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

//Price		<- current price
//Wine		<- order at hand
//History	<- current held history
//NavyBook	<- returned book
//CyanBook	<- holders
type BAR struct {
	Price		int
	Wine		ORM
	History		Record
	NavyBook	[]Order
	CyanBook	map[int]float32
}

//Type		<- 0 price req | 1 sell | 2 buy | 3 history | 4 buy book | 5 sell book | 6 full book | 7 close market
//Pike		<- returning channel for BAR
//Blood		<- order at hand
type BAC struct {
	Type		int
	Pike		<-chan BAR
	Blood		Order
}
