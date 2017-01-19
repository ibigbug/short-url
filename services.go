package main

import (
	"math"
	"sync"
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

type SimpleStorage struct {
	sync.RWMutex
	shortToLong map[string]string
	longToShort map[string]string
}

func (s *SimpleStorage) GetByOrigin(origin string) string {
	s.RLock()
	defer s.RUnlock()
	return s.longToShort[origin]
}

func (s *SimpleStorage) GetByShort(short string) string {
	s.RLock()
	defer s.RUnlock()
	return s.shortToLong[short]
}

func (s *SimpleStorage) Save(short, origin string) {
	s.Lock()
	defer s.Unlock()
	s.shortToLong[short] = origin
	s.longToShort[origin] = short
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
	// TODO: may overflow
	r := make([]byte, 0, 6)
	for i != 0 {
		mod := int(math.Mod(float64(i), 62.0))
		r = append(r, CHAR_MAPPING[mod])
		i = i / 62
	}
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func NewShortService(idGen IdGenerator, storage Storage) *SimpleShorter {
	if idGen == nil {
		idGen = &SimpleIdGen{
			start: 0,
		}
	}
	if storage == nil {
		storage = &SimpleStorage{
			shortToLong: map[string]string{},
			longToShort: map[string]string{},
		}
	}
	return &SimpleShorter{
		idGen:   idGen,
		storage: storage,
	}
}
