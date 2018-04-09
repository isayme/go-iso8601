package iso8601

import (
	"fmt"
	"time"
)

// Format format a time to ISO8601 layout
func Format(t *time.Time) *string {
	if t == nil {
		return nil
	}

	utc := t.UTC()

	s := fmt.Sprintf(
		"%04d-%02d-%02dT%02d:%02d:%02d.%03dZ",
		utc.Year(),
		utc.Month(),
		utc.Day(),
		utc.Hour(),
		utc.Minute(),
		utc.Second(),
		utc.Nanosecond()/int(time.Millisecond),
	)

	return &s
}
