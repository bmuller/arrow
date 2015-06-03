package arrow

import (
	"strconv"
	"strings"
	"time"
)

type Arrow struct {
	time.Time
}

// Like time's constants, but with Day and Week
const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
	Week                      = 7 * Day
)

func New(t time.Time) Arrow {
	return Arrow{t}
}

func UTC() Arrow {
	return New(time.Now().UTC())
}

func Now() Arrow {
	return New(time.Now())
}

func (a Arrow) UTC() Arrow {
	return Arrow{a.Time.UTC()}
}

func (a Arrow) Sub(b Arrow) time.Duration {
	return a.Time.Sub(b.Time)
}

// Add any duration parseable by time.ParseDuration
func (a Arrow) AddDuration(duration string) Arrow {
	if pduration, err := time.ParseDuration(duration); err == nil {
		return Arrow{a.Add(pduration)}
	}
	return a
}

func (a Arrow) AddDays(days int) Arrow {
	return Arrow{a.AddDate(0, 0, days)}
}

func (a Arrow) AtBeginningOfMinute() Arrow {
	return Arrow{a.Truncate(Minute)}
}

func (a Arrow) AtBeginningOfHour() Arrow {
	return Arrow{a.Truncate(Hour)}
}

func (a Arrow) AtBeginningOfDay() Arrow {
	d := time.Duration(-a.Hour()) * Hour
	return Arrow{a.AtBeginningOfHour().Add(d)}
}

func (a Arrow) AtBeginningOfWeek() Arrow {
	days := time.Duration(-1*int(a.Weekday())) * Day
	return Arrow{a.AtBeginningOfDay().Add(days)}
}

func (a Arrow) AtBeginningOfMonth() Arrow {
	days := time.Duration(-1*int(a.Day())+1) * Day
	return Arrow{a.AtBeginningOfDay().Add(days)}
}

func (a Arrow) AtBeginningOfYear() Arrow {
	days := time.Duration(-1*int(a.YearDay())+1) * Day
	return Arrow{a.AtBeginningOfDay().Add(days)}
}

// Add any durations parseable by time.ParseDuration
func (a Arrow) AddDurations(durations []string) Arrow {
	for _, duration := range durations {
		a = a.AddDuration(duration)
	}
	return a
}

func formatConvert(format string) string {
	// create mapping from strftime to time in Go
	strftimeMapping := map[string]string{
		"%a": "Mon",
		"%A": "Monday",
		"%b": "Jan",
		"%B": "January",
		"%c": "", // locale not supported
		"%C": "06",
		"%d": "02",
		"%D": "01/02/06",
		"%e": "_2",
		"%E": "", // modifiers not supported
		"%F": "2006-01-02",
		"%G": "%G", // special case, see below
		"%g": "%g", // special case, see below
		"%h": "Jan",
		"%H": "15",
		"%I": "03",
		"%j": "%j", // special case, see below
		"%k": "%k", // special case, see below
		"%l": "_3",
		"%m": "01",
		"%M": "04",
		"%n": "\n",
		"%O": "", // modifiers not supported
		"%p": "PM",
		"%P": "pm",
		"%r": "03:04:05 PM",
		"%R": "15:04",
		"%s": "%s", // special case, see below
		"%S": "05",
		"%t": "\t",
		"%T": "15:04:05",
		"%u": "%u", // special case, see below
		"%U": "%U", // special case, see below
		"%V": "%V", // special case, see below
		"%w": "%w", // special case, see below
		"%W": "%W", // special case, see below
		"%x": "%x", // locale not supported
		"%X": "%X", // locale not supported
		"%y": "06",
		"%Y": "2006",
		"%z": "-0700",
		"%Z": "MST",
		"%+": "Mon Jan _2 15:04:05 MST 2006",
		"%%": "%%", // special case, see below
	}

	for fmt, conv := range strftimeMapping {
		format = strings.Replace(format, fmt, conv, -1)
	}

	return format
}

func CParse(layout, value string) (Arrow, error) {
	t, e := time.Parse(formatConvert(layout), value)
	return Arrow{t}, e
}

func CParseInLocation(layout, value string, loc *time.Location) (Arrow, error) {
	t, e := time.ParseInLocation(formatConvert(layout), value, loc)
	return Arrow{t}, e
}

func (a Arrow) CFormat(format string) string {
	format = a.Format(formatConvert(format))

	year, week := a.ISOWeek()
	yearday := a.YearDay()
	weekday := a.Weekday()
	syear := strconv.Itoa(year)
	sweek := strconv.Itoa(week)
	syearday := strconv.Itoa(yearday)
	sweekday := strconv.Itoa(int(weekday))

	if a.Year() > 999 {
		format = strings.Replace(format, "%G", syear, -1)
		format = strings.Replace(format, "%g", syear[2:4], -1)
	}

	format = strings.Replace(format, "%j", syearday, -1)
	if a.Hour() < 10 {
		shour := " " + strconv.Itoa(a.Hour())
		format = strings.Replace(format, "%k", shour, -1)
	}
	format = strings.Replace(format, "%s", strconv.FormatInt(a.Unix(), 10), -1)

	if weekday == 0 {
		format = strings.Replace(format, "%u", "7", -1)
	} else {
		format = strings.Replace(format, "%u", sweekday, -1)
	}

	format = strings.Replace(format, "%U", weekNumber(a, time.Sunday), -1)
	format = strings.Replace(format, "%U", sweek, -1)
	format = strings.Replace(format, "%w", sweekday, -1)
	format = strings.Replace(format, "%W", weekNumber(a, time.Monday), -1)
	return strings.Replace(format, "%%", "%", -1)
}

// Used for %U and %W:
// %U: The week number of the current year as a decimal number, range
// 00 to 53, starting with the first Sunday as the first day of week 01.
//
// %W: The week number of the current year as a decimal number, range
// 00 to 53, starting with the first Monday as the first day of week 01.
func weekNumber(a Arrow, firstday time.Weekday) string {
	dayone := a.AtBeginningOfYear()
	for dayone.Weekday() != time.Sunday {
		dayone = dayone.AddDays(1)
	}
	week := int(a.Sub(dayone.AddDays(-7)) / Week)
	return strconv.Itoa(week)
}
