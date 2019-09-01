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
	START_BYTES    = 4
	END_BYTES      = 4
	OFFSET_BYTES   = 4
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
			if _, err := handleFile(file); err != nil {
				fail += 1
				//fmt.Printf(err.Error())
			} else {
				pass += 1
			}
			fmt.Println()
		}

		math := 100 * pass / (pass+fail)
		fmt.Printf("Results (%v%%): %v pass, %v fail \n", int(math), pass, fail)
	} else {
		filepath := os.Args[1]
		handleFile(filepath)
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

		if _, sExif, err := sl.FindExif(); err != nil {
			return nil, err
		} else {
			//fmt.Printf("* (sExif) %x %s %x\n", sExif.Offset, sExif.MarkerName, len(sExif.Data))


			bytesCount := 0
			startExifBytes := START_BYTES
			endExifBytes := END_BYTES
			for _, s := range sl.Segments() {

				if s.MarkerName == sExif.MarkerName {
					if startExifBytes == START_BYTES {
						startExifBytes = bytesCount
						endExifBytes = startExifBytes + len(s.Data) + OFFSET_BYTES
					} else {
						endExifBytes += len(s.Data) + OFFSET_BYTES
					}
					fmt.Printf("* (sExif) %x %s %v (%x)\n", s.Offset, s.MarkerName, len(s.Data), s.Offset+len(s.Data))
				} else {
					fmt.Printf("* (s) %x %s %v (%x)\n", s.Offset, s.MarkerName, len(s.Data), s.Offset+len(s.Data))
				}
				bytesCount += len(s.Data) + OFFSET_BYTES


			}

			//filtered = data
			filtered = data[:startExifBytes]
			filtered = append(filtered, data[endExifBytes:]...)

			fmt.Printf("* (size) %v %v  (%v)\n", len(data), len(filtered), len(data)-len(filtered))

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
