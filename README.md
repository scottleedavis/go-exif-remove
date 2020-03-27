# go-exif-remove
[![Build Status](https://img.shields.io/circleci/project/github/scottleedavis/go-exif-remove/master.svg)](https://circleci.com/gh/scottleedavis/go-exif-remove) [![codecov](https://codecov.io/gh/scottleedavis/go-exif-remove/branch/master/graph/badge.svg)](https://codecov.io/gh/scottleedavis/go-exif-remove)  [![GoDoc](https://godoc.org/github.com/scottleedavis/go-exif-remove?status.svg)](https://godoc.org/github.com/scottleedavis/go-exif-remove)


Removes EXIF information from JPG and PNG files

Uses [go-exif](https://github.com/dsoprea/go-exif) to extract EXIF information and overwrites the EXIF region.

```go
import 	"github.com/scottleedavis/go-exif-remove"

noExifBytes, err := exifremove.Remove(imageBytes)
```

# installation
## install to PATH
go install github.com/aakarim/go-exif-remove

# example usage

```bash
#run against all in img folder
remove-exif

#run against single file
remove-exif img/jpg/11-tests.jpg
```
