package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSecondInversion(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/second-inversion.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	piece := translateSecondInversion(content)

	expected := "Violin Sonata No.1"
	if piece.Title != expected {
		t.Errorf("Title = %v, expected %v", piece.Title, expected)
	}

	expected = "Frederic Delius"
	if piece.Composer != expected {
		t.Errorf("Composer = %v, expected %v", piece.Composer, expected)
	}
}

func TestCounterstream(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/counterstream.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	piece := translateCounterstream(content)

	expected := "Serenade"
	if piece.Title != expected {
		t.Errorf("Title = %v, expected %v", piece.Title, expected)
	}

	expected = "Edward T. Cone"
	if piece.Composer != expected {
		t.Errorf("Composer = %v, expected %v", piece.Composer, expected)
	}
}
