package sim

import "time"

// Time layout to parse our input times.
// Using Go's "magic numbers":
// https://golang.org/pkg/time/#pkg-constants
const timeLayout = "2006-01-02 15:04:05"

// parseTime parses the given time string and swallows errors
func parseTime(raw string) (time.Time, error) {
	return time.Parse(timeLayout, raw)
}

// formatDuration formats the given duration
func formatDuration(t time.Duration) string {
	return t.String()
}
