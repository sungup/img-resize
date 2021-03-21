package resizer

import (
	"fmt"
	"path"
	"strings"
)

type ImgFormat int

const (
	PNG ImgFormat = iota
	JPG

	UnsupportedImg = ImgFormat(0xfe)
)

var (
	imgFormToStr = map[ImgFormat]string{
		PNG: "png",
		JPG: "jpg",
	}

	imgExtensions = map[ImgFormat][]string{
		PNG: {"png"},
		JPG: {"jpg", "jpeg"},
	}

	strToImgForm = map[string]ImgFormat{
		"png": PNG,
		"jpg": JPG,
	}
)

func (i ImgFormat) String() string {
	return imgFormToStr[i]
}

func (i ImgFormat) Ext() string {
	return i.String()
}

func ToImgForm(format string) (ImgFormat, error) {
	if imgForm, ok := strToImgForm[format]; ok {
		return imgForm, nil
	} else {
		return UnsupportedImg, fmt.Errorf("unsupported type: %s", format)
	}
}

func SupportImgFormat(imgPath string) (ImgFormat, error) {
	imgExt := strings.ToLower(path.Ext(imgPath))[1:]

	for format, extensions := range imgExtensions {
		for _, ext := range extensions {
			if ext == imgExt {
				return format, nil
			}
		}
	}

	return UnsupportedImg, fmt.Errorf("unsupported type: %s", imgExt)
}
