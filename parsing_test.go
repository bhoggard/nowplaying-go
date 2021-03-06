package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

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

func TestQ2(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/q2.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	piece := translateQ2(content)

	expected := "Turangalila-symphonie"
	if piece.Title != expected {
		t.Errorf("Title = %v, expected %v", piece.Title, expected)
	}

	expected = "Olivier Messiaen"
	if piece.Composer != expected {
		t.Errorf("Composer = %v, expected %v", piece.Composer, expected)
	}
}

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

func TestYle(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/yle.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	piece := translateYle(content)

	expected := "Godard: Pianokonsertto n:o 1 a-molli. (Howard Shelley ja Tasmanian SO)."
	if piece.Title != expected {
		t.Errorf("Title = %v, expected %v", piece.Title, expected)
	}

	expected = "Godard, Benjamin [1849-1895]"
	if piece.Composer != expected {
		t.Errorf("Composer = %v, expected %v", piece.Composer, expected)
	}
}
