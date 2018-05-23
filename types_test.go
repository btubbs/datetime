package datetime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newDefaultUTC(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) DefaultUTC {
	return DefaultUTC(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

func TestParse(t *testing.T) {

	tt := []struct {
		input  string
		loc    *time.Location
		output time.Time
	}{
		{
			input:  "2007-11-11T17:38:12.000000001Z",
			loc:    time.UTC,
			output: time.Date(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
		},
	}
	for _, tc := range tt {
		tyme, err := Parse(tc.input, tc.loc)
		assert.Equal(t, tc.output, tyme)

		assert.Nil(t, err)
	}
}

func TestString(t *testing.T) {
	tt := []struct {
		dt     DefaultUTC
		output string
	}{
		{
			dt:     newDefaultUTC(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
			output: "2007-11-11T17:38:12.432Z",
		},
		{
			dt:     newDefaultUTC(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
			output: "2007-11-11T17:38:12.000000001Z",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.output, tc.dt.String())
	}
}

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		input []byte
		dt    DefaultUTC
		err   error
	}{
		{
			input: []byte(`"2007-11-11T17:38:12.432Z"`),
			dt:    newDefaultUTC(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
		},
		{
			input: []byte(`"2007-11-11T17:38:12.000000001Z"`),
			dt:    newDefaultUTC(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
		},
		{
			input: []byte(`2007`),
			err:   errors.New("2007 does not begin and end with double quotes"),
		},
		{
			input: []byte(`"A"`),
			err:   errors.New("found A, expected number"),
		},
	}

	for _, tc := range tt {
		var dt DefaultUTC
		err := json.Unmarshal(tc.input, &dt)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.dt, dt)
	}
}

func TestScan(t *testing.T) {
	tt := []struct {
		input interface{}
		dt    DefaultUTC
		err   error
	}{
		{
			input: []byte("2007-11-11T17:38:12.432Z"),
			dt:    newDefaultUTC(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
		},
		{
			input: []byte("invalid"),
			err:   errors.New("found i, expected number"),
		},
		{
			input: "2007-11-11T17:38:12.000000001Z",
			dt:    newDefaultUTC(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
		},
		{
			input: "invalid",
			err:   errors.New("found i, expected number"),
		},
		{
			input: 2007,
			err:   errors.New("can only scan string and []byte, not int"),
		},
	}

	for _, tc := range tt {
		var dt DefaultUTC
		err := dt.Scan(tc.input)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.dt, dt)
	}
}

func TestValue(t *testing.T) {
	tt := []struct {
		dt     DefaultUTC
		output driver.Value
	}{
		{
			dt:     newDefaultUTC(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
			output: "2007-11-11T17:38:12.432Z",
		},
		{
			dt:     newDefaultUTC(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
			output: "2007-11-11T17:38:12.000000001Z",
		},
	}

	for _, tc := range tt {
		val, err := tc.dt.Value()
		assert.Nil(t, err)
		assert.Equal(t, tc.output, val)
	}
}
