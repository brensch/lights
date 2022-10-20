package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brensch/lights/pkg/sunutil"
	"github.com/brensch/lights/pkg/wled"
)

func doRandomEffect(s *wled.Server) {
	isLight, err := sunutil.TimeIsLight(time.Now())
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("hour", time.Now().Local().Hour())

	// turn off at midnight
	if time.Now().Local().Hour() < 12 {
		isLight = true
	}

	// turn off if it is light
	if isLight {
		err = s.Power(false)
		if err != nil {
			log.Println(err)
		}
		return
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
