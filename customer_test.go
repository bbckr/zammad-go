package zammad

import (
	"testing"
)

func TestCustomerID_String(t *testing.T) {
	test := []struct {
		name  string
		value interface{}
		want  string
	}{
		{name: "should handle int", value: 123, want: "123"},
		{name: "should handle nil", value: nil, want: ""},
		{name: "should handle string", value: "guess:{hello@example.com}", want: "guess:{hello@example.com}"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerID{Value: tt.value}
			if got.String() != tt.want {
				t.Errorf("expected %s, got %s", tt.want, got.String())
			}
		})
	}
}

func TestCustomerID_Int(t *testing.T) {
	test := []struct {
		name  string
		value interface{}
		want  int
	}{
		{name: "should convert int", value: 123, want: 123},
		{name: "should handle nil", value: nil, want: -1},
		{name: "should handle string", value: "123", want: 123},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerID{Value: tt.value}
			if got.Int() != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got.Int())
			}
		})
	}

}

func TestCustomerID_MarshalJSON(t *testing.T) {
	test := []struct {
		name  string
		value interface{}
		want  string
	}{
		{name: "should marshal int", value: 123, want: "123"},
		{name: "should marshal nil", value: nil, want: "null"},
		{name: "should marshal string", value: "guess:{hello@example.com}", want: "\"guess:{hello@example.com}\""},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerID{Value: tt.value}

			b, err := got.MarshalJSON()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if string(b) != tt.want {
				t.Errorf("expected %s, got %s", tt.want, string(b))
			}
		})
	}

}

func TestCustomerID_UnmarshalJSON(t *testing.T) {
	test := []struct {
		name  string
		input string
		want  interface{}
	}{
		{name: "should unmarshal int", input: "123", want: float64(123)},
		{name: "should unmarshal null", input: "null", want: nil},
		{name: "should unmarshal string", input: "\"guess:{hello@example.com}\"", want: "guess:{hello@example.com}"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			var got CustomerID

			err := got.UnmarshalJSON([]byte(tt.input))
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got.Value != tt.want {
				t.Errorf("expected %v (%T), got %v (%T)", tt.want, tt.want, got.Value, got.Value)
			}
		})
	}
}
