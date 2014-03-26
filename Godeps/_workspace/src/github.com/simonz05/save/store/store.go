package store

import (
	"github.com/simonz05/save"
)

type foo struct {
}

func (s *foo) Save(v string) error {
	return nil
}

var _ save.Saver = &foo{}
