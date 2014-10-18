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

type CounterstreamFeed struct {
	XMLName xml.Name           `xml:"Playlist"`
	Entry   CounterstreamEntry `xml:"PlaylistEntry"`
}

type CounterstreamEntry struct {
	Title  string
	Artist string
}

type SecondInversionFeed struct {
	XMLName xml.Name             `xml:"nexgen_audio_export"`
	Audio   SecondInversionAudio `xml:"audio"`
}

type SecondInversionAudio struct {
	Title    string `xml:"title"`
	Composer string `xml:"composer"`
}

type YleFeed struct {
	XMLName xml.Name `xml:"RMPADEXPORT"`
	Item    YleItem  `xml:"ITEM"`
}

type YleItem struct {
	Composer    string         `xml:"COMPOSER,attr"`
	PublishData YlePublishData `xml:"PUBLISH-DATA"`
}

type YlePublishData struct {
	Title string `xml:"TEXT"`
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func translateCounterstream(data []byte) Piece {
	var feed CounterstreamFeed
	err := xml.Unmarshal(data, &feed)
	checkErr(err, "translateCounterstream unmarshal error")
	return Piece{feed.Entry.Title, feed.Entry.Artist}
}

func counterstream() Piece {
	resp, err := http.Get(counterstreamUrl)
	checkErr(err, "failed get of counterstream feed")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "failed to read response body")
	return translateCounterstream(body)
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
	checkErr(err, "failed to read response body")
	return translateSecondInversion(body)
}

func translateYle(data []byte) Piece {
	var feed YleFeed
	err := xml.Unmarshal(data, &feed)
	checkErr(err, "translateYle unmarshal error")
	return Piece{feed.Item.PublishData.Title, feed.Item.Composer}
}

func yle() Piece {
	resp, err := http.Get(yleUrl)
	checkErr(err, "failed get of YLE feed")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "failed to read response body")
	return translateYle(body)
}
