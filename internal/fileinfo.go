package internal

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sungup/img-resize/pkg/resizer"
)

type ByMDate []*resizer.ConvertInfo

func (l ByMDate) Len() int           { return len(l) }
func (l ByMDate) Less(i, j int) bool { return l[i].MDate.Before(l[j].MDate) }
func (l ByMDate) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func MakeConvertInfo(filepath string, opts Options) (*resizer.ConvertInfo, error) {
	stat, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return nil, err
	}

	format, err := resizer.SupportImgFormat(filepath)
	if err != nil {
		return nil, err
	}

	mdate := time.Now()
	if opts.KeepFileDate {
		mdate = stat.ModTime()
	}

	return &resizer.ConvertInfo{
		MDate:     mdate,
		SrcPath:   filepath,
		DstPath:   path.Join(opts.DestDir, stat.Name()[:len(stat.Name())-3]+opts.Format.Ext()),
		SrcFormat: format,
		DstFormat: opts.Format,
		Width:     opts.Width,
		Height:    opts.Height,
		Quality:   opts.Quality,
	}, nil
}

func MakeConvertInfoInDir(dirpath string, opts Options) []*resizer.ConvertInfo {
	convertList := make([]*resizer.ConvertInfo, 0)

	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return convertList
	}

	for _, item := range entries {
		if item.IsDir() {
			continue
		}

		if convertInfo, err := MakeConvertInfo(path.Join(dirpath, item.Name()), opts); err == nil {
			convertList = append(convertList, convertInfo)
		} else {
			log.Println(err)
		}

	}

	return convertList
}
