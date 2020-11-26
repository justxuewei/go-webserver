package util

import "testing"

func TestTrimFirstRune(t *testing.T) {
	s := "/public"
	s1 := TrimFirstRune(s)
	if s1 != "public" {
		t.Error("TrimFirstRune() Error.")
	}
}
