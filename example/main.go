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

	// fractions of a second
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22.321")) // 2007-11-22 12:30:22.321 -0700 MST <nil>

	// omitting dashes and colons, as ISO 8601 allows
	fmt.Println(datetime.ParseTime("20071122T123022")) // 2007-11-22 12:30:22 -0700 MST <nil>

	// specifying a timezone offset
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22+0800")) // 2007-11-22 12:30:22 +0800 +0800 <nil>

	// adding separators to the offset too
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22+08:00")) // 2007-11-22 12:30:22 +0800 +08:00 <nil>

	// using a shorthand for UTC
	fmt.Println(datetime.ParseTime("2007-11-22T12:30:22Z")) // 2007-11-22 12:30:22 +0000 UTC <nil>
}
