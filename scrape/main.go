package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SlyMarbo/rss"
	goose "github.com/advancedlogic/GoOse"
	"github.com/smaxwellstewart/articlescrape/article"
)

var Feeds = []string{
	"http://www.streetupdates.com/feed/",
	"http://feeds.bizjournals.com/bizj_newyork",
	"http://www.thestreet.com/feeds/rss/index.xml",
	"http://blogs.barrons.com/stockstowatchtoday/feed/",
	"http://stockcharts.com/blogs/atom.xml",
	"http://www.moneycontrol.com/rss/buzzingstocks.xml",
	"http://www.ft.com/rss/markets/equities",
	"http://feeds.feedburner.com/marginalrevolution/feed",
	"http://www.nytimes.com/services/xml/rss/nyt/Dealbook.xml",
	"http://fortune.com/tag/street-sweep/feed/",
	"http://articlefeeds.nasdaq.com/nasdaq/categories?category=Stocks",
	"http://feeds.finance.yahoo.com/rss/2.0/headline?s=googl&region=US&lang=en-US",
	"http://gizmodo.com/rss",
	"http://www.newyorker.com/feed/business",
	"http://seekingalpha.com/tag/wall-st-breakfast.xml",
	"http://feeds.wsjonline.com/wsj/video/markets/feed",
	"http://fivethirtyeight.com/economics/feed/",
	"http://www.theglobeandmail.com/report-on-business/?service=rss",
	"https://news.google.com/news?cf=all&hl=en&pz=1&ned=us&output=rss",
	"https://en.wikinews.org/w/index.php?title=Special:NewsFeed&feed=rss&categories=Published&notcategories=No%20publish%7CArchived%7cAutoArchived%7cdisputed&namespace=0&count=15&ordermethod=categoryadd&stablepages=only",
	"http://feeds.feedburner.com/FP_TopStories",
	"https://www.project-syndicate.org/rss",
	"http://feeds.reuters.com/reuters/topNews",
	"http://www.cnbc.com/id/100003114/device/rss/rss.html",
	"http://feeds.feedburner.com/yahoo/DGcR",
	"http://feeds.reuters.com/reuters/businessNews",
}

func main() {
	// fetch rss
	items := make([]*rss.Item, 0)

	for _, url := range Feeds {
		fmt.Println("Fetching XML rss feed:", url)
		feed, err := rss.Fetch(url)
		if err != nil {
			fmt.Println("Error fetching rss feed", err)
			continue
		}
		items = append(items, feed.Items...)
	}

	g := goose.New()
	articles := []article.Article{}
	for _, item := range items {
		a, err := g.ExtractFromURL(item.Link)
		if err != nil {
			fmt.Println("Error extracting article", err)
			continue
		}
		if a.Title != "" {
			articles = append(articles, article.Article{Title: a.Title, Body: a.CleanedText, Description: a.MetaDescription, Keywords: a.MetaKeywords, Time: item.Date.Unix(), URL: item.Link})
		}
	}

	b, err := json.Marshal(articles)
	if err != nil {
		log.Fatal("Error marshalling json", err)
	}

	err = ioutil.WriteFile("../sample.json", b, 0644)
	if err != nil {
		log.Fatal("Error saving json to file", err)
	}
}
