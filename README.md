# datetime [![Build Status](https://travis-ci.org/btubbs/datetime.svg?branch=master)](https://travis-ci.org/btubbs/datetime) [![Coverage Status](https://coveralls.io/repos/github/btubbs/datetime/badge.svg?branch=master)](https://coveralls.io/github/btubbs/datetime?branch=master)

`datetime` provides a ParseTime function for turning commonly-used 
[ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) date/time formats into
Golang time.Time variables. 

Unlike Go's built-in RFC-3339 time format, this package automatically supports ISO 8601 date and
time stamps with varying levels of granularity.  Examples:

```go
package main

import (
	"fmt"

	"github.com/btubbs/datetime"
)

func main() {
	// just a year. Local time zone.
	fmt.Println(datetime.ParseTime("2007")) // 2007-01-01 00:00:00 -0700 MST <nil>

	// a year and a month
	fmt.Println(datetime.ParseTime("2007-11")) // 2007-11-01 00:00:00 -0600 MDT <nil>

	// a full date
	fmt.Println(datetime.ParseTime("2007-11-22")) // 2007-11-22 00:00:00 -0700 MST <nil>

	// adding time
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22")) // 2007-11-22 12:30:22 -0700 MST <nil>

	// omitting dashes and colons, as ISO 8601 allows
	fmt.Println(datetime.ParseTime("20071122T123022")) // 2007-11-22 12:30:22 -0700 MST <nil>

	// specifying a timezone offset
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22+0800")) // 2007-11-22 12:30:22 +0800 +0800 <nil>

	// adding separators to the offset too
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22+08:00")) // 2007-11-22 12:30:22 +0800 +08:00 <nil>

	// using a shorthand for UTC
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22Z")) // 2007-11-22 12:30:22 +0000 UTC <nil>
}

```

A `DateTime` type is also provided, which implements Scan, Value, and UnmarshalJSON methods for easy
de/serialization of ISO 8601 timestamps with external systems.
