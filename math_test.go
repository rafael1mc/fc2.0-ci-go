package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Invalid sum result: Result %d. Expected: %d", total, 30)
	}
}
