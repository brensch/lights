package wled

import "testing"

func TestGetState(t *testing.T) {

	s, err := InitServer([]string{
		"192.168.1.2",
	}, 600)
	if err != nil {
		t.Error("couldn't init server", err)
		return
	}

	states, err := s.GetStates()
	if err != nil {
		t.Error("couldn't get state", err)
		return
	}

	t.Log(states[0].Segments[0].EffectId)
}

func TestRandomEffect(t *testing.T) {

	s, err := InitServer([]string{
		"192.168.1.2",
	}, 600)
	if err != nil {
		t.Error("couldn't init server", err)
		return
	}

	err = s.RandomEffect()
	if err != nil {
		t.Error("couldn't get state", err)
		return
	}

}
