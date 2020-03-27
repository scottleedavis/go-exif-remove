module github.com/aakarim/remove-exif

go 1.14

require (
	github.com/aakarim/remove-exif/exif_remove v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.5.1
)

replace github.com/aakarim/remove-exif/exif_remove => ./exif_remove
