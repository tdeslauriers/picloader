package dao

import (
	"testing"

	"github.com/picloader/model"
)

// must create album record in db
func TestCreateAlbum(t *testing.T) {

	test := model.Album{Album: "2018"}
	createAlbum(test)
}

// must find album by name/value

// obtain function
// find by name if exists, return || !exists, create
