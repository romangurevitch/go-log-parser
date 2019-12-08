package analyzer

import (
	"fmt"
	"github.com/romangurevitch/go-log-parser/httplog"
	"strings"
)

const top = 3

func newIpAnalyzer() *ipAnalyzer {
	return &ipAnalyzer{
		ips:          map[string]int{},
		topActiveIps: maxValues{},
	}
}

type ipAnalyzer struct {
	ips          map[string]int
	topActiveIps maxValues
}

func (a *ipAnalyzer) AddEntry(entry *httplog.Log) {
	ip := string(entry.IPAddress)
	if _, ok := a.ips[ip]; ok {
		a.ips[ip]++
	} else {
		a.ips[ip] = 1
	}

	a.updateTopIps(ip)
}

func (a *ipAnalyzer) Report() string {
	var reports []string
	reports = append(reports, fmt.Sprint("Number of unique IP addresses: ", len(a.ips)))
	reports = append(reports, fmt.Sprint("Top 3 most active IP addresses: ", strings.Join(getValues(a.topActiveIps), ", ")))
	return strings.Join(reports, "\n")
}

func (a *ipAnalyzer) updateTopIps(ip string) {
	a.topActiveIps.addValue(CountValue{a.ips[ip], ip}, top)
}
