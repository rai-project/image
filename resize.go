package image

import (
	"image"
	"image/draw"

	"github.com/pkg/errors"
	"github.com/rai-project/image/asm"
	"github.com/rai-project/image/types"
)

func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

func doResize(targetPixels []uint8, srcPixels []uint8, targetWidth, targetHeight, srcWidth, srcHeight int) error {
	return asm.IResizeBilinear(targetPixels, srcPixels, targetWidth, targetHeight, srcWidth, srcHeight)
}

func Resize(inputImage types.Image, width, height int) (types.Image, error) {

	switch inputImage.(type) {
	case *types.RGBImage:
		out := types.NewRGBImage(image.Rect(0, 0, width, height))
		inPix := inputImage.(*types.RGBImage).Pix
		doResize(out.Pix, inPix, width, height, inputImage.Bounds().Dx(), inputImage.Bounds().Dy())
		return out, nil
	case *types.BGRImage:
		out := types.NewBGRImage(image.Rect(0, 0, width, height))
		inPix := inputImage.(*types.BGRImage).Pix
		doResize(out.Pix, inPix, width, height, inputImage.Bounds().Dx(), inputImage.Bounds().Dy())
		return out, nil
	default:
		return nil, errors.New("input image was neither an RGB nor a BGR image")
	}

}
