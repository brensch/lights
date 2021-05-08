package wled

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func InitServer(ips []string, stringLength int) (server *Server, err error) {

	server = &Server{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ips:  ips,
		leds: make([]*Led, stringLength),
	}

	for i := range server.leds {
		server.leds[i] = new(Led)
	}

	for _, ip := range ips {

		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, WARLSPORT))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		server.conns = append(server.conns, conn)

	}

	// start led sync loop
	// TODO: maybe make this able to be turned off
	go func() {
		for {
			server.WriteLedState()
			time.Sleep(1 * time.Millisecond)
		}
	}()

	return
}

func (s *Server) Leds() []*Led {
	s.lock.Lock()
	leds := s.leds
	s.lock.Unlock()
	return leds
}

// WriteLedState blasts state of server to all connected wled devices
func (s *Server) WriteLedState() {
	packets := ConstructPackets(s.Leds())
	for _, packet := range packets {
		for _, conn := range s.conns {
			conn.Write(packet)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
