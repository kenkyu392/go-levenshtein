# go-levenshtein

[![test](https://github.com/kenkyu392/go-levenshtein/workflows/test/badge.svg?branch=master)](https://github.com/kenkyu392/go-levenshtein)
[![codecov](https://codecov.io/gh/kenkyu392/go-levenshtein/branch/master/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-levenshtein)
[![GoDoc](https://godoc.org/github.com/kenkyu392/go-levenshtein?status.svg)](https://godoc.org/github.com/kenkyu392/go-levenshtein)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenkyu392/go-levenshtein)](https://goreportcard.com/report/github.com/kenkyu392/go-levenshtein)

Go implementation of [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance).

## Installation

```
go get -u github.com/kenkyu392/go-levenshtein
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/kenkyu392/go-levenshtein"
)

func main() {
	var (
		a = []rune("I Am a Cat")
		b = []rune("I Am a Dog")
	)
	p := levenshtein.Percent(a, b)
	c := levenshtein.Distance(a, b)
	// -> 70% match. 3 character difference.
	fmt.Printf("%.0f%% match. %d character difference.", p, c)
}
```

## License

[MIT](https://github.com/kenkyu392/go-levenshtein/blob/master/LICENSE)
