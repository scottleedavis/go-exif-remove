package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/dsoprea/go-exif"
	"github.com/dsoprea/go-jpeg-image-structure"
	"github.com/dsoprea/go-png-image-structure"
)

const (
	JpegMediaType  = "jpeg"
	PngMediaType   = "png"
	OtherMediaType = "other"
)

type MediaContext struct {
	MediaType string
	RootIfd   *exif.Ifd
	RawExif   []byte
	Media     interface{}
}

type IfdEntry struct {
	IfdPath     string      `json:"ifd_path"`
	FqIfdPath   string      `json:"fq_ifd_path"`
	IfdIndex    int         `json:"ifd_index"`
	TagId       uint16      `json:"tag_id"`
	TagName     string      `json:"tag_name"`
	TagTypeId   uint16      `json:"tag_type_id"`
	TagTypeName string      `json:"tag_type_name"`
	UnitCount   uint32      `json:"unit_count"`
	Value       interface{} `json:"value"`
	ValueString string      `json:"value_string"`
}

func main() {
	file := os.Args[1]
	if data, err := ioutil.ReadFile(file); err != nil {
		fmt.Printf(err.Error())
		return
	} else {
		if filtered, err := extractEXIF(data); err != nil {
			fmt.Printf(err.Error())
			return
		} else {
			if err := ioutil.WriteFile("img_output/test.jpg", filtered, 0644); err != nil {
				fmt.Printf(err.Error())
			}
		}
	}
}

func extractEXIF(data []byte) ([]byte, error) {
	jmp := jpegstructure.NewJpegMediaParser()
	pmp := pngstructure.NewPngMediaParser()
	mc := &MediaContext{
		MediaType: OtherMediaType,
		RootIfd:   nil,
		RawExif:   nil,
		Media:     nil,
	}
	filtered := []byte{}

	if jmp.LooksLikeFormat(data) {
		mc.MediaType = JpegMediaType
		sl, err := jmp.ParseBytes(data)
		if err != nil {
			return nil, err
		}
		mc.Media = sl

		if rootIfd, rawExif, err := sl.Exif(); err != nil {
			return nil, err
		} else {
			mc.RootIfd = rootIfd
			mc.RawExif = rawExif
		}

		if _, sExif, err := sl.FindExif(); err != nil {
			return nil, err
		} else {
			fmt.Printf("****(exif) %x %s %x\n", sExif.Offset, sExif.MarkerName, len(sExif.Data))

			bytesCount := 0
			startExifBytes := 4
			endExifBytes := 4
			for _, s := range sl.Segments() {

				if s.MarkerName == sExif.MarkerName {
					if startExifBytes == 4 {
						startExifBytes = bytesCount
						endExifBytes = startExifBytes + len(s.Data)
					} else {
						endExifBytes += len(s.Data)
					}
				}
				bytesCount += len(s.Data)

				fmt.Printf("%x %s %v (%x)\n", s.Offset, s.MarkerName, len(s.Data), s.Offset+len(s.Data))

			}

			filtered = data[:startExifBytes]
			filtered = append(filtered, data[endExifBytes:]...)

			fmt.Printf("********(size) %v %v  (%v)\n", len(data), len(filtered), len(data)-len(filtered))

			_, _, err = image.Decode(bytes.NewReader(filtered))
			if err != nil {
				return nil, errors.New("EXIF extraction corrupted image " + err.Error() + "\n")
			}

		}

	} else if pmp.LooksLikeFormat(data) {
		mc.MediaType = PngMediaType
	}

	return filtered, nil
}
