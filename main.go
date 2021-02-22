package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/picloader/model"
	"github.com/picloader/util"
	"github.com/rwcarlsen/goexif/exif"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Atomic dog...")

	dir := "/home/tombomb/Pictures/test/source/"
	thumbdir := "/home/tombomb/Pictures/test/thumbs/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// need to test for nil
	imgs := make(model.Pics, 0, 100)
	for _, f := range files {

		p, err := os.Open(dir + f.Name())
		if err != nil {
			log.Panicf("Could not open file: %s.\n%v", f.Name(), err)
		}

		x, err := exif.Decode(p)
		if err != nil {
			log.Panicf("Could not decode image: %s.\n%v", f.Name(), err)
		}

		tm, _ := x.DateTime()

		pic := model.Pic{Filename: uuid.New(), Date: tm}
		imgs = append(imgs, pic)

		tmb, _ := x.JpegThumbnail()
		thumb := thumbdir + pic.Filename.String() + "_thumb.jpg"
		util.MakeThumb(tmb, thumb)

		err = os.Rename(dir+f.Name(), dir+pic.Filename.String()+".jpg")
		if err != nil {
			panic(err)
		}
	}

}
