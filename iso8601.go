package iso8601

import (
	"fmt"
	"time"
)

// Format format a time to ISO8601 layout
func Format(t time.Time) string {
	t = t.UTC()

	return fmt.Sprintf(
		"%04d-%02d-%02dT%02d:%02d:%02d.%03dZ",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/int(time.Millisecond),
	)
}
