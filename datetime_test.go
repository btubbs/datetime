package datetime

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidFormats(t *testing.T) {
	tt := []struct {
		input  string
		output time.Time
	}{
		{
			input:  "2007",
			output: time.Date(2007, time.January, 1, 0, 0, 0, 0, time.Local),
		},
		{
			input:  "2007-11",
			output: time.Date(2007, time.November, 1, 0, 0, 0, 0, time.Local),
		},
		{
			input:  "20071130",
			output: time.Date(2007, time.November, 30, 0, 0, 0, 0, time.Local),
		},
		{
			input:  "2007-11-30",
			output: time.Date(2007, time.November, 30, 0, 0, 0, 0, time.Local),
		},
		{
			input:  "2007-11-30T10",
			output: time.Date(2007, time.November, 30, 10, 0, 0, 0, time.Local),
		},
		{
			input:  "20071130T10",
			output: time.Date(2007, time.November, 30, 10, 0, 0, 0, time.Local),
		},
		{
			input:  "20071130T1010",
			output: time.Date(2007, time.November, 30, 10, 10, 0, 0, time.Local),
		},
		{
			input:  "20071130T101010",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.Local),
		},
		{
			input:  "20071130T101010.123",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 123000000, time.Local),
		},
		{
			input:  "2007-11-30T10:10",
			output: time.Date(2007, time.November, 30, 10, 10, 0, 0, time.Local),
		},
		{
			input:  "2007-11-30T10:10:10",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.Local),
		},
		{
			input:  "2007-11-30T10:10:10.123",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 123000000, time.Local),
		},
		{
			input:  "2007-11-30T10:10:10.000000001",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 1, time.Local),
		},
		{
			input:  "2007-11-30T10:10:10.123Z",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 123000000, time.UTC),
		},
		{
			input:  "20071130T10:10:10.123Z",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 123000000, time.UTC),
		},
		{
			input:  "20071130T10:10:10.0123Z",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 12300000, time.UTC),
		},
		{
			input:  "2007-11-30T10:10:10Z",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.UTC),
		},
		{
			input:  "2007-11-30T10:10:10+02",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.FixedZone("+02", 60*60*2)),
		},
		{
			input:  "2007-11-30T10:10:10-02",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.FixedZone("-02", 60*60*-2)),
		},
		{
			input:  "2007-11-30T10:10:10+0230",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.FixedZone("+0230", (60*60*2)+(30*60))),
		},
		{
			input:  "2007-11-30T10:10:10+02:30",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.FixedZone("+02:30", (60*60*2)+(30*60))),
		},
		{
			input:  "2007-11-30T10:10:10+00:00",
			output: time.Date(2007, time.November, 30, 10, 10, 10, 0, time.FixedZone("+00:00", 0)),
		},
		{
			input:  "2007-11-30T10Z",
			output: time.Date(2007, time.November, 30, 10, 0, 0, 0, time.UTC),
		},
		{
			input:  "2007T10",
			output: time.Date(2007, time.January, 1, 10, 0, 0, 0, time.Local),
		},
	}

	for _, tc := range tt {
		when, err := ParseTime(tc.input)
		assert.Equal(t, tc.output, when, tc.input)
		assert.Nil(t, err, tc.input)
	}
}

func TestErrors(t *testing.T) {
	tt := []struct {
		input string
		err   string
	}{
		{
			input: "asdf",
			err:   "found a, expected number",
		},
		{
			input: "2007-",
			err:   "found , expected number",
		},
		{
			input: "20077",
			err:   "found 20077, expected yyyy-mm-dd or yyyymmdd",
		},
		{
			input: "2007-13",
			err:   "13 is not a valid month",
		},
		{
			input: "2007-11-31",
			err:   "31 is not a valid day in November",
		},
		{
			input: "2007-11-30TA",
			err:   "expected number. got A",
		},
		{
			input: "2007-11-30A",
			err:   "found A, expected T or EOF",
		},
		{
			input: "2007-11-30T12B",
			err:   "expected colon, EOF, or timezone offset. got B",
		},
		{
			input: "2007-11-30T12:C",
			err:   "found C, expected number",
		},
		{
			input: "2007-11-10T25",
			err:   "25 is not a valid hour",
		},
		{
			input: "2007-11-30T12:112",
			err:   "112 is not a valid minute",
		},
		{
			input: "2007-11-30T12:11:70",
			err:   "70 is not a valid second",
		},
		{
			input: "2007-11-30T12:11:20.Q",
			err:   "expected fraction of seconds.  got Q",
		},
		{
			input: "2007-11-30T12:11:20.",
			err:   "expected fraction of seconds.  got ",
		},
		{
			input: "2007-11-30T12:11:20.456Q",
			err:   "expected Z, timezone offset, or EOF. got Q",
		},
		{
			input: "2007-11-30T12:11:20.456+111",
			err:   "expected ±hh:mm, ±hhmm, or ±hh timezone offset format. got 111",
		},
		{
			input: "2007-11-30T12:11:20.456+Q",
			err:   "expected number. got Q",
		},
		{
			input: "2007-11-30T10:10:10+02:30Q",
			err:   "expected EOF. got Q",
		},
		{
			input: "2007-11-30T10:10:10+02Q",
			err:   "expected colon or EOF. got Q",
		},
		{
			input: "2007-11-30T10:10:10+02:Q",
			err:   "expected number. got Q",
		},
		{
			input: "2007-11-30T10:10:Q",
			err:   "found Q, expected number",
		},
		{
			input: "2007Q",
			err:   "found Q, expected dash, T, or EOF",
		},
		{
			input: "2007-11Q",
			err:   "found Q, expected dash, T, or EOF",
		},
		{
			input: "2007-11-Q",
			err:   "found Q, expected number",
		},
		{
			input: "2007-11-10T1010101",
			err:   "expected time. got 1010101",
		},
	}

	for _, tc := range tt {
		when, err := ParseTime(tc.input)
		assert.Equal(t, zeroTime, when, tc.input)
		assert.Equal(t, errors.New(tc.err), err, tc.input)
	}
}

func TestParseInt(t *testing.T) {
	// most cases are tested higher up, but the panic case can't be, since all the values fed into
	// parseInt by callers are already checked as safe.  Test the panic case here.
	assert.Panics(t, func() { parseInt("") })
}

func TestRound(t *testing.T) {
	assert.Equal(t, 1.0, round(0.9))
}
