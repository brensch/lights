package main

import (
	"fmt"

	"github.com/brensch/lights/pkg/wled"
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

	err = s.Power(true)
	if err != nil {
		fmt.Println(err)
	}

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
