package main

type ShortenParam struct {
	Short string `json:"short"`
}

type OriginalParam struct {
	Original string `json:"original"`
}

type Storage interface {
	GetByShort(s string) string
	GetByOrigin(r string) string
	Save(short, origin string)
}

type IdGenerator interface {
	Gen() string
}
