package rand

import (
	"strconv"
	"testing"
)

func TestRandom(t *testing.T) {
	digit, err := Random(1, Ldigit)
	if err != nil {
		t.Error(err.Error())
		return
	}

	d, err := strconv.Atoi(digit)
	if err != nil {
		t.Error(err.Error())
		return
	} else if d > 10 || d < 0 {
		t.Error("invalid digit:", d)
	}
}

func TestUUID(t *testing.T) {
	uuid, err := UUID()
	if err != nil {
		t.Error(err.Error())
		return
	} else if len(uuid) != 36 {
		t.Error("invalid uuid:", uuid)
	}
}
