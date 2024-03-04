package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
