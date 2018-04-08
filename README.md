# go-iso8601

[![Build Status](https://travis-ci.org/isayme/go-iso8601.svg?branch=master)](https://travis-ci.org/isayme/go-iso8601)

Time format with ISO8601 layout: `2006-01-02T15:04:05.999Z`

## Usage
```
package main

import (
  "fmt"
  "time"
  "github.com/isayme/go-iso8601"
)

func main() {
  t := time.Now()
  s := iso8601.Format(t)
  fmt.Println(s)
}
```

## Refer
- [Javascript Date.prototype.toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
- [Wiki ISO8601](https://en.wikipedia.org/wiki/ISO_8601)
