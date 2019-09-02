package exifremove

import (
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
	pngstructure "github.com/dsoprea/go-png-image-structure"
)

func TestRemove(t *testing.T) {

	data, err := ioutil.ReadFile("exif-remove-tool/img/jpg/11-tests.jpg")
	assert.Nil(t, err)
	filtered, err := Remove(data)
	assert.Nil(t, err)
	jmp := jpegstructure.NewJpegMediaParser()
	sl, err := jmp.ParseBytes(filtered)
	_, _, err = sl.Exif()
	assert.NotNil(t, err)


	data, err = ioutil.ReadFile("exif-remove-tool/img/png/exif.png")
	assert.Nil(t, err)
	filtered, err = Remove(data)
	assert.Nil(t, err)
	pmp := pngstructure.NewPngMediaParser()
	cs, err := pmp.ParseBytes(filtered)
	_, _, err = cs.Exif()
	assert.NotNil(t, err)


}
