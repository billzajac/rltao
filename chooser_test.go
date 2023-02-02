package rltao

import "testing"

func TestGenerateId(t *testing.T) {
	for i := 1; i <= 5000; i++ {
		got := GenerateId()
		if got > 81 || got < 1 {
			t.Errorf("GenerateId returned a value out of range: %q", got)
		}
	}
}
