package resizer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/nfnt/resize"
)

const (
	interpolationFn = resize.Lanczos3
)

type ConvertInfo struct {
	MDate time.Time

	SrcPath string
	DstPath string

	SrcFormat ImgFormat
	DstFormat ImgFormat

	Width   uint
	Height  uint
	Quality int
}

func (c *ConvertInfo) decodeFromSrc() (image.Image, error) {
	fp, _ := os.Open(c.SrcPath)
	defer func() { _ = fp.Close() }()

	switch c.SrcFormat {
	case PNG:
		return png.Decode(fp)
	case JPG:
		return jpeg.Decode(fp)
	}

	return nil, fmt.Errorf("unsupported source format: %s", c.SrcFormat.String())
}

func (c *ConvertInfo) encodeToDst(img image.Image) error {
	var (
		fp  *os.File
		err error
	)

	if fp, err = os.Create(c.DstPath); err != nil {
		return err
	}
	defer func() { _ = fp.Close() }()

	switch c.DstFormat {
	case PNG:
		err = png.Encode(fp, img)
	case JPG:
		err = jpeg.Encode(fp, img, &jpeg.Options{Quality: c.Quality})
	default:
		err = fmt.Errorf("unsupported dest format: %s", c.DstFormat.String())
	}

	os.Chtimes(c.DstPath, c.MDate, c.MDate)

	return err
}

func (c *ConvertInfo) Convert() error {
	if img, err := c.decodeFromSrc(); err != nil {
		return err
	} else {
		return c.encodeToDst(resize.Resize(c.Width, c.Height, img, interpolationFn))
	}
}
