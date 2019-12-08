package httplog

import (
	"reflect"
	"testing"
)

func TestLog_Extract(t *testing.T) {
	tests := []struct {
		name     string
		expected *Log
		line     []byte
	}{
		{
			name:     "simple",
			expected: &Log{nil, []byte("50.112.00.11"), []byte("11/Jul/2018:17:33:01 +0200"), []byte("GET"), []byte("/asset.css")},
			line:     []byte(`50.112.00.11 - admin [11/Jul/2018:17:33:01 +0200] "GET /asset.css HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"`),
		},
		{
			name:     "simple2",
			expected: &Log{nil, []byte("177.71.128.21"), []byte("10/Jul/2018:22:21:28 +0200"), []byte("GET"), []byte("/intranet-analytics/")},
			line:     []byte(`177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := &Log{}
			_, err := actual.Extract(tt.line)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tt.expected.IPAddress, actual.IPAddress) {
				t.Error("Expected IP:", string(tt.expected.IPAddress), "Actual IP:", string(actual.IPAddress))
			}
			if !reflect.DeepEqual(tt.expected.Date, actual.Date) {
				t.Error("Expected Date:", string(tt.expected.Date), "Actual Date:", string(actual.Date))
			}
			if !reflect.DeepEqual(tt.expected.Method, actual.Method) {
				t.Error("Expected Method:", string(tt.expected.Method), "Actual Method:", string(actual.Method))
			}
			if !reflect.DeepEqual(tt.expected.Url, actual.Url) {
				t.Error("Expected Url:", string(tt.expected.Url), "Actual Url:", string(actual.Url))
			}
		})
	}
}
