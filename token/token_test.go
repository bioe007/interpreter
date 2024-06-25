package token

import "testing"

func TestTokenLulz(t *testing.T) {
	thisToken := Token{
		"FOO",
		"BAR",
	}
	if thisToken.Type != "FOO" {
		t.Fatal("you suck")
	}
}
