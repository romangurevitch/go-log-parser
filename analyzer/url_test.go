package analyzer

import (
	"github.com/romangurevitch/go-log-parser/httplog"
	"reflect"
	"testing"
)

func TestUrlAnalyzer_AddEntry(t *testing.T) {
	entries := []struct {
		log   *httplog.Log
		count int
	}{
		{&httplog.Log{Url: []byte("10")}, 10},
		{&httplog.Log{Url: []byte("5")}, 5},
		{&httplog.Log{Url: []byte("15")}, 15},
		{&httplog.Log{Url: []byte("2")}, 2},
		{&httplog.Log{Url: []byte("1")}, 1},
		{&httplog.Log{Url: []byte("30")}, 30},
	}
	analyzer := newUrlAnalyzer()

	for _, entry := range entries {
		for i := 0; i < entry.count; i++ {
			analyzer.AddEntry(entry.log)
		}
	}

	expected := []string{"30", "15", "10"}
	actual := getValues(analyzer.topVisitedUrls)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Expected: ", expected, "Got: ", actual)
	}
}
