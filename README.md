# envpath

Small Go library for appending and prepending paths to a PATH var.

It annoyed me that [`filepath.SplitList`](https://pkg.go.dev/path/filepath#SplitList) exists, but nothing to append or prepend strings to a PATH formatted environment variable.

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/bigkevmcd/envpath"
)

func main() {
	path := os.Getenv("PATH")

	fmt.Println(envpath.Append(path, "/usr/local/sbin"))
}
```
