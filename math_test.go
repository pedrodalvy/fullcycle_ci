package main

import "testing"

func TestSooma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Expected 30, got %d", total)
	}
}
