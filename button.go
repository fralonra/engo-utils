package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Button struct {
	Clickable
	Font     *common.Font
	Text     string
	Texture  string
	IsHud    bool
	Position engo.Point
	World    *ecs.World

	label   *Label
	texture *Texture
}

func (b *Button) Init() {
	b.label = &Label{
		Font:     b.Font,
		IsHud:    b.IsHud,
		Position: b.Position,
		Text:     b.Text,
		World:    b.World,
	}
	b.label.Init()

	if b.Texture != "" {
		b.texture = &Texture{
			World:    b.World,
			Url:      b.Texture,
			Position: b.Position,
		}
		if b.IsHud {
			b.texture.IsHud = true
		}
		b.texture.Init()
	}

	b.Clickable.World = b.World
	b.Clickable.Position = b.Position
	if b.IsHud {
		b.Clickable.IsHud = true
	}
	if b.texture != nil {
		b.Clickable.SpaceComponent = b.texture.SpaceComponent
	} else {
		b.Clickable.SpaceComponent = b.label.SpaceComponent
	}
	b.Clickable.Init()
}
