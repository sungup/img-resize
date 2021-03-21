package resizer

import "time"

type ConvertInfo struct {
	MDate time.Time

	SrcPath string
	DstPath string

	SrcFormat ImgFormat
	DstFormat ImgFormat

	Width   int
	Height  int
	Quality int
}
