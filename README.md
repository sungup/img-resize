# Img-Resize

CLI based image resizing tool

## Install

```bash
$ go get github.com/sungup/img-resize/cmd/img-resize
```

And you can run `img-resize` on your terminal.

## Usage

```txt
Usage of img-resize:
  -dest string
    	destination path (default "converted")
  -format string
    	converted file format [jpg|png] (default "jpg")
  -height uint
    	reduced height size (default 1440)
  -interval duration
    	converting interval (default 1s)
  -keep-filedate
    	keep original file atime and mtime (default true)
  -quality int
    	converted image quality (default 100)
  -width uint
    	reduced width size (default 2560)
```

## License

Copyright (c) 2021 Sungu Moon <sungup@me.com>.
img-resize under a MIT style license.

## Reference

- [nfnt/resize](https://github.com/nfnt/resize)
