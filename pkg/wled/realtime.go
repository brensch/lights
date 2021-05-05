package wled

import (
	"encoding/binary"
)

// details are outlined at https://github.com/Aircoookie/WLED/wiki/UDP-Realtime-Control

const (
	RealtimeTimeoutSeconds = 2

	ModeDNRGB = 4

	MaxLedsPerPacket = 489

	BytePositionRealtimeMode   = 0
	BytePositionTimeout        = 1
	BytePositionStartIndexHigh = 2
	BytePositionStartIndexLow  = 3
	BytePositionRedOffset      = 4
	BytePositionGreenOffset    = 5
	BytePositionBlueOffset     = 6
)

func ConstructPackets(leds []*Led) (packets [][]byte) {

	for i := 0; i < len(leds); i += MaxLedsPerPacket {

		endOfSegment := i + MaxLedsPerPacket

		if endOfSegment > len(leds) {
			endOfSegment = len(leds) - 1
		}

		packet := make([]byte, 4+3*(endOfSegment-i))

		packet[BytePositionRealtimeMode] = ModeDNRGB
		packet[BytePositionTimeout] = RealtimeTimeoutSeconds

		start := uint16(i)
		binary.BigEndian.PutUint16(packet[BytePositionStartIndexHigh:BytePositionStartIndexLow+1], start)
		// fmt.Println("yeeeert: ", far[0], far[1])

		for j := 0; j < endOfSegment-i; j++ {
			leds[i+j].lock.Lock()
			packet[BytePositionRedOffset+j*3] = leds[i+j].state.Red
			packet[BytePositionGreenOffset+j*3] = leds[i+j].state.Green
			packet[BytePositionBlueOffset+j*3] = leds[i+j].state.Blue
			leds[i+j].lock.Unlock()
		}

		packets = append(packets, packet)
	}

	return
}

// func generatePackets
