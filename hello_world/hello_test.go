package main

import "testing"


func asserCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}


func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		asserCorrectMessage(t, got, want)
	
	})

	t.Run("say Hello world when and empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		asserCorrectMessage(t, got, want)
	})


	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		asserCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		asserCorrectMessage(t, got, want)
	})
}
