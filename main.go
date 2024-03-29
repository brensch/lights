package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/brensch/lights/pkg/wled"
)

var (
	changeInterval = 1 * time.Minute

	IPs = "192.168.1.2" // comma separated IPs
)

func init() {
	flag.StringVar(&IPs, "ips", IPs, "comma separated list of IPs of WLED devices")
}

func main() {

	// the sun api had a bad certificate for like a month meaning the lights weren't coming on.
	// if someone wants to spoof it to make the lights turn on more than they should, good for them.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	s, err := wled.InitServer([]string{
		"192.168.1.2",
	}, 600)
	if err != nil {
		log.Println(err)
		return
	}

	// check daylight every five minutes and if it's not daylight and not after 12am, pick a random pattern
	ticker := time.NewTicker(changeInterval)

	doRandomEffect(s)
	for {
		select {
		case <-ticker.C:
			doRandomEffect(s)
		}

	}

	// go func() {
	// 	lightsOn := 0
	// 	for {

	// 		// lightsOn += rand.Intn(3)*2 - 1
	// 		lightsOn++
	// 		if lightsOn < 0 || lightsOn > len(s.Leds()) {
	// 			lightsOn = 0
	// 		}

	// 		// log.Println(lightsOn)

	// 		for i, led := range s.Leds() {

	// 			if i == lightsOn || i == lightsOn-20 {

	// 				led.SetState(wled.LedState{
	// 					// Red:   byte(rand.Intn(rand.Intn(255) + 1)),
	// 					// Green: byte(rand.Intn(rand.Intn(255) + 1)),
	// 					// Blue:  byte(rand.Intn(rand.Intn(255) + 1)),
	// 					Red:   255,
	// 					Green: 255,
	// 					Blue:  255,
	// 				})
	// 			} else {
	// 				led.SetState(wled.LedState{
	// 					Red:   0,
	// 					Green: 0,
	// 					Blue:  0,
	// 					// Red:   255,
	// 					// Green: 255,
	// 					// Blue:  255,
	// 				})
	// 			}

	// 			time.Sleep(400 * time.Nanosecond)
	// 		}
	// 	}

	// }()
	// time.Sleep(10 * time.Second)

	// // err = s.Power(false)
	// // if err != nil {
	// // 	log.Println(err)
	// // }
}
