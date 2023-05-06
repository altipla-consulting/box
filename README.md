
# box

[![Go Reference](https://pkg.go.dev/badge/github.com/altipla-consulting/box.svg)](https://pkg.go.dev/github.com/altipla-consulting/box)

Renders messages inside boxes in the terminal.


## Install

```shell
go get github.com/altipla-consulting/box
```


### Usage

```go
package main

import (
	"github.com/altipla-consulting/box"
	"github.com/logrusorgru/aurora"
)

func main() {
	var o box.Box
	o.AddLine("foo", "bar")
	o.AddLine("before", aurora.Red("colored"), "after")
	o.Render()
}
```


## Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `make gofmt`.


## License

[MIT License](LICENSE)
