package regex

import (
	"github.com/zhang1career/module/datastruct/bytes"
	"testing"
)

var hypertext       = []byte(`<a id="1" href="http://www.zhenai.com/zhenghun/yushu1" class="">玉树</a>`)
var verityHypertext = [][]byte{
	[]byte(`<a id="1" href="http://www.zhenai.com/zhenghun/yushu1" class="">玉树</a>`),
	[]byte(`http://www.zhenai.com/zhenghun/yushu1`),
	[]byte(`玉树`),
}
func TestHypertext(t *testing.T) {
	match := Hypertext(hypertext)
	for i, ma := range match {
		if bytes.Equal(&ma, &(verityHypertext[i])) {
			t.Logf("OK, got: %s\n", ma)
		} else {
			t.Errorf("WRONG, expected: %s, got: %s\n", verityHypertext[i], ma)
		}
	}
}


var hypertextBatch  = []byte(`
	<a id="1" href="http://www.zhenai.com/zhenghun/anhui" class="">安徽</a>
	<a href="http://www.zhenai.com/zhenghun/baotou" class="">包头</a>
`)
var verityHypertextBatch = [][][]byte{
	[][]byte{
		[]byte(`<a id="1" href="http://www.zhenai.com/zhenghun/anhui" class="">安徽</a>`),
		[]byte(`http://www.zhenai.com/zhenghun/anhui`),
		[]byte(`安徽`),
	},
	[][]byte{
		[]byte(`<a href="http://www.zhenai.com/zhenghun/baotou" class="">包头</a>`),
		[]byte(`http://www.zhenai.com/zhenghun/baotou`),
		[]byte(`包头`),
	},
}
func TestHypertextBatch(t *testing.T) {
	matches := HypertextBatch(hypertextBatch)
	for i, match := range matches {
		t.Logf("test %s...\n", verityHypertextBatch[i][0])
		for j, ma := range match {
			if bytes.Equal(&ma, &(verityHypertextBatch[i][j])) {
				t.Logf("OK, got: %s\n", ma)
			} else {
				t.Errorf("WRONG, expected: %s, got: %s\n", verityHypertextBatch[i][j], ma)
			}
		}
	}
}

var PatternFilter = `<a[^{href}]*href="([^"]*)"[^>]*>([^<]+)</a>`
func TestFilterBatch(t *testing.T) {
	matches := FilterBatch(hypertextBatch, &PatternFilter)
	for i, match := range matches {
		t.Logf("test %s...\n", verityHypertextBatch[i][0])
		for j, ma := range match {
			if bytes.Equal(&ma, &(verityHypertextBatch[i][j])) {
				t.Logf("OK, got: %s\n", ma)
			} else {
				t.Errorf("WRONG, expected: %s, got: %s\n", verityHypertextBatch[i][j], ma)
			}
		}
	}
}