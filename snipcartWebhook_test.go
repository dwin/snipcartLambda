package main

import "testing"

func TestOfSpacemap(t *testing.T) {
	withWhitespace := "435 KMN"
	woWhitespace := "435KMN"
	result := spacemap(withWhitespace)
	if result == "435KMN" {
		return
	}

	t.Fatalf("Expected %s but got %s", woWhitespace, result)
}
