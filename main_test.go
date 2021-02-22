package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/picloader/model"
	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {

	t.Log("Atomic dog test...")

	i := 0
	for i < 10 {

		j := uuid.New()
		t.Log(j)
		i++
	}
}

//must load an array of images
func TestImageLoad(t *testing.T) {

	imgs := make(model.Pics, 10, 10)
	for j := 0; j < 10; j++ {
		u := uuid.New()
		d := time.Now()
		i := model.Pic{Filename: u, Date: d}
		imgs[j] = i
	}
	for _, j := range imgs {
		t.Logf("image {\n\tfilename: %v\n\tdate: %v\n}\n", j.Filename, j.Date.Format("2006-01-02"))
	}
	assert.True(t, len(imgs) == 10)
}

// needs to change jpg file name: done

// needs to assign exif data

// needs to create thumbnail if not

// needs to load db by year OR!! topic
