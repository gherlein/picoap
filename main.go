package main

import (
	"encoding/hex"
	"log/slog"
	"machine"
	"time"

	"github.com/soypat/cyw43439"
)

const (
	// Set your Wifi SSID and passphrase here
	ssid string = "picoap"
	pass string = "password"
)

func main() {
	// Wait for USB to initialize:
	time.Sleep(time.Second)

	logger := slog.New(slog.NewTextHandler(machine.Serial, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	dev := cyw43439.NewPicoWDevice()
	cfg := cyw43439.DefaultWifiConfig()
	cfg.Logger = logger // Uncomment to see in depth info on wifi device functioning.

	err := dev.Init(cfg)
	if err != nil {
		panic(err)
	}

	err = dev.StartAP(ssid, pass, 1)
	if err != nil {
		panic(err)
	}

	addr, err := dev.HardwareAddr6()
	if err != nil {
		println("no MAC address!")
		panic(err)
	}
	a := addr[:]
	addrs := hex.EncodeToString(a)
	println("MAC:", addrs[0], addrs[1], addrs[2], addrs[3], addrs[4], addrs[5])

	/*
		up := dev.IsLinkUp()
		if up != true {
			println("WiFi Link is NOT up")
			panic(err)
		}
	*/

	for {
		/*
			err = dev.GPIOSet(0, true)
			if err != nil {
				println("err", err.Error())
			} else {
				println("LED ON")
			}
			time.Sleep(500 * time.Millisecond)
			err = dev.GPIOSet(0, false)
			if err != nil {
				println("err", err.Error())
			} else {
				println("LED OFF")
			}
		*/
		time.Sleep(500 * time.Millisecond)
	}
}
