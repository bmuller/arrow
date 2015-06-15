# Arrow: Go Date Parsing

[![GoDoc](https://godoc.org/github.com/bmuller/arrow/lib?status.png)](https://godoc.org/github.com/bmuller/arrow/lib)
[![Build Status](https://travis-ci.org/bmuller/arrow.png?branch=master)](https://travis-ci.org/bmuller/arrow)

Arrow provides a C style format based parsing in Golang (among other helpful date/time functions).

```
Time flies like an arrow; fruit flies like a banana.
```

# Sprintf Compatability
The problem with formatting times in Golang is that the [format string you give](http://golang.org/pkg/time/#Time.Format) is based on re-formatting a single date that is a pain to remember (I know it's 1/2 3:04:05 2006 -0700, but [I'm lazy](http://threevirtues.com/)).  Most languages based on C (Python, etc) use a string formatting based on [sprintf](http://man7.org/linux/man-pages/man3/strftime.3.html), which is what I think most people are familiar with.

So here's `sprintf` compatability for Golang with `CFormat` and `CParse`:

```go
package main

import (
	"fmt"

	"github.com/voxmedia/arrow/lib"
)

func main() {
     // formatting
     fmt.Println("Current date: ", arrow.Now().CFormat("%Y-%m-%d %H:%M"))

     // parsing
     fmt.Println("Some other date: ", arrow.CParse("%Y-%m-%d", "2015-06-03"))
}
```

# Additional Fun
You can also utilize helpful functions to get things like the beginning of the minute, hour, day, week, month, and year.

```go
func main() {
     day := arrow.Now().AtBeginningOfWeek().CFormat("%Y-%m-%d")
     fmt.Println("First day of week: ", day)

     hour := arrow.Now().AtBeginningOfHour().CFormat("%Y-%m-%d %H:%M")
     fmt.Println("First second of hour: ", hour)
}
```

You can also more easily sleep until specific times:

```go
func main() {
     // sleep until the next minute starts
     arrow.SleepUntil(arrow.NextMinute())
     fmt.Println(arrow.Now().CFormat("%H:%M:%S"))
}
```

There are also helpers to get today, yesterday, and UTC times:

```go
func main() {
     day := arrow.Yesterday().CFormat("%Y-%m-%d")
     fmt.Println("Yesterday: ", day)

     dayutc := arrow.UTC().Yesterday().CFormat("%Y-%m-%d %H:%M")
     fmt.Println("Yesterday, UTC: ", dayutc)

     newyork := arrow.InTimezone("America/New_York").CFormat("%H:%M:%s")
     fmt.Println("Time in New York: ", newyork)
}
```

And for generating ranges when you need to interate:

```go
func main() {
     // Print every minute from now until 24 hours from now
     for _, a := range arrow.Now().UpTo(arrow.Tomorrow(), arrow.Minute) {
          fmt.Println(a.CFormat("%Y-%m-%d %H:%M:%S"))
     }
}
```
