wifiNotifier
===

Detect when WiFi changed (connected or disconnected or ssid changed), support Windows, Mac OS, Linux.  
Additional functionality:
- get current WiFi SSID
- scan for a WiFi networks
- get current network cipher (macOS, Windows)

examples

```
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
	fmt.Printf("* current WiFi security: %v\n", wifiNotifier.GetCurrentNetworkSecurity())

	wifiNotifier.SetWifiNotifier(func(ssid string) {
		fmt.Printf("* onWifiChanged. Current ssid: %v (security %v)\n", ssid, wifiNotifier.GetCurrentNetworkSecurity())
	})

	for {
		time.Sleep(time.Second)
	}
}
```
