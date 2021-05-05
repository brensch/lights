package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/brensch/lights/pkg/wled"
)

const (
	changeInterval = 1 * time.Minute
)

func main() {
	fmt.Println("yo")

	s, err := wled.InitServer([]string{
		"192.168.1.2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = s
	// this is cool as bro

	// for i := time.Now(); i.Before(time.Now().Add(24 * time.Hour)); i = i.Add(10 * time.Minute) {
	// 	// fmt.Println(i)

	// 	isDark, err := sunutil.TimeIsDark(i)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	fmt.Println("is dark:", i, isDark)
	// }

	// results, err := sunutil.GetSun(targetTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("twilight", results.AstronomicalTwilightEnd.Format(time.RFC3339))
	// fmt.Println(results.NauticalTwilightEnd.Sub(targetTime))

	// fmt.Println(results)

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

	err = s.Power(true)
	if err != nil {
		fmt.Println(err)
	}

	RemoteAddress := "192.168.1.2:21324"

	raddr, err := net.ResolveUDPAddr("udp", RemoteAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Println("Error dialing udp")
		log.Println(err)
	}

	i := 0
	fmt.Println("starting")

	var leds []*wled.Led

	for i := 0; i <= 1000; i++ {
		leds = append(leds, &wled.Led{})
	}

	go func() {
		for i := 0; i <= 1000; i++ {
			for _, led := range leds {

				led.SetState(wled.LedState{
					Red: byte(i),
				})

			}
			time.Sleep(10 * time.Millisecond)
		}

	}()

	for {

		i++

		packets := wled.ConstructPackets(leds)
		for _, packet := range packets {
			conn.Write(packet)
			time.Sleep(10 * time.Millisecond)
		}

		// conn.Write([]byte{1, 1, byte(i), 255, 155, 0})

	}

	// err = s.Power(false)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func makeHue(r, g, b, offset byte) []byte {
	currentState := make([]byte, 4+3*255)

	currentState[0] = 4
	currentState[1] = 255
	currentState[2] = 0
	currentState[3] = offset

	for i := 0; i < 255; i++ {
		currentState[4+i*3] = r
		currentState[5+i*3] = g
		currentState[6+i*3] = b
	}

	return currentState
}
