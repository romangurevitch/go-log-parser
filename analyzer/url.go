package analyzer

import (
	"fmt"
	"github.com/romangurevitch/go-log-parser/httplog"
	"strings"
)

const maxTopUrls = 3

func newUrlAnalyzer() *urlAnalyzer {
	return &urlAnalyzer{
		urls:           map[string]int{},
		topVisitedUrls: maxValues{},
	}
}

type urlAnalyzer struct {
	urls           map[string]int
	topVisitedUrls maxValues
}

func (a *urlAnalyzer) AddEntry(entry *httplog.Log) {
	url := string(entry.Url)
	if _, ok := a.urls[url]; ok {
		a.urls[url]++
	} else {
		a.urls[url] = 1
	}

	a.updateTopUrls(url)
}

func (a *urlAnalyzer) Report() string {
	return fmt.Sprint("Top 3 most visited URL addresses: ", strings.Join(getValues(a.topVisitedUrls), ", "))
}

func (a *urlAnalyzer) updateTopUrls(url string) {
	a.topVisitedUrls.addValue(CountValue{a.urls[url], url}, maxTopUrls)
}
