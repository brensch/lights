package sunutil

import (
	"encoding/json"
	"fmt"
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

	q.Add("lat", "37.3996161838102")
	q.Add("lng", "-122.10145712094311")
	q.Add("formatted", "0")
	// fmt.Println("utc time", targetTime.Format("2006-01-02"))
	q.Add("date", targetTime.Format("2006-01-02"))
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

func TimeIsDark(targetTime time.Time) (dark bool, err error) {

	sun, err := GetSun(targetTime)
	if err != nil {
		return
	}

	dark = sun.AstronomicalTwilightEnd.Before(targetTime) || sun.AstronomicalTwilightBegin.After(targetTime)

	return

}
