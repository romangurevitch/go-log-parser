package analyzer

import (
	"github.com/romangurevitch/go-log-parser/httplog"
	"reflect"
	"testing"
)

func Test_ipAnalyzer_AddEntry(t *testing.T) {
	entries := []struct {
		log   *httplog.Log
		count int
	}{
		{&httplog.Log{IPAddress: []byte("10")}, 10},
		{&httplog.Log{IPAddress: []byte("5")}, 5},
		{&httplog.Log{IPAddress: []byte("15")}, 15},
		{&httplog.Log{IPAddress: []byte("2")}, 2},
		{&httplog.Log{IPAddress: []byte("1")}, 1},
		{&httplog.Log{IPAddress: []byte("30")}, 30},
	}
	analyzer := newIpAnalyzer()

	for _, entry := range entries {
		for i := 0; i < entry.count; i++ {
			analyzer.AddEntry(entry.log)
		}
	}

	expected := []string{"30", "15", "10"}
	actual := getValues(analyzer.topActiveIps)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Expected: ", expected, "Got: ", actual)
	}

	uniqueIps := len(analyzer.ips)
	expectedIps := 6
	if uniqueIps != expectedIps {
		t.Fatal("Expected: ", expectedIps, "Got: ", uniqueIps)
	}
}
