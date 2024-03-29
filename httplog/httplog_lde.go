// Code generated by ldetool --package httplog httplog.lde. DO NOT EDIT.

package httplog

import (
	"bytes"
)

// Using https://sirkon.github.io/ldetool/ for log parsing
// See detailed rules syntax at: https://github.com/sirkon/ldetool/blob/master/TOOL_RULES.md
type Log struct {
	Rest      []byte
	IPAddress []byte
	Date      []byte
	Method    []byte
	Url       []byte
}

// Extract ...
func (p *Log) Extract(line []byte) (bool, error) {
	p.Rest = line
	var pos int

	// Take until ' ' as IPAddress(string)
	pos = bytes.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		p.IPAddress = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '[' and then pass it
	pos = bytes.IndexByte(p.Rest, '[')
	if pos >= 0 {
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until ']' as Date(string)
	pos = bytes.IndexByte(p.Rest, ']')
	if pos >= 0 {
		p.Date = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '"' and then pass it
	pos = bytes.IndexByte(p.Rest, '"')
	if pos >= 0 {
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until ' ' as Method(string)
	pos = bytes.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		p.Method = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until ' ' as Url(string)
	pos = bytes.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		p.Url = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	return true, nil
}
