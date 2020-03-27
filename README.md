# Remove EXIF
![Go](https://github.com/aakarim/remove-exif/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/aakarim/remove-exif/exif_remove?status.svg)](https://godoc.org/github.com/aakarim/remove-exif/exif_remove)


Removes EXIF information from JPG and PNG files

Uses [go-exif](https://github.com/dsoprea/go-exif) to extract EXIF information and overwrites the EXIF region.

# Tool
## Installation into $PATH
`go install github.com/aakarim/remove-exif`

# Example Usage

```bash
#run against all in img folder
remove-exif

#run against single file
remove-exif img/jpg/11-tests.jpg
```

# Library
## Installation
`go get -u github.com/aakarim/remove-exif/exif_remove`
## Example Usage
```go
import 	"github.com/aakarim/remove-exif/exif_remove"

noExifBytes, err := exifremove.Remove(imageBytes)
```
