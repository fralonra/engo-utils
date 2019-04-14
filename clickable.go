package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	eventClick = iota
	eventMouseover
)

type Clickable struct {
	eventHandlers map[int]func()
	IsHud         bool
	Position      engo.Point
	World         *ecs.World

	ecs.BasicEntity
	common.MouseComponent
	common.RenderComponent
	common.SpaceComponent
}

func (c *Clickable) Init() {
	c.BasicEntity = ecs.NewBasic()

	c.eventHandlers = make(map[int]func())

	c.MouseComponent = common.MouseComponent{}

	if c.IsHud {
		c.SetShader(common.HUDShader)
	}

	for _, system := range c.World.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&c.BasicEntity, &c.MouseComponent, &c.SpaceComponent, &c.RenderComponent)
		case *ClickableSystem:
			sys.Add(c)
		}
	}
}

func (c *Clickable) On(event int, f func()) {
	c.eventHandlers[event] = f
}

func (c *Clickable) OnClick(f func()) {
	c.On(eventClick, f)
}

func (c *Clickable) OnMouseOver(f func()) {
	c.On(eventMouseover, f)
}

type ClickableSystem struct {
	entities []*Clickable
}

func (s *ClickableSystem) Add(c *Clickable) {
	s.entities = append(s.entities, c)
}

func (s *ClickableSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *ClickableSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.Clicked {
			if f, ok := e.eventHandlers[eventClick]; ok {
				f()
			}
		}
		if e.Enter {
			engo.SetCursor(engo.CursorHand)
		}
		if e.Leave {
			engo.SetCursor(engo.CursorArrow)
		}
	}
}
