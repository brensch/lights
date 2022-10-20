package sunutil

import (
	"testing"
	"time"
)

func TestTimeIsLight(t *testing.T) {

	light, err := TimeIsLight(time.Now())
	if err != nil {
		t.Error("error getting light state", err)
		return
	}

	t.Log(light)

}
