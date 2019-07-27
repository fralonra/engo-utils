package utils

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Box struct {
	Color       color.Color
	Width       float32
	Height      float32
	BorderColor color.Color
	BorderWidth float32
	IsHud       bool
	Position    engo.Point
	World       *ecs.World

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (b *Box) Init() {
	b.BasicEntity = ecs.NewBasic()

	if b.Height == 0 {
		b.Height = b.Width
	}

	b.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{
			BorderWidth: b.BorderWidth,
			BorderColor: b.BorderColor,
		},
		Color: b.Color,
	}
	b.SpaceComponent = common.SpaceComponent{
		Width:    b.Width,
		Height:   b.Height,
		Position: b.Position,
	}

	if b.IsHud {
		b.SetShader(common.HUDShader)
	}

	for _, system := range b.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&b.BasicEntity, &b.RenderComponent, &b.SpaceComponent)
		}
	}
}
