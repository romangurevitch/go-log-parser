package analyzer

import (
	"github.com/romangurevitch/go-log-parser/httplog"
	"strings"
)

func New() *compositeAnalyzer {
	return &compositeAnalyzer{
		[]Analyzer{
			newIpAnalyzer(),
			newUrlAnalyzer(),
		},
	}
}

type Analyzer interface {
	AddEntry(entry *httplog.Log)
	Report() string
}

type compositeAnalyzer struct {
	analyzers []Analyzer
}

func (c *compositeAnalyzer) AddEntry(entry *httplog.Log) {
	for _, v := range c.analyzers {
		v.AddEntry(entry)
	}
}

func (c *compositeAnalyzer) Report() string {
	var reports []string
	for _, v := range c.analyzers {
		reports = append(reports, v.Report())
	}
	return strings.Join(reports, "\n")
}
