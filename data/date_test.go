package data

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestDateMarshalJSON(t *testing.T) {
	tests := []struct {
		input    Date
		expected []byte
	}{
		{Date{}, []byte(`null`)},
		{Present(), []byte(`"present"`)},
		{YM("2011-11"), []byte(`"2011-11"`)},
		{YMD("2011-11-11"), []byte(`"2011-11-11"`)},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			bs, err := test.input.MarshalJSON()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !bytes.Equal(bs, test.expected) {
				t.Fatalf("expected '%s' but got '%s'", string(test.expected), string(bs))
			}
		})
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    []byte
		expected Date
	}{
		{[]byte(`null`), Date{}},
		{[]byte(`"present"`), Present()},
		{[]byte(`"2011-11"`), YM("2011-11")},
		{[]byte(`"2011-11-11"`), YMD("2011-11-11")},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var date Date
			if err := (&date).UnmarshalJSON(test.input); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(date, test.expected) {
				t.Fatalf("expected '%#v' but got '%#v'", test.expected, date)
			}
		})
	}
}
