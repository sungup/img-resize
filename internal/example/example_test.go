package example

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/nfnt/resize"
	"github.com/stretchr/testify/assert"
	"github.com/sungup/img-resize/test"
)

func Test_helloworld(t *testing.T) {
	tcInterpolation := map[string]resize.InterpolationFunction{
		"NearestNeighbor":   resize.NearestNeighbor,
		"Bininear":          resize.Bilinear,
		"Bicubic":           resize.Bicubic,
		"MitchellNetravali": resize.MitchellNetravali,
		"Lanczos2":          resize.Lanczos2,
		"Lanczos3":          resize.Lanczos3,
	}

	a := assert.New(t)

	srcPath, err := test.GetTestFilePath("images/source_image.png")
	a.NoError(err)
	srcStat, err := os.Stat(srcPath)
	a.NoError(err)

	srcByte, err := test.LoadTestFile("images/source_image.png")
	a.NoError(err)
	a.NotEmpty(srcByte)

	srcImage, err := png.Decode(bytes.NewReader(srcByte))
	a.NoError(err)

	for name, fn := range tcInterpolation {
		newImage := resize.Resize(2440, 0, srcImage, fn)

		dstPath, _ := test.GetTestFilePath(fmt.Sprintf("images/dest_image_%s.jpeg", name))
		out, err := os.Create(dstPath)
		a.NoError(err)
		defer out.Close()

		a.NoError(jpeg.Encode(out, newImage, &jpeg.Options{Quality: 100}))

		a.NoError(os.Chtimes(dstPath, srcStat.ModTime(), srcStat.ModTime()))
	}

}
