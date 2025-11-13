package catalogue

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FlexibleFloat is a custom type that can unmarshal from both string and float64.
// This handles the API inconsistency where numeric values may be returned as strings.
type FlexibleFloat float64

// UnmarshalJSON implements json.Unmarshaler for FlexibleFloat.
// It accepts both numeric values and string representations of numbers.
func (f *FlexibleFloat) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as float64 first
	var numVal float64
	if err := json.Unmarshal(data, &numVal); err == nil {
		*f = FlexibleFloat(numVal)
		return nil
	}

	// Try unmarshaling as string
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return fmt.Errorf("failed to unmarshal as number or string: %w", err)
	}

	// Parse string as float
	parsed, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse string %q as float: %w", strVal, err)
	}

	*f = FlexibleFloat(parsed)
	return nil
}

// MarshalJSON implements json.Marshaler for FlexibleFloat.
// It always marshals as a numeric value.
func (f FlexibleFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

// Float64 returns the float64 value of the FlexibleFloat.
func (f FlexibleFloat) Float64() float64 {
	return float64(f)
}

// FlexibleDate is a custom type that can unmarshal dates in multiple formats.
// This handles the API inconsistency where dates may be returned in different formats.
type FlexibleDate struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler for FlexibleDate.
// It tries multiple common date formats including RFC3339 and YYYY-MM-DD.
func (d *FlexibleDate) UnmarshalJSON(data []byte) error {
	// Remove quotes
	s := strings.Trim(string(data), "\"")

	// Handle empty/null values
	if s == "" || s == "null" {
		return nil
	}

	// Try common date/time formats in order of likelihood
	formats := []string{
		"2006-01-02",                 // YYYY-MM-DD (most common for dates)
		time.RFC3339,                 // 2006-01-02T15:04:05Z07:00
		time.RFC3339Nano,             // 2006-01-02T15:04:05.999999999Z07:00
		"2006-01-02T15:04:05",        // YYYY-MM-DDTHH:MM:SS (no timezone)
		"2006-01-02 15:04:05",        // YYYY-MM-DD HH:MM:SS
		time.DateTime,                // 2006-01-02 15:04:05
		time.DateOnly,                // 2006-01-02
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			d.Time = t
			return nil
		}
	}

	return fmt.Errorf("unable to parse date %q: tried formats %v", s, formats)
}

// MarshalJSON implements json.Marshaler for FlexibleDate.
// It marshals as RFC3339 format for consistency.
func (d FlexibleDate) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(d.Time.Format(time.RFC3339))
}
