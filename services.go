package main

import (
	"fmt"
	"math"
	"sync/atomic"
)

const (
	CHAR_MAPPING = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type SimpleShorter struct {
	idGen   IdGenerator
	storage Storage
}

func (s *SimpleShorter) Short(original string) (short string) {
	short = s.storage.GetByOrigin(original)
	if short != "" {
		return
	}
	short = s.idGen.Gen()
	s.storage.Save(short, original)
	return
}

func (s *SimpleShorter) Original(short string) (original string) {
	original = s.storage.GetByShort(short)
	return
}

type SimpleIdGen struct {
	start int64
}

func (s *SimpleIdGen) Gen() string {
	return s.base62(atomic.AddInt64(&(s.start), 1))
}

func (s *SimpleIdGen) base62(i int64) string {
	if i == 0 {
		return "0"
	}
	r := make([]byte, 0, 6)
	for i != 0 {
		fmt.Println("i", i)
		mod := int(math.Mod(float64(i), 62.0))
		fmt.Println("mod", mod)
		r = append(r, CHAR_MAPPING[mod])
		i = i / 62
	}
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func NewShortService(idGen IdGenerator, storage Storage) *SimpleShorter {
	return &SimpleShorter{
		idGen:   idGen,
		storage: storage,
	}
}
