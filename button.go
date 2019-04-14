package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type Button struct {
	Clickable
	Label
	IsHud    bool
	Position engo.Point
	World    *ecs.World
}

func (b *Button) Init() {
	b.Label.World = b.World
	b.Label.Position = b.Position
	if b.IsHud {
		b.Label.IsHud = true
	}
	b.Label.Init()

	b.Clickable.World = b.World
	b.Clickable.Position = b.Position
	if b.IsHud {
		b.Clickable.IsHud = true
	}
	b.Clickable.SpaceComponent = b.Label.SpaceComponent
	b.Clickable.Init()
}
