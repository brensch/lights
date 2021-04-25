package wled

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const (
	JSONEndpoint = "/json/state"
)

func (s *Server) Power(on bool) (err error) {

	stateReq := APIState{
		On: &on,
	}

	return s.SubmitState(stateReq)

}

func (s *Server) SubmitState(state APIState) (err error) {

	reqBytes, err := json.Marshal(state)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ip := range s.ips {

		target := fmt.Sprintf("http://%s%s", ip, JSONEndpoint)
		res, err := s.httpClient.Post(target, "application/json", bytes.NewBuffer(reqBytes))
		if err != nil {
			return err
		}

		var resBody APIResponse
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			return err
		}

		if resBody.Success == nil || !*resBody.Success {
			return fmt.Errorf("got success = false")
		}

	}

	return
}

func (s *Server) RandomEffect() (err error) {

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	state := APIState{
		Segments: []APISegment{
			{
				EffectId:        r.Intn(118),
				ColourPaletteID: r.Intn(56),
			},
		},
	}

	err = s.SubmitState(state)
	if err != nil {
		return
	}

	return
}
