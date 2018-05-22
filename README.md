# datetime [![Build Status](https://travis-ci.org/btubbs/datetime.svg?branch=master)](https://travis-ci.org/btubbs/datetime) [![Coverage Status](https://coveralls.io/repos/github/btubbs/datetime/badge.svg?branch=master)](https://coveralls.io/github/btubbs/datetime?branch=master)

`datetime` provides a Parse function for turning commonly-used 
[ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) date/time formats into
Golang time.Time variables.  Unlike Go's built-in RFC-3339 time format, this package support's date
and time stamps with varying levels of granularity.

It also provides a DateTime type that wraps the aforementioned Parse function in the golang
Unmarshaler interface, so you can use it as a struct field to parse a wide variety of timestamps in
JSON payloads.
