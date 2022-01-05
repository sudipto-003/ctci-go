package stringRotation

import "testing"

func TestStringRotation(t *testing.T) {
	tests := []struct {
		input string
		rot   string
		valid bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"helloworld", "elloworldh", true},
		{"abcdefgh", "ghabcdef", true},
		{"abcdefgh", "cdefabgh", false},
		{"", "", false},
	}

	for i, tt := range tests {
		result := IsRotation(tt.input, tt.rot)
		if result != tt.valid {
			t.Errorf("Error in test: %d. String %s is not rotation of %s", i, tt.input, tt.rot)
		}
	}
}
