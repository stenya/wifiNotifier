package main

import (
	"fmt"
	"time"

	"github.com/stenya/wifiNotifier"
)

func main() {

	fmt.Println("* available networks:")
	fmt.Println(wifiNotifier.GetAvailableSSIDs())

	fmt.Printf("* current ssid: %v\n", wifiNotifier.GetCurrentSSID())
	fmt.Printf("* current WiFi is insecure: %v\n", wifiNotifier.GetCurrentNetworkIsInsecure())

	wifiNotifier.SetWifiNotifier(func(ssid string) {
		fmt.Printf("* onWifiChanged. Current ssid: %v (insecure %v)\n", ssid, wifiNotifier.GetCurrentNetworkIsInsecure())
	})

	for {
		time.Sleep(time.Second)
	}
}
