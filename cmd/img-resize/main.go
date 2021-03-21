package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"time"

	"github.com/sungup/img-resize/internal"
	"github.com/sungup/img-resize/pkg/resizer"
)

const (
	DefaultDestDir     = "converted"
	DefaultDestFormat  = resizer.JPG
	DefaultDestWidth   = 2560
	DefaultDestHeight  = 1440
	DefaultDestQuality = 100

	DefaultConvertInterval = time.Second
	DefaultKeepFileDate    = true

	DefaultDirPermission = 0755
)

func argParse() (internal.Options, []string, error) {
	var (
		format string
		opts   internal.Options
		err    error
	)

	flag.StringVar(&opts.DestDir, "dest", DefaultDestDir, "destination path")
	flag.StringVar(&format, "format", DefaultDestFormat.String(), "converted file format [jpg|png]")
	flag.UintVar(&opts.Width, "width", DefaultDestWidth, "reduced width size")
	flag.UintVar(&opts.Height, "height", DefaultDestHeight, "reduced height size")
	flag.IntVar(&opts.Quality, "quality", DefaultDestQuality, "converted image quality")
	flag.DurationVar(&opts.ConvertInterval, "interval", DefaultConvertInterval, "converting interval")
	flag.BoolVar(&opts.KeepFileDate, "keep-filedate", DefaultKeepFileDate, "keep original file atime and mtime")

	flag.Parse()

	opts.Format, err = resizer.ToImgForm(format)

	return opts, flag.Args(), err
}

func main() {
	// 1. argument parsing
	opts, args, err := argParse()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// 2. make file list
	convertList := make([]*resizer.ConvertInfo, 0)
	for _, arg := range args {
		stat, err := os.Stat(arg)
		if os.IsNotExist(err) {
			log.Printf("input dir/file is not exists: %s\n", arg)
			continue
		}

		if stat.IsDir() {
			convertList = append(convertList, internal.MakeConvertInfoInDir(arg, opts)...)
		} else if convertInfo, err := internal.MakeConvertInfo(arg, opts); err == nil {
			convertList = append(convertList, convertInfo)
		} else {
			log.Printf("%s is not supported file: %v", arg, err)
		}
	}

	// 3. sort files in time-inorder
	sort.Sort(internal.ByMDate(convertList))

	// 4. create destination directory if possible
	if err := os.MkdirAll(opts.DestDir, DefaultDirPermission); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// 4. convert files into dest
	sleepTerm := time.Duration(0)

	for _, item := range convertList {
		time.Sleep(sleepTerm)

		if err := item.Convert(); err != nil {
			log.Printf("Convert failed: %v\n", err)
		}

		sleepTerm = opts.ConvertInterval
	}
}
