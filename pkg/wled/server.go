package wled

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func InitServer(ips []string) (server *Server, err error) {

	server = &Server{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ips: ips,
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

	return
}

func (s *Server) Leds() []Led {
	s.lock.Lock()
	leds := s.leds
	s.lock.Unlock()
	return leds
}
