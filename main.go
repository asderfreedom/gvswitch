package main

import (
	"fmt"
	"gvswitch/dtun"
	"log"
)

func main() {
	config := dtun.Config{
		DeviceType: dtun.TUN,
	}
	config.Name = "dTun0"

	ifce, err := dtun.New(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Interface Name: %s\n", ifce.Name())

	packet := make([]byte, 2000)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %d bytes from interface %s\n", n, ifce.Name())
	}
}
