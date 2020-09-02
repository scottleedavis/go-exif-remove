你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# go-exif-remove
[![Build Status](https://img.shields.io/circleci/project/github/scottleedavis/go-exif-remove/master.svg)](https://circleci.com/gh/scottleedavis/go-exif-remove) [![codecov](https://codecov.io/gh/scottleedavis/go-exif-remove/branch/master/graph/badge.svg)](https://codecov.io/gh/scottleedavis/go-exif-remove)  [![GoDoc](https://godoc.org/github.com/scottleedavis/go-exif-remove?status.svg)](https://godoc.org/github.com/scottleedavis/go-exif-remove)


Removes EXIF information from JPG and PNG files

Uses [go-exif](https://github.com/dsoprea/go-exif) to extract EXIF information and overwrites the EXIF region.

```go
import 	"github.com/scottleedavis/go-exif-remove"

noExifBytes, err := exifremove.Remove(imageBytes)
```

See example usage in [exif-remove-tool](exif-remove-tool)

