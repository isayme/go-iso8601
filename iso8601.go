package iso8601

import (
	"fmt"
	"time"
)

// Time time with iso8601 json format
type Time time.Time

// MarshalJSON custom MarshalJSON
func (ts Time) MarshalJSON() ([]byte, error) {
	t := (time.Time)(ts)
	s := Format(t)

	b := make([]byte, 0, 26)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// Format format a time to ISO8601 layout
func Format(t time.Time) string {
	utc := t.UTC()

	return fmt.Sprintf(
		"%04d-%02d-%02dT%02d:%02d:%02d.%03dZ",
		utc.Year(),
		utc.Month(),
		utc.Day(),
		utc.Hour(),
		utc.Minute(),
		utc.Second(),
		utc.Nanosecond()/int(time.Millisecond),
	)
}
