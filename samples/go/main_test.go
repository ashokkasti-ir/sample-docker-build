package main

import "testing"

func TestSmoke(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("math smoke test failed")
	}
}
