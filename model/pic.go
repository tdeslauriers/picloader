package model

import (
	"time"

	"github.com/google/uuid"
)

// Pic used for idb population + file rename
type Pic struct {
	ID       int64
	Filename uuid.UUID
	Date     time.Time
	AlbumID  int64
}

type Pics []Pic

// Album used for populating db
type Album struct {
	ID    int64
	Album string
}
