wifiNotifier
===

Detect when wifi changed(connected or disconnected or ssid changed),current only support Windows & Mac OS

examples

```
package main

import (
	"log"
	"time"

	"github.com/stenya/wifiNotifier"
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
```