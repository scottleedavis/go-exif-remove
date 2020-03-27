package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"

	exifremove "github.com/aakarim/remove-exif/exif_remove"
)

func main() {
	// create output folder in current directory if none
	outputPath := filepath.Join(".", "img_output")
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.Mkdir(outputPath, os.ModePerm)
	}
	if len(os.Args) == 1 {
		var files []string
		root := "img"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if path != "img" && path != "img/png" && path != "img/jpg" {
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
				ioutil.WriteFile(filepath.Join(outputPath, f), b, 0644)
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
			ioutil.WriteFile(filepath.Join(outputPath, file), b, 0644)
			fmt.Println("Removed EXIF data successfully")
		}
	}

}

func handleFile(filepath string) ([]byte, error) {
	if data, err := ioutil.ReadFile(filepath); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	} else {
		if _, _, err := image.Decode(bytes.NewReader(data)); err != nil {
			fmt.Printf("ERROR: original image is corrupt" + err.Error() + "\n")
			return nil, err
		}
		if filtered, err := exifremove.Remove(data); err != nil {
			fmt.Printf("* " + err.Error() + "\n")
			return nil, errors.New(err.Error())
		} else {
			if _, _, err = image.Decode(bytes.NewReader(filtered)); err != nil {
				fmt.Printf("ERROR: filtered image is corrupt" + err.Error() + "\n")
				return nil, err
			}
			return filtered, nil
		}
	}
}
