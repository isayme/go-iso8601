package iso8601

import (
	"path"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func assert(t *testing.T, exp, got interface{}) {
	if reflect.DeepEqual(exp, got) == false {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("Expecting '%v' got '%v' @%s:%d\n", exp, got, path.Base(file), line)
	}
}

func TestISO8601(t *testing.T) {
	t.Run("time.Unix(0, 0)", func(t *testing.T) {
		s := Format(time.Unix(0, 0))
		assert(t, "1970-01-01T00:00:00.000Z", s)
	})

	t.Run("time.Unix(0, 1000)", func(t *testing.T) {
		s := Format(time.Unix(0, 1000))
		assert(t, "1970-01-01T00:00:00.000Z", s)
	})

	t.Run("time.Unix(0, 100000)", func(t *testing.T) {
		s := Format(time.Unix(0, 100000))
		assert(t, "1970-01-01T00:00:00.000Z", s)
	})

	t.Run("time.Unix(0, 1000000)", func(t *testing.T) {
		s := Format(time.Unix(0, 1000000))
		assert(t, "1970-01-01T00:00:00.001Z", s)
	})

	t.Run("time.Unix(0, 10000000)", func(t *testing.T) {
		s := Format(time.Unix(0, 10000000))
		assert(t, "1970-01-01T00:00:00.010Z", s)
	})

	t.Run("time.Unix(0, 100000000)", func(t *testing.T) {
		s := Format(time.Unix(0, 100000000))
		assert(t, "1970-01-01T00:00:00.100Z", s)
	})

	t.Run("time.Unix(0, 1523199549840ms)", func(t *testing.T) {
		s := Format(time.Unix(0, 1523199549840*int64(time.Millisecond)))
		assert(t, "2018-04-08T14:59:09.840Z", s)
	})

	t.Run("tz", func(t *testing.T) {
		ts := time.Unix(0, 0)
		loc, _ := time.LoadLocation("Asia/Shanghai")
		ts = ts.In(loc)

		s := Format(ts)
		assert(t, "1970-01-01T00:00:00.000Z", s)

		loc, _ = time.LoadLocation("America/New_York")
		ts = ts.In(loc)
		s = Format(ts)
		assert(t, "1970-01-01T00:00:00.000Z", s)
	})
}
