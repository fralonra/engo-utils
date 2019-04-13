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
	Position *engo.Point
	World    *ecs.World

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (t *Texture) Init() {
	t.BasicEntity = ecs.NewBasic()

	texture, err := common.LoadedSprite(t.Url)
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}
	t.Width = texture.Width()
	t.Height = texture.Height()

	if t.IsHud {
		t.SetShader(common.HUDShader)
	}

	t.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale: engo.Point{
			X: 1,
			Y: 1,
		},
	}

	log.Printf("tex: pos: %#v", t.RenderComponent.Drawable)

	t.SpaceComponent = common.SpaceComponent{
		Width:  texture.Width(),
		Height: texture.Height(),
	}
	if t.Position != nil {
		t.SpaceComponent.Position = *t.Position
	}

	for _, system := range t.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
		}
	}
}

func (t *Texture) Translate(x float32, y float32) {
	t.Position = &engo.Point{
		X: x,
		Y: y,
	}
	t.SpaceComponent.Position = *t.Position
}
