package main

import "testing"

func TestDouble(t *testing.T) {

	given := 2
	want := 4

	result := given
	double(&result)

	if result != want {
		t.Errorf("Given %d, Want %d. But %d", given, want, result)
	}
}

func TestAppGreeting(t *testing.T) {

	given := "Bob"
	want := "Hi, Bob"

	result := given
	appendGreeting(&result)
	if result != want {
		t.Errorf("Given %s, Want %s. But %s", given, want, result)
	}

}
