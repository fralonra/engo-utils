package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type MenuEntity struct {
	Text    string
	Texture string
	Gap     float32
	OnClick func()
}

type Menu struct {
	Font     *common.Font
	Texture  string
	Gap      float32
	IsHud    bool
	Position engo.Point
	World    *ecs.World

	Entities []MenuEntity
}

func (m *Menu) Init() {
	lastY := m.Position.Y
	for _, entity := range m.Entities {
		button := Button{
			World: m.World,
			Font:  m.Font,
			Text:  entity.Text,
			IsHud: m.IsHud,
			Position: engo.Point{
				X: m.Position.X,
			},
		}
		if entity.Texture != "" {
			button.Texture = entity.Texture
		} else {
			button.Texture = m.Texture
		}
		if entity.Gap != 0 {
			button.Position.Y = lastY + entity.Gap
			lastY += entity.Gap
		} else {
			button.Position.Y = lastY + m.Gap
			lastY += m.Gap
		}
		button.Init()
		button.OnClick(entity.OnClick)
		_, buttonHeight := button.Dimensions()
		lastY += buttonHeight
	}
}
