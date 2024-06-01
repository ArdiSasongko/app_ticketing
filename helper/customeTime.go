package helper

import (
	"fmt"
	"time"
)

const customTimeLayout = "2006-01-02T15:04:05Z07:00"

// CustomTime wraps time.Time for custom parsing
type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses the JSON string into a CustomTime object
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	// Remove quotes
	s = s[1 : len(s)-1]
	ct.Time, err = time.Parse(customTimeLayout, s)
	return
}

// MarshalJSON converts the CustomTime object to a JSON string
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Time.Format(customTimeLayout))), nil
}
