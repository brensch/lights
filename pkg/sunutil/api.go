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

	q.Add("lat", latitude)
	q.Add("lng", longitude)
	q.Add("formatted", "0")
	q.Add("date", targetTime.Local().Format("2006-01-02"))
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
	light = targetTime.Before(sun.Sunset) && targetTime.After(sun.Sunrise)
	log.Printf("light: %t: target: %s, sunset: %s, sunrise: %s",
		light,
		targetTime.Local().Format(time.RFC3339),
		sun.Sunset.Local().Format(time.RFC3339),
		sun.Sunrise.Local().Format(time.RFC3339),
	)

	return

}
