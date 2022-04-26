package arrow

import (
	"testing"
	"time"
)

func loadLocation(name string) *time.Location {
	if loc, err := time.LoadLocation(name); err == nil {
		return loc
	}

	return nil
}

func TestCParse(t *testing.T) {
	tests := []struct {
		dateStr  string
		timezone string
		format   string
		expected time.Time
	}{
		{
			dateStr:  "1980-06-19",
			timezone: "America/New_York",
			format:   "%Y-%m-%d",
			expected: time.Date(1980, time.June, 19, 0, 0, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-19",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d",
			expected: time.Date(1980, time.June, 19, 0, 0, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-19 9:02",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d %-I:%M",
			expected: time.Date(1980, time.June, 19, 9, 2, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-19 9:02",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d %-H:%M",
			expected: time.Date(1980, time.June, 19, 9, 2, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-19 09:02",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d %-H:%M",
			expected: time.Date(1980, time.June, 19, 9, 2, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-19 9:02",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d %H:%M",
			expected: time.Date(1980, time.June, 19, 9, 2, 0, 0, loadLocation("America/New_York")),
		},
		{
			dateStr:  "1980-6-0109:02",
			timezone: "America/New_York",
			format:   "%Y-%-m-%d%H:%M",
			expected: time.Date(1980, time.June, 1, 9, 2, 0, 0, loadLocation("America/New_York")),
		},
	}
	for _, test := range tests {
		loc, _ := time.LoadLocation(test.timezone)
		actual, err := CParseInLocation(test.format, test.dateStr, loc)
		if err != nil {
			t.Errorf("Error parsing date: %s", err)
		}
		if !actual.Time.Equal(test.expected) {
			t.Errorf("CParseInLocation(%s, %s, %s): Expected %s, got %s", test.format, test.dateStr, test.timezone, test.expected, actual)
		}
	}
}

func TestCFormat(t *testing.T) {
	tests := []struct {
		datetime time.Time
		format   string
		expected string
	}{
		{
			datetime: time.Date(1980, time.June, 19, 0, 0, 0, 0, time.UTC),
			format:   "%Y-%m-%d",
			expected: "1980-06-19",
		},
		{
			datetime: time.Date(2022, time.January, 8, 0, 0, 0, 0, time.UTC),
			format:   "%Y-%-m-%-d",
			expected: "2022-1-8",
		},
		{
			datetime: time.Date(2022, time.January, 8, 21, 2, 0, 0, time.UTC),
			format:   "%Y-%-m-%-d %-I:%-M",
			expected: "2022-1-8 9:2",
		},
		{
			datetime: time.Date(2022, time.January, 8, 4, 2, 0, 0, time.UTC),
			format:   "%Y-%-m-%-d %-H:%-M",
			expected: "2022-1-8 4:2",
		},
		{
			datetime: time.Date(2022, time.January, 8, 13, 2, 0, 0, time.UTC),
			format:   "%Y-%-m-%-d %-H:%-M",
			expected: "2022-1-8 13:2",
		},
		{
			datetime: time.Date(2006, time.August, 9, 23, 59, 11, 0, time.UTC),
			format:   "%Y-%-m-%-dT%H:%M:%S",
			expected: "2006-8-9T23:59:11",
		},
	}
	for _, test := range tests {
		a := New(test.datetime)
		actual := a.CFormat(test.format)
		if actual != test.expected {
			t.Errorf("CFormat(%s) returned %s, expected %s", test.format, actual, test.expected)
		}
	}
}
