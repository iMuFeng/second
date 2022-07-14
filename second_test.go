package second

import (
	"testing"
)

var pairs = []struct {
	k string
	v float64
	e bool
}{
	{k: "five days", e: true},
	{k: "100", e: true},
	{k: "1m", v: 60},
	{k: "1m", v: 60},
	{k: "1h", v: 3600},
	{k: "2d", v: 172800},
	{k: "3w", v: 1814400},
	{k: "1s", v: 1},
	{k: "1y", v: 31557600},
	{k: "1.5h", v: 5400},
	{k: "1   s", v: 1},
	{k: "1H", v: 3600},
	{k: "0.5s", v: 0.5},
	{k: "-1s", v: -1},
	{k: "-1.5h", v: -5400},
	{k: "-10.5h", v: -37800},
	{k: "-.5h", v: -1800},
}

func TestParse(t *testing.T) {
	for _, r := range pairs {
		n, err := Parse(r.k)

		if err != nil {
			if !r.e {
				t.Error(err)
			}
		} else {
			if r.e {
				t.Errorf("Failed to find out not matched string: %s", r.k)
			} else if n != r.v {
				t.Errorf("%s's result %f is not equal to %f", r.k, r.v, n)
			}
		}
	}
}
