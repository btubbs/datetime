package datetime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// a simple test, for a simple function
	dt, err := Parse("2007-11-11T17:38:12.000000001Z")
	assert.Equal(t,
		New(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
		dt,
	)

	assert.Nil(t, err)
}

func TestString(t *testing.T) {
	tt := []struct {
		dt     DateTime
		output string
	}{
		{
			dt:     New(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
			output: "2007-11-11T17:38:12.432Z",
		},
		{
			dt:     New(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
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
		dt    DateTime
		err   error
	}{
		{
			input: []byte(`"2007-11-11T17:38:12.432Z"`),
			dt:    New(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
		},
		{
			input: []byte(`"2007-11-11T17:38:12.000000001Z"`),
			dt:    New(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
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
		var dt DateTime
		err := json.Unmarshal(tc.input, &dt)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.dt, dt)
	}
}

func TestScan(t *testing.T) {
	tt := []struct {
		input interface{}
		dt    DateTime
		err   error
	}{
		{
			input: []byte("2007-11-11T17:38:12.432Z"),
			dt:    New(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
		},
		{
			input: []byte("invalid"),
			err:   errors.New("found i, expected number"),
		},
		{
			input: "2007-11-11T17:38:12.000000001Z",
			dt:    New(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
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
		var dt DateTime
		err := dt.Scan(tc.input)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.dt, dt)
	}
}

func TestValue(t *testing.T) {
	tt := []struct {
		dt     DateTime
		output driver.Value
	}{
		{
			dt:     New(2007, time.November, 11, 17, 38, 12, 432000000, time.UTC),
			output: "2007-11-11T17:38:12.432Z",
		},
		{
			dt:     New(2007, time.November, 11, 17, 38, 12, 1, time.UTC),
			output: "2007-11-11T17:38:12.000000001Z",
		},
	}

	for _, tc := range tt {
		val, err := tc.dt.Value()
		assert.Nil(t, err)
		assert.Equal(t, tc.output, val)
	}
}
