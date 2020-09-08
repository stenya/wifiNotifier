package main

import (
	"log"
	"time"

	"github.com/play175/wifiNotifier"
)

func main() {

	wifiNotifier.SetWifiNotifier(func(ssid string) {
		log.Println("onWifiChanged,current ssid:" + ssid)
	})

	log.Println("current ssid:" + wifiNotifier.GetCurrentSSID())

	for {
		time.Sleep(time.Millisecond * 1000)
	}
}
