package utils

import (
	"log"

	"github.com/EngoEngine/engo"
)

type Scene struct {
	Name string
}

func (s *Scene) Type() string {
	return s.Name
}

func (*Scene) Preload() {}

func (*Scene) Setup(engo.Updater) {}

func (s *Scene) Hide() {
	log.Println(s.Name + " is now hidden")
}

func (s *Scene) Show() {
	log.Println(s.Name + " is now shown")
}
