package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/smaxwellstewart/articlescrape/article"
)

var ClusterThreshold = flag.Float64("threshold", 0.9, "cluster similarity score threshold")

const (
	// NOTE: these would need to be fine tuned
	WithinDayScore  = 0.0
	WithinWeekScore = 0.0
	// Time constants
	DayInSeconds  = 60 * 60 * 24
	WeekInSeconds = 60 * 60 * 24 * 7
)

// Article models an article plus it's similarity scores
type Article struct {
	article.Article
	Index int
}

// Similarities is a matrix of similarity scores
type Similarities [][]float64

// NOTE: not currently in use but could be used to implement full decision tree
func InCluster(a Article, b Article, similarities [][]float64) bool {
	// NOTE: time score would probably be better as a continuous function
	delta := a.Time - b.Time
	var timeScore float64
	if delta < DayInSeconds {
		timeScore = WithinDayScore
	} else if delta < WeekInSeconds {
		timeScore = WithinWeekScore
	}
	similarityScore := similarities[a.Index][b.Index]
	return (timeScore + similarityScore) > *ClusterThreshold
}

func main() {
	flag.Parse()

	file, err := ioutil.ReadFile("../similarities.json")
	if err != nil {
		log.Fatal("Could not load ../similarities.json:", err)
	}

	var s Similarities
	json.Unmarshal(file, &s)

	file, err = ioutil.ReadFile("../sample.json")
	if err != nil {
		log.Fatal("Could not load ../sample.json:", err)
	}

	var articles []Article
	json.Unmarshal(file, &articles)

	_clusters := make([][]Article, len(articles))
	for n := 0; n < len(_clusters); n++ {
		_clusters[n] = make([]Article, len(articles))
	}

	for i := 0; i < len(articles); i++ {
		for j := i + 1; j < len(articles); j++ {

			if s[i][j] > *ClusterThreshold {
				fmt.Printf("Similar articles found: %s | %s\n", articles[i].Title, articles[j].Title)
				fmt.Printf("(%d, %d) -> %f\n", i, j, s[i][j])
			}
		}
	}
}
