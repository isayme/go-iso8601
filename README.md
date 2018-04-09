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
  fmt.Println(s) // 2018-04-09T15:23:54.595Z
}
```

## Use for time.Time json format
```
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/isayme/go-iso8601"
)

type Person struct {
	Birthday time.Time
}

func (p Person) MarshalJSON() ([]byte, error) {
	type _person Person
	return json.Marshal(&struct {
		Birthday iso8601.Time
	}{
		Birthday: (iso8601.Time)(p.Birthday),
	})
}

func main() {
	p := Person{
		Birthday: time.Unix(0, 0),
	}

	b, _ := json.Marshal(p)
	fmt.Println(string(b)) // {"Birthday":"1970-01-01T00:00:00.000Z"}
}
```

## Refer
- [Javascript Date.prototype.toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
- [Wiki ISO8601](https://en.wikipedia.org/wiki/ISO_8601)
- [Custom JSON Marshalling in Go](http://choly.ca/post/go-json-marshalling/)