package txtstore

import "testing"

func TestStore(t *testing.T) {
	tests := []struct {
		s  string
		id string
	}{
		{s: "simple string 1"},
		{s: "simple string 2"},
		{s: "simple string 3"},
	}

	for i, tt := range tests {
		tests[i].id = Write(tt.s)
	}
	for _, tt := range tests {
		got := Read(tt.id)
		if got != tt.s {
			t.Errorf("got: %s expect: %s", got, tt.s)
		}
	}
}
