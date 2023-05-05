package internal

import "github.com/italorfeitosa/go-hands-on-prometheus/pkg/base62"

type Codec interface {
	Encode(int64) string
	Decode(string) int64
}

type Base62Codec struct{}

func (Base62Codec) Encode(id int64) string {
	return base62.Encode(id)
}

func (Base62Codec) Decode(slug string) int64 {
	return base62.Decode(slug)
}
