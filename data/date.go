package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date represents date with a special case for 'Present'
type Date struct {
	Time    time.Time
	YM      bool
	Present bool
}

func (d Date) String() string {
	switch {
	case d.Present:
		return "present"
	case d.Time.IsZero():
		return ""
	case d.YM:
		return d.Time.Format("2006-01")
	default:
		return d.Time.Format("2006-01-02")
	}
}

// UnmarshalJSON decodes a JSON string or null data into a Date
// If the string is "present" then it will refer to the present date
// Otherwise, it will try to parse the string from the form YYYY-MM
func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "": // Unmarhsal([]byte("null"),&s) is a no-op and leaves s empty
		return nil
	case "present":
		d.Present = true
	default:
		ym := true
		t, err := time.Parse("2006-01", s)
		if err != nil {
			t, err = time.Parse("2006-01-02", s)
			ym = false
		}
		if err != nil {
			return err
		}
		// assign after error check rather than during the time.Parse line so there are no parse-error side-effects
		d.Time = t
		d.YM = ym
	}
	return nil
}

// MarshalJSON converts the Date into a JSON string representation
func (d Date) MarshalJSON() ([]byte, error) {
	switch {
	case d.Present:
		return json.Marshal("present")
	case !d.Time.IsZero():
		if d.YM {
			return json.Marshal(d.Time.Format("2006-01"))
		}
		return json.Marshal(d.Time.Format("2006-01-02"))
	}
	return []byte("null"), nil
}

// Present returns the special "Present" date
func Present() Date {
	return Date{Present: true}
}

// YM parses a string in YYYY-MM form and returns a Date
// If the string given does not parse, it panics
func YM(str string) Date {
	t, err := time.Parse("2006-01", str)
	if err != nil {
		panic(fmt.Errorf("error parsing '%s': %v", str, err))
	}
	return Date{Time: t, YM: true}
}

// YMD returns a time.Time of midnight UTC on the given day of the year
func YMD(str string) Date {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		panic(fmt.Errorf("error parsing '%s': %v", str, err))
	}
	return Date{Time: t}
}
