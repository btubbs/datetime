package datetime

import (
	"bytes"
	"time"
)

func Parse(s string) (time.Time, error) {
	p := newParser(bytes.NewBuffer([]byte(s)))
	return p.parse()
}
