package arrow

import (
	"testing"
	"time"
)

func TestCParse(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	a, _ := CParseInLocation("%Y-%m-%d", "1980-06-19", loc)
	if a.Format("2006-01-02") != "1980-06-19" {
		t.Error("CParseInLocation")
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
