package main

import (
	"flag"
	"log"
	"time"

	"github.com/brensch/lights/pkg/sunutil"
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

	s, err := wled.InitServer([]string{
		"192.168.1.2",
	}, 300)
	if err != nil {
		log.Println(err)
		return
	}

	// check daylight every five minutes and if it's not daylight and not after 12am, pick a random pattern
	ticker := time.NewTicker(changeInterval)

	for {

		select {
		case <-ticker.C:
			isDark, err := sunutil.TimeIsDark(time.Now())
			if err != nil {
				log.Println(err)
				break
			}

			// turn off if it isn't dark or it's after 12am
			if !isDark || time.Now().Hour() < 12 || true {
				err = s.Power(false)
				if err != nil {
					log.Println(err)
				}

				continue
			}

			err = s.Power(true)
			if err != nil {
				log.Println(err)
			}

			err = s.RandomEffect()
			if err != nil {
				log.Println(err)
			}

		}

	}

	// for i := time.Now(); i.Before(time.Now().Add(24 * time.Hour)); i = i.Add(10 * time.Minute) {
	// 	// log.Println(i)

	// 	isDark, err := sunutil.TimeIsDark(i)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}

	// 	log.Println("is dark:", i, isDark)
	// }

	// results, err := sunutil.GetSun(targetTime)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Println("twilight", results.AstronomicalTwilightEnd.Format(time.RFC3339))
	// log.Println(results.NauticalTwilightEnd.Sub(targetTime))

	// log.Println(results)

	// ticker := time.NewTicker(changeInterval)
	// log.Println("changing pattern")

	// s.RandomEffect()

	// for {

	// 	select {
	// 	case <-ticker.C:
	// 		log.Println("changing pattern")
	// 		err = s.RandomEffect()
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 	}

	// }

	// err = s.Power(true)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println("starting")

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
