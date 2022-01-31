package id

import (
	"strings"

	shortid "github.com/jasonsoft/go-short-id"
)

func GenShortId() string {
	opt := shortid.Options{
		Number:        6,
		StartWithYear: true,
		EndWithHost:   false,
	}
	return strings.ToLower(shortid.Generate(opt))
}
