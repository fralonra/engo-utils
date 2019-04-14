package utils

import (
	"github.com/EngoEngine/engo/common"
	"image/color"
)

type Font struct {
	Font *common.Font
	URL  string
	Size float64
	BG   color.Color
	FG   color.Color
}

func (f *Font) Init() {
	f.Font = &common.Font{
		URL:  f.URL,
		FG:   f.FG,
		Size: f.Size,
	}

	err := f.Font.CreatePreloaded()
	if err != nil {
		panic(err)
	}
}

func (f *Font) SetFG(fg color.Color) {
	f.FG = fg
	f.Init()
}
