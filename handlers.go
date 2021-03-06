package main

import (
	"encoding/json"
	"net/http"
	"path"
)

const q2Url = "http://www.wqxr.org/api/whats_on/q2/2/"
const counterstreamUrl = "http://www.live365.com/pls/front?handler=playlist&cmd=view&viewType=xml&handle=amcenter&maxEntries=1"
const secondInversionUrl = "http://filesource.abacast.com/king/TRE/inversion2.xml"
const yleUrl = "http://yle.fi/radiomanint/LiveXML/r17/item(0).xml"

func index(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "index.html")
	http.ServeFile(w, r, fp)
}

func counterstreamJSON(w http.ResponseWriter, r *http.Request) {
	piece := counterstream()
	js, err := json.Marshal(piece)
	checkErr(err, "json.Marshal failed")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func q2JSON(w http.ResponseWriter, r *http.Request) {
	piece := q2()
	js, err := json.Marshal(piece)
	checkErr(err, "json.Marshal failed")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func secondInversionJSON(w http.ResponseWriter, r *http.Request) {
	piece := secondInversion()
	js, err := json.Marshal(piece)
	checkErr(err, "json.Marshal failed")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func yleJSON(w http.ResponseWriter, r *http.Request) {
	piece := yle()
	js, err := json.Marshal(piece)
	checkErr(err, "json.Marshal failed")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
