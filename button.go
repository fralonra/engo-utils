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

	b.label = &Label{
		Font:  b.Font,
		IsHud: b.IsHud,
		Text:  b.Text,
		World: b.World,
	}
	if b.texture != nil {
		width, height, _ := b.label.TextDimensions()
		b.label.Position = engo.Point{
			X: b.texture.SpaceComponent.Position.X + (b.texture.SpaceComponent.Width-float32(width))/2,
			Y: b.texture.SpaceComponent.Position.Y + (b.texture.SpaceComponent.Height-float32(height))/2,
		}
	} else {
		b.label.Position = b.Position
	}
	b.label.Init()

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

func (b *Button) Dimensions() (float32, float32) {
	return b.Clickable.SpaceComponent.Width, b.Clickable.SpaceComponent.Height
}
