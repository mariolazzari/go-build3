package main

import "encoding/xml"

type RSS struct {
	XMLName xml.Name
	Channel *Channel
}

type Channel struct {
	Title    string
	ItemList []Item
}

type News struct {
	Headline     string
	HeadlineLink string
}

type Item struct {
	Title     string
	Link      string
	Traffic   string
	NewsItems []News
}

func main() {
	readGoogleTrends()
}

func readGoogleTrends() {
	getGoogleTrends()
}

func getGoogleTrends() {

}
