package main

import (
	"fmt"
	"github.com/Cmon-Aa/Sesnrortest/internal/iota"
	"log"
	"time"

	"github.com/shanghuiyang/rpi-devices/dev"
	"github.com/stianeikeland/go-rpio"
)

const (
	pinTrig = 21
	pinEcho = 26
)

func main() {

	// see this for generating a seed: https://docs.iota.org/docs/getting-started/1.1/transfer-tokens/create-a-seed
	iota, err := iota.WithSeed("DCUFIZDXDODUSKAOFKWQGZEWMJQJ9LI9YOPZGYYPKBZERPUMRVLWLXCDUHVUADV9CWBKLLABJLVAKFODR")
	if err != nil {
		panic(err)
	}

	// example usage
	tx, err := iota.SendToTangle("IOTA is fun")
	if err != nil {
		return
	}

	fmt.Printf("tangle transaction: %v\n", tx)
	// see the transaction on https://explorer.iota.org/devnet

	if err := rpio.Open(); err != nil {
		log.Fatalf("failed to open rpio, error: %v", err)
		return
	}
	defer rpio.Close()

	hcsr04 := dev.NewHCSR04(pinTrig, pinEcho)
	for {
		dist := hcsr04.Dist()
		fmt.Printf("%.2f cm\n", dist)
		time.Sleep(1 * time.Second)
	}
}