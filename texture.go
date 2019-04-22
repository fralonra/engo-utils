package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

type Texture struct {
	Url      string
	IsHud    bool
	Position engo.Point
	Width    float32
	Height   float32
	World    *ecs.World

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (t *Texture) Init() {
	t.BasicEntity = ecs.NewBasic()

	texture, err := common.LoadedSprite(t.Url)
	if err != nil {
		log.Printf("Unable to load texture: %#v", err.Error())
	}
	if t.Width == 0 {
		t.Width = texture.Width()
	}
	if t.Height == 0 {
		t.Height = texture.Height()
	}

	t.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale: engo.Point{
			X: 1,
			Y: 1,
		},
	}

	t.SpaceComponent = common.SpaceComponent{
		Width:  texture.Width(),
		Height: texture.Height(),
		Position: t.Position,
	}

	if t.IsHud {
		t.SetShader(common.HUDShader)
	}

	for _, system := range t.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
		}
	}
}

func (t *Texture) Translate(x float32, y float32) {
	t.Position = engo.Point{
		X: x,
		Y: y,
	}
	t.SpaceComponent.Position = t.Position
}
