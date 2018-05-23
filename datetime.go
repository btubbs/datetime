package datetime

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

type DateTime time.Time

func New(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) DateTime {
	return DateTime(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

func Parse(input string) (DateTime, error) {
	return parseBytes([]byte(input))
}

// TODO:
// sql value and scan methods

// String returns the DateTime's RFC3339Nano representation.
func (d DateTime) String() string {
	t := time.Time(d)
	return t.Format(time.RFC3339Nano)
}

const doubleQuote byte = 34

// UnmarshalJSON implements the JSON Unmarshaler interface, allowing datetime.DateTime struct fields
// to be read from JSON string fields.
func (d *DateTime) UnmarshalJSON(data []byte) error {
	// strip enclosing quotes
	if data[0] != doubleQuote || data[len(data)-1] != doubleQuote {
		return fmt.Errorf("%s does not begin and end with double quotes", data)
	}
	trimmed := data[1 : len(data)-1]

	t, err := parseBytes(trimmed)
	if err != nil {
		return err
	}
	*d = t
	return nil
}

// Scan implements the sql Scanner interface, allowing datetime.DateTime fields to be read from
// database columns.
func (d *DateTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		dt, err := parseBytes(v)
		if err != nil {
			return err
		}
		*d = dt
	case string:
		dt, err := parseBytes([]byte(v))
		if err != nil {
			return err
		}
		*d = dt
	default:
		return fmt.Errorf("can only scan string and []byte, not %v", reflect.TypeOf(value))
	}
	return nil
}

// Value implements the sql Valuer interface, allowing datetime.DateTime fields to be saved to
// database columns.
func (d DateTime) Value() (driver.Value, error) {
	return d.String(), nil
}

func parseBytes(b []byte) (DateTime, error) {
	p := newParser(bytes.NewBuffer(b))
	t, err := p.parse()
	return DateTime(t), err
}
