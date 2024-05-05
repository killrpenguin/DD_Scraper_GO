package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)


type XMLItem struct {
	XMLName xml.Name `xml:"item"`
	Title string `xml:"itunes:title"`
	Episode int `xml:"itunes:episode"`
	Season int `xml:"itunes:season"`
	Url string `xml:"url,attr"`
}

type DnPodcasts struct {
	DnPodcasts []XMLItem
}

func DnRequest (link string) (page []byte ) {
	resp, resp_err := http.Get(link)
	if resp_err != nil && resp.StatusCode < 299 { 
		log.Println(resp_err)
		defer resp.Body.Close()
	}
	byteValue, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return byteValue
}

func DnParseXML (bval []byte) DnPodcasts {
	var pcData DnPodcasts
	xml.Unmarshal(bval, &pcData)
	for i:=0; i < len(pcData.DnPodcasts); i++ {
		fmt.Printf("Title: ", pcData.DnPodcasts[i].Title)
	}
	return pcData
}
