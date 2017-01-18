package main

type ShortenParam struct {
	Url string `json:"url"`
}

type OriginalParam struct {
	Short string `json:"short"`
}

type Storage interface {
	GetByShort(s string) string
	GetByOrigin(r string) string
	Save(short, origin string)
}

type IdGenerator interface {
	Gen() string
}
