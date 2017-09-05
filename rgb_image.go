package image

import (
	"image"
	"image/color"

	context "golang.org/x/net/context"

	"github.com/pkg/errors"
)

// RGBImage is an in-memory image whose At method returns RGB values.
type RGBImage struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []float32
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p RGBImage) ColorModel() color.Model { return RGBModel }

func (p RGBImage) Bounds() image.Rectangle { return p.Rect }

func (p RGBImage) Mode() mode { return RGBMode }

func (p RGBImage) At(x, y int) color.Color {
	return p.RGBAt(x, y)
}

func (p *RGBImage) RGBAt(x, y int) RGB {
	if !(image.Point{x, y}.In(p.Rect)) {
		return RGB{}
	}
	i := p.PixOffset(x, y)
	return RGB{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGBImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGBModel.Convert(c).(RGB)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
}

func (p *RGBImage) SetRGB(x, y int, c RGB) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBImage) SubImage(r image.Rectangle) Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

func (p *RGBImage) fillFromRGBAImage(ctx context.Context, rgbaImage *image.RGBA) error {
	if p.Bounds() != rgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), rgbaImage.Bounds())
	}

	width := rgbaImage.Bounds().Dx()
	height := rgbaImage.Bounds().Dy()
	stride := rgbaImage.Stride

	rgbaImagePixels := rgbaImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		rgbaOffset := y * stride
		rgbOffset := y * width
		for x := 0; x < width; x++ {
			rgbImagePixels[rgbOffset+0] = float32(rgbaImagePixels[rgbaOffset+0])
			rgbImagePixels[rgbOffset+1] = float32(rgbaImagePixels[rgbaOffset+1])
			rgbImagePixels[rgbOffset+2] = float32(rgbaImagePixels[rgbaOffset+2])
			rgbaOffset += 4
			rgbOffset += 3
		}
	}

	return nil
}

func (p *RGBImage) fillFromNRGBAImage(ctx context.Context, nrgbaImage *image.NRGBA) error {
	if p.Bounds() != nrgbaImage.Bounds() {
		return errors.Errorf("the bounds %v and %v did not match", p.Bounds(), nrgbaImage.Bounds())
	}

	width := nrgbaImage.Bounds().Dx()
	height := nrgbaImage.Bounds().Dy()
	stride := nrgbaImage.Stride

	nrgbaImagePixels := nrgbaImage.Pix
	rgbImagePixels := p.Pix
	for y := 0; y < height; y++ {
		nrgbaOffset := y * stride
		rgbOffset := y * width
		for x := 0; x < width; x++ {
			rgbImagePixels[rgbOffset+0] = float32(nrgbaImagePixels[nrgbaOffset+0])
			rgbImagePixels[rgbOffset+1] = float32(nrgbaImagePixels[nrgbaOffset+1])
			rgbImagePixels[rgbOffset+2] = float32(nrgbaImagePixels[nrgbaOffset+2])
			nrgbaOffset += 4
			rgbOffset += 3
		}
	}

	return nil
}

// NewRGBImage returns a new RGBImage image with the given bounds.
func NewRGBImage(r image.Rectangle) *RGBImage {
	w, h := r.Dx(), r.Dy()
	buf := make([]float32, 3*w*h)
	return &RGBImage{buf, 3 * w, r}
}
