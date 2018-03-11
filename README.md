# gproject
gproject is a library for getting to Google Cloud Platform's ProjectID.

## Installation
```bash
$ go get -u github.com/orisano/gproject
```

## How to use
```go
package main

import (
	"fmt"

	"github.com/orisano/gproject"
)

func main() {
	fmt.Println(gproject.Default())
}
```

## Method
https://github.com/GoogleCloudPlatform/google-auth-library-python/blob/master/google/auth/_default.py#L186


## Dependencies
Google Cloud SDK (optional)

## License
MIT
