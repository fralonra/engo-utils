package utils

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
)

const (
	radioBoxSize   = 20
	radioInnerSize = 12

	eventChange = iota
)

type radioBoxInner struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type radioBox struct {
	isChecked bool
	isHud     bool
	position  engo.Point
	world     *ecs.World
	inner     *radioBoxInner

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (r *radioBox) Init() {
	r.BasicEntity = ecs.NewBasic()

	r.RenderComponent = common.RenderComponent{
		Drawable: common.Circle{
			BorderWidth: 2,
			BorderColor: color.Black,
		},
		Color: color.White,
	}
	r.SpaceComponent = common.SpaceComponent{
		Width:    radioBoxSize,
		Height:   radioBoxSize,
		Position: r.position,
	}

	r.inner = &radioBoxInner{}
	r.inner.BasicEntity = ecs.NewBasic()

	r.inner.RenderComponent = common.RenderComponent{
		Drawable: common.Circle{},
		Color:    color.Black,
	}
	radioInnerOffset := float32((radioBoxSize - radioInnerSize) / 2)
	r.inner.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: r.position.X + radioInnerOffset,
			Y: r.position.Y + radioInnerOffset,
		},
	}
	if r.isChecked {
		r.inner.SpaceComponent.Width = radioInnerSize
		r.inner.SpaceComponent.Height = radioInnerSize
	}

	if r.isHud {
		r.SetShader(common.HUDShader)
		r.inner.SetShader(common.HUDShader)
	}

	for _, system := range r.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&r.BasicEntity, &r.RenderComponent, &r.SpaceComponent)
			sys.Add(&r.inner.BasicEntity, &r.inner.RenderComponent, &r.inner.SpaceComponent)
		}
	}
}

func (r *radioBox) toggle(isChecked bool) {
	r.isChecked = isChecked
	if r.isChecked {
		r.inner.SpaceComponent.Width = radioInnerSize
		r.inner.SpaceComponent.Height = radioInnerSize
	} else {
		r.inner.SpaceComponent.Width = 0
		r.inner.SpaceComponent.Height = 0
	}
}

type Radio struct {
	Clickable
	Label
	index     int
	isChecked bool
	isHud     bool
	Position  engo.Point
	World     *ecs.World

	radioBox
	group *RadioGroup
}

func (r *Radio) Init() {
	r.radioBox.world = r.World
	r.radioBox.position = engo.Point{
		X: r.Position.X + 10,
		Y: r.Position.Y + 10,
	}
	r.radioBox.isChecked = r.isChecked
	r.radioBox.isHud = r.isHud
	r.radioBox.Init()

	r.Label.World = r.World
	r.Label.Position = engo.Point{
		X: r.Position.X + 50,
		Y: r.Position.Y + 10,
	}
	r.Label.IsHud = r.isHud
	r.Label.Init()

	r.Clickable.World = r.World
	r.Clickable.Position = r.Position
	r.Clickable.IsHud = r.isHud
	r.Clickable.SpaceComponent = r.Label.SpaceComponent
	r.Clickable.Init()
	r.Clickable.OnClick(func() {
		if r.isChecked {
			return
		}
		r.group.handleChange(r.index)
	})
}

func (r *Radio) toggle(isChecked bool) {
	r.isChecked = isChecked
	r.radioBox.toggle(isChecked)
}

type RadioGroup struct {
	activeIndex   int
	eventHandlers map[int]func()
	hasChanged    bool
	RadioList     []*Radio
	World         *ecs.World

	ecs.BasicEntity
}

func (r *RadioGroup) Init() {
	r.BasicEntity = ecs.NewBasic()

	r.eventHandlers = make(map[int]func())

	for i, radio := range r.RadioList {
		if r.activeIndex >= 0 && r.activeIndex == i {
			radio.isChecked = true
		}
		radio.index = i
		radio.group = r
		radio.Init()
	}

	for _, system := range r.World.Systems() {
		switch sys := system.(type) {
		case *RadioSystem:
			sys.Add(r)
		}
	}
}

func (r *RadioGroup) On(event int, f func()) {
	r.eventHandlers[event] = f
}

func (r *RadioGroup) OnChange(f func()) {
	r.On(eventChange, f)
}

func (r *RadioGroup) GetActiveIndex() int {
	return r.activeIndex
}

func (r *RadioGroup) handleChange(index int) {
	r.activeIndex = index
	r.hasChanged = true
}

type RadioSystem struct {
	entities []*RadioGroup
}

func (s *RadioSystem) Add(r *RadioGroup) {
	s.entities = append(s.entities, r)
}

func (s *RadioSystem) Remove(basic ecs.BasicEntity) {
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

func (s *RadioSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.hasChanged {
			for i, r := range e.RadioList {
				if e.activeIndex == i {
					r.toggle(true)
				} else {
					r.toggle(false)
				}
			}
			if f, ok := e.eventHandlers[eventChange]; ok {
				f()
			}
			e.hasChanged = false
		}
	}
}
