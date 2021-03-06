package internal

import (
	"time"

	"github.com/sungup/img-resize/pkg/resizer"
)

type Options struct {
	DestDir string
	Format  resizer.ImgFormat

	Width   uint // new file width
	Height  uint // new file height
	Quality int  // convert format for JPEG

	ConvertInterval time.Duration
	KeepFileDate    bool
}
