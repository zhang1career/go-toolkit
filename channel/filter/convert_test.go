package filter_test

import (
	"fmt"
	"github.com/zhang1career/lib/channel/filter"
	"github.com/zhang1career/lib/channel/source"
	"testing"
)

func TestToInt(t *testing.T) {
	done := make(chan interface{})
	
	for value := range filter.ToInt(done, source.RandSource(10)) {
		t.Logf("%d", value)
	}
}

func TestToString(t *testing.T) {
	done := make(chan interface{})
	
	var message string
	for token := range filter.ToString(done, filter.Take(done, source.Repeat(done, "I", "am."), 5)) {
		message += token
	}
	
	fmt.Printf("message: %s...", message)
}
