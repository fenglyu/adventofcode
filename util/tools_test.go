package util

import "testing"

func TestMod(t *testing.T) {
	if Mod(-1, 5) != 4 {
		t.Fatal("unexpected result")
	}
}
