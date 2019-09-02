# go-exif-remove
Removes EXIF information from JPG and PNG files

Uses [go-exif](https://github.com/dsoprea/go-exif) to extract EXIF information and overwrites the EXIF region.

```go
import 	"github.com/scottleedavis/go-exif-remove"

noExifBytes, err := exifremove.Remove(imageBytes)
```

See example usage in [exif-remove-tool](exif-remove-tool)

