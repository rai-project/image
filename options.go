package image

import (
	"github.com/rai-project/image/types"
	context "golang.org/x/net/context"
)

type Options struct {
	resizeWidth  int
	resizeHeight int
	mode         types.Mode
	mean         [3]float32
	layout       Layout
	ctx          context.Context
}

type Option func(o *Options)

func Resized(width, height int) Option {
	return func(o *Options) {
		o.resizeWidth = width
		o.resizeHeight = height
	}
}

func Mean(mean [3]float32) Option {
	return func(o *Options) {
		o.mean = mean
	}
}

func MeanValue(mean float32) Option {
	return Mean([3]float32{mean, mean, mean})
}

func ChannelLayout(layout Layout) Option {
	return func(o *Options) {
		o.layout = layout
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Mode(mode types.Mode) Option {
	return func(o *Options) {
		o.mode = mode
	}
}

func NewOptions() *Options {
	return &Options{
		mean:   [3]float32{0, 0, 0},
		mode:   types.RGBMode,
		layout: HWCLayout,
	}
}
