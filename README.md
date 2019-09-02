# go-exif-remove
remove EXIF information from JPG and PNG files

Uses [go-exif](https://github.com/dsoprea/go-exif) to extract EXIF information and fills the region with 0's.

```go
noExifBytes, err := exifremove.RemoveEXIF(imageBytes)
```

_TODO png_
