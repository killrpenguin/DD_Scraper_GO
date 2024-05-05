package main

import "fmt"



func main() {
    fmt.Println("Hello, world.")
	body := DnRequest("https://feeds.megaphone.fm/darknetdiaries")
	pcData := DnParseXML(body)
	fmt.Printf("Podcasts Found: %d",len(pcData.DnPodcasts))
}
