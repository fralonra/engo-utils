package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Label struct {
	Font     *common.Font
	IsHud    bool
	Position engo.Point
	Text     string
	World    *ecs.World

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (l *Label) Init() {
	l.BasicEntity = ecs.NewBasic()

	width, height, _ := l.TextDimensions()
	l.RenderComponent.Drawable = common.Text{
		Font: l.Font,
		Text: l.Text,
	}
	l.SpaceComponent = common.SpaceComponent{
		Width:    float32(width),
		Height:   float32(height),
		Position: l.Position,
	}

	if l.IsHud {
		l.SetShader(common.TextHUDShader)
	}

	for _, system := range l.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&l.BasicEntity, &l.RenderComponent, &l.SpaceComponent)
		}
	}
}

func (l *Label) TextDimensions() (int, int, int) {
	return l.Font.TextDimensions(l.Text)
}

func (l *Label) SetFont(font *common.Font) {
	l.Font = font
	l.setDrawble()
}

func (l *Label) SetText(text string) {
	l.Text = text
	l.setDrawble()
}

func (l *Label) setDrawble() {
	l.RenderComponent.Drawable = common.Text{
		Font: l.Font,
		Text: l.Text,
	}
}
