package coyote

import "testing"

func TestConnect(t *testing.T) {
	err := Connect("localhost:6379")
	if err != nil {
		t.Error("Error connecting.")
	}
}
