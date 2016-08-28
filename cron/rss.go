package cron

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type RSS struct {
	XMLName      xml.Name `xml:"rss"`
	Title        string   `xml:"channel>title"`
	Descriptioin string   `xml:"channel>descriptioin"`
	Items        []Item   `xml:"channel>item"`
}

type Item struct {
	Title        string `xml:"title"`
	Link         string `xml:"link"`
	PubDate      string `xml:"pubDate"` // Sat, 20 Aug 2016 16:08:14 GMT
	Descriptioin string `xml:"description"`
}

func Parse(url string) ([]Item, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rss := RSS{}
	err = xml.Unmarshal(content, &rss)

	return rss.Items, err
}
