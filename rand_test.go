package rand

import (
	"strconv"
	"testing"
)

func TestRandom(t *testing.T) {
	digits, err := Random(6, Ldigit)
	if err != nil {
		t.Error(err.Error())
		return
	} else if len(digits) != 6 {
		t.Error("invalid digit:", digits)
		return
	}

	for _, b := range digits {
		d, err := strconv.Atoi(string(b))
		if err != nil {
			t.Error(err.Error())
			return
		} else if d > 10 || d < 0 {
			t.Error("invalid digit:", d)
		}
	}
}

func TestShortStr(t *testing.T) {
	g, err := ShortStr([]byte(MustUUID()), 6, Ldigit)
	if err != nil {
		t.Error(err.Error())
		return
	} else if len(g) != 4 {
		t.Error("invalid short group:", g)
		return
	}

	for _, item := range g {
		for _, b := range item {
			d, err := strconv.Atoi(string(b))
			if err != nil {
				t.Error(err.Error())
				return
			} else if d > 10 || d < 0 {
				t.Error("invalid digit:", d)
			}
		}
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
