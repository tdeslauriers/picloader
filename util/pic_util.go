package util

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

func MakeThumb(thumbRaw []byte, name string) {

	thumb, _, err := image.Decode(bytes.NewReader(thumbRaw))
	if err != nil {
		panic(err)
	}

	out, _ := os.Create(name)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 30

	err = jpeg.Encode(out, thumb, &opts)

	out.Close()
}
