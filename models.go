package main

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/json"
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

type Q2Feed struct {
	Item Q2Item `json:"current_playlist_item"`
}

type Q2Item struct {
	Entry Q2Entry `json:"catalog_entry"`
}

type Q2Entry struct {
	Title    string
	Composer Q2Composer `json:"composer"`
}

type Q2Composer struct {
	Name string `json:"name"`
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

func translateQ2(data []byte) Piece {
	var feed Q2Feed
	err := json.Unmarshal(data, &feed)
	checkErr(err, "translateQ2 unmarshal error")
	return Piece{feed.Item.Entry.Title, feed.Item.Entry.Composer.Name}
}

func q2() Piece {
	resp, err := http.Get(q2Url)
	checkErr(err, "failed get of Q2 feed")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "failed to read response body")
	return translateQ2(body)
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
