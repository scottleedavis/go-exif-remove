package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dsoprea/go-exif"
	"github.com/dsoprea/go-jpeg-image-structure"
	"github.com/dsoprea/go-png-image-structure"
)

const (
	JpegMediaType  = "jpeg"
	PngMediaType   = "png"
	OtherMediaType = "other"
	StartBytes     = 4
	EndBytes       = 4
	OffsetBytes    = 4
)

type MediaContext struct {
	MediaType string
	RootIfd   *exif.Ifd
	RawExif   []byte
	Media     interface{}
}

func main() {

	if len(os.Args) == 1 {
		var files []string
		root := "img"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if path != "img" {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		pass := 0
		fail := 0
		for _, file := range files {
			fmt.Println(file)
			if b, err := handleFile(file); err != nil {
				fail += 1
			} else {
				pass += 1
				f := filepath.Base(file)
				ioutil.WriteFile("img_output/"+f, b, 0644)
			}
			fmt.Println()
		}

		percentage := 100 * pass / (pass + fail)
		fmt.Printf("Results (%v%%): %v pass, %v fail \n", int(percentage), pass, fail)
	} else {
		path := os.Args[1]
		if b, err := handleFile(path); err != nil {
			fmt.Printf(err.Error())
		} else {
			file := filepath.Base(path)
			ioutil.WriteFile("img_output/"+file, b, 0644)
		}
	}

}

func handleFile(filepath string) ([]byte, error) {
	if data, err := ioutil.ReadFile(filepath); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	} else {
		_, err = jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Printf("ERROR: original image is corrupt" + err.Error() + "\n")
			return nil, err
		}
		filtered, err := extractEXIF(data)
		if err != nil {
			if !strings.EqualFold(err.Error(), "no exif data") {
				fmt.Printf("* " + err.Error() + "\n")
				return nil, errors.New(err.Error())
			}
		}
		return filtered, nil
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

		if _, _, err := sl.FindExif(); err != nil {
			return nil, err
		} else {

			startExifBytes := StartBytes
			endExifBytes := EndBytes

			if bytes.Contains(data, mc.RawExif) {
				for i := 0; i < len(data)-len(mc.RawExif); i++ {
					if bytes.Compare(data[i:i+len(mc.RawExif)], mc.RawExif) == 0 {
						startExifBytes = i
						endExifBytes = i + len(mc.RawExif)
					}
				}
			}

			fill := make([]byte, len(data[startExifBytes:endExifBytes]))
			copy(data[startExifBytes:endExifBytes], fill)
			filtered = data

			_, _, err = image.Decode(bytes.NewReader(filtered))
			if err != nil {
				return nil, errors.New("EXIF removal corrupted " + err.Error())
			}

		}

	} else if pmp.LooksLikeFormat(data) {
		mc.MediaType = PngMediaType
	}

	return filtered, nil
}
