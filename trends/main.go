package main

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
