package main

import (
	"testing"

	"github.com/brensch/lights/pkg/wled"
)

func TestDoRandomEffect(t *testing.T) {

	s, err := wled.InitServer([]string{
		"192.168.1.2",
	}, 600)
	if err != nil {
		t.Error("couldn't init server", err)
		return
	}

	doRandomEffect(s)

}
