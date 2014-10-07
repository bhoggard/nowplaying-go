package main

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"log"
)

type Piece struct {
	Title    string
	Composer string
}

type SecondInversionFeed struct {
	XMLName xml.Name             `xml:nexgen_audio_export`
	Audio   SecondInversionAudio `xml:"audio"`
}

type SecondInversionAudio struct {
	Title    string `xml:"title"`
	Composer string `xml:"composer"`
}

func translateSecondInversion(data []byte) Piece {
	var feed SecondInversionFeed

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(&feed)

	if err != nil {
		log.Fatal("decoder error:", err)
	}

	return Piece{feed.Audio.Title, feed.Audio.Composer}
}
