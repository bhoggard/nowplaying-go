package main

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Piece struct {
	Title    string `json:"title"`
	Composer string `json:"composer"`
}

type SecondInversionFeed struct {
	XMLName xml.Name             `xml:nexgen_audio_export`
	Audio   SecondInversionAudio `xml:"audio"`
}

type SecondInversionAudio struct {
	Title    string `xml:"title"`
	Composer string `xml:"composer"`
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func translateSecondInversion(data []byte) Piece {
	var feed SecondInversionFeed
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(&feed)
	checkErr(err, "translateSecondInversion decode error")
	return Piece{feed.Audio.Title, feed.Audio.Composer}
}

func secondInversion() Piece {
	resp, err := http.Get(secondInversionUrl)
	checkErr(err, "failed get of secondInversion feed")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "failed to read respsone body")
	return translateSecondInversion(body)
}
