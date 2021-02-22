package dao

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/picloader/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateImage(t *testing.T) {

	test := model.Pic{Filename: uuid.New(), Date: time.Now()}
	CreateImage(test)
}

func TestFindImageById(t *testing.T) {

	test := model.Pic{Filename: uuid.New(), Date: time.Now()}
	id, err := CreateImage(test)
	if err != nil {
		t.Log(err)
	}
	f := FindImageById(id)
	t.Log(f.ID, f.Filename, f.Date)
	assert.Equal(t, id, f.ID)

}
