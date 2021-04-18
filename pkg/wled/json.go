package wled

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(resBytes))
	}

	return
}
