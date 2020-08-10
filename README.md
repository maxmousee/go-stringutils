# stringutils

[![Build Status](https://travis-ci.com/maxmousee/stringutils.svg?branch=master)](https://travis-ci.org/maxmousee/stringutils)
[![Go Report](https://goreportcard.com/badge/github.com/maxmousee/stringutils)](https://goreportcard.com/report/github.com/maxmousee/stringutils)
[![Coverage Status](https://coveralls.io/repos/github/maxmousee/stringutils/badge.svg?branch=master)](https://coveralls.io/github/maxmousee/stringutils?branch=master)
[![GoDoc](https://godoc.org/github.com/maxmousee/stringutils?status.svg)](https://godoc.org/github.com/maxmousee/stringutils)

StringUtils for Go! :) 

StringUtils contains common utility functions available in other programming languages and frameworks
to make it easier for you to manipulate strings, create a parser or a lexer in your project.

------

Installation
============

To install stringutils, use `go get`:

    go get github.com/maxmousee/stringutils

This will then make the following package(s) available to you:

    github.com/maxmousee/stringutils

Import the `maxmousee/stringutils` package into your code using this template:

```go
package main

import (
	"fmt"
	"github.com/maxmousee/stringutils"
)

func main() {
	equalsIgoringCase := stringutils.EqualsIgnoreCase("a", "A")
	fmt.Println(equalsIgoringCase)
}
```

------

Staying up to date
==================

To update stringutils to the latest version, use `go get -u github.com/maxmousee/stringutils`.

------

Supported go versions
==================

Support for the last major Go version.

------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, please include a complete test function that demonstrates the issue.


------

License
=======

This project is licensed under the terms of the MIT license.