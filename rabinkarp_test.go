package rabinkarp

import (
	"testing"
)

func Test32(t *testing.T) {
	tests := map[string]uint32{
		"0":           48,
		"5":           53,
		"a":           97,
		"b":           98,
		"c":           99,
		"abc":         3371236466,
		"lorem ipsum": 932161751,
		"Lorem Ipsum": 554694039,
	}
	h := New32()
	for str, exp := range tests {
		h.Write([]byte(str))
		if sum := h.Sum32(); sum != exp {
			t.Errorf("[%s] Expected %d; Got %d", str, exp, sum)
		}
		h.Reset()
	}
}
