//2021 Kevin Szmyd | zalatwic

package main

import (
	"fmt"
	auction	"./src/Auction"
	hk	"./src/HK"
)

func main() {
	//market setup
	KOF := auction.Auction{[]hk.Order{}, []hk.Order{}, make(map[int]float32), []hk.Record{}, 50.0, 0, false}
	kcc := make(chan hk.BAC)
	krc := make(chan hk.BAR)

	//control messages
	//exit the matket
	qq := hk.BAC{Type: 7, Pike: krc}

	//place a buy order
	t1 := hk.BAC{Type: 2, Pike: krc, Blood: hk.Order{32.58, 10, 1000000, false, 1, 1}}

	//place a low sell order
	t2 := hk.BAC{Type: 1, Pike: krc, Blood: hk.Order{10.58, 10, 1000000, false, 1, 1}}

	//get the order book
	//ob := hk.BAC{Type: 6, Pike: krc}

	go KOF.Open(kcc)
	kcc <- t2
	fmt.Println(<-krc)
	kcc <- t1
	fmt.Println(<-krc)
	kcc <- qq
	fmt.Println(<-krc)
}
