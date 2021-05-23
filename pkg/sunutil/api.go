package sunutil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetSun(targetTime time.Time) (results *APIResults, err error) {

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	q := req.URL.Query()

	zone, err := time.LoadLocation(timeZone)
	if err != nil {
		return
	}

	q.Add("lat", latitude)
	q.Add("lng", longitude)
	q.Add("formatted", "0")
	q.Add("date", targetTime.In(zone).Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response APIResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.Status != "OK" {
		err = fmt.Errorf(res.Status)
		return
	}

	results = response.Results

	return
}

func TimeIsLight(targetTime time.Time) (light bool, err error) {

	sun, err := GetSun(targetTime)
	if err != nil {
		return
	}

	// start and end of twilight at when it starts and ends being light
	light = targetTime.After(sun.CivilTwilightBegin) && targetTime.Before(sun.CivilTwilightEnd)
	log.Printf("light: %t: target: %s, twilight start: %s, twilight end: %s",
		light,
		targetTime.Local().Format(time.RFC3339),
		sun.CivilTwilightBegin.Local().Format(time.RFC3339),
		sun.CivilTwilightEnd.Local().Format(time.RFC3339),
	)

	return

}
