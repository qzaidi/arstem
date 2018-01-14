package arstem

import (
	"testing"
)

func TestArStem(t *testing.T) {
	input := []string{
		"صِف",
	}

	expect := []string{
		"صف",
	}

	for idx, w := range input {
		result := StemASRI(w)
		if result != expect[idx] {
			t.Error("failed for ", input, result, "!=", expect[idx])
		}
	}

}
