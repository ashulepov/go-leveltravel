package leveltravel

import "testing"

func TestNewAPI(t *testing.T) {
	if lt := NewAPI("key", ""); lt == nil {
		t.Logf("init LevelTravel API error")
	}
}
