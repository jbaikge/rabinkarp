package rabinkarp

import (
	"hash"
)

type (
	sum32 uint32
)

const (
	prime32 = 16777619
)

var (
	_ hash.Hash32 = new(sum32)
)

func New32() hash.Hash32 {
	return new(sum32)
}

func (s *sum32) BlockSize() int { return 1 }

func (s *sum32) Reset() { *s = 0 }

func (s *sum32) Size() int { return 4 }

func (s *sum32) Sum(in []byte) []byte {
	v := uint32(*s)
	return append(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (s *sum32) Sum32() uint32 { return uint32(*s) }

func (s *sum32) Write(data []byte) (int, error) {
	hash := *s
	for i := 0; i < len(data); i++ {
		hash = hash*prime32 + sum32(data[i])
	}
	*s = hash
	return len(data), nil
}
