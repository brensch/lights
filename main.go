package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brensch/lights/pkg/wled"
)

const (
	changeInterval = 1 * time.Minute
)

func main() {
	fmt.Println("yo")

	s, err := wled.InitServer([]string{
		"192.168.1.15",
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

	ticker := time.NewTicker(changeInterval)
	log.Println("changing pattern")

	s.RandomEffect()

	for {

		select {
		case <-ticker.C:
			log.Println("changing pattern")
			err = s.RandomEffect()
			if err != nil {
				log.Println(err)
			}
		}

	}

	// err = s.Power(true)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// RemoteAddress := "192.168.1.15:21324"

	// raddr, err := net.ResolveUDPAddr("udp", RemoteAddress)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// conn, err := net.DialUDP("udp", nil, raddr)
	// if err != nil {
	// 	log.Println("Error dialing udp")
	// 	log.Println(err)
	// }

	// i := 0
	// for {

	// 	i++

	// 	// data := makeHue(0, 0, byte(i))
	// 	conn.Write([]byte{1, 1, byte(i), 255, 155, 0})
	// 	time.Sleep(10 * time.Millisecond)

	// 	if i > 255 {
	// 		i = 0
	// 	}

	// }
}

func makeHue(r, g, b byte) []byte {
	currentState := make([]byte, 2+4*255)

	currentState[0] = 1
	currentState[1] = 255
	for i := 0; i < 255; i++ {
		currentState[2+i*4] = byte(i)
		currentState[3+i*4] = r
		currentState[4+i*4] = g
		currentState[5+i*4] = b
	}

	return currentState
}
