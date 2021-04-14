package regex

import (
	"github.com/zhang1career/golab/log"
	"regexp"
)

const PatternHypertext = `<a[^{href}]*href="([^"]*)"[^>]*>([^<]+)</a>`

// param
//   in: html content
// return
//   match
//     [0]: whole hypertext
//     [1]: hyperlink
//     [2]: text
func Hypertext(in []byte) (match [][]byte) {
	re := regexp.MustCompile(PatternHypertext)
	match = re.FindSubmatch(in)
	return match
}

// param
//   html: html content
// return
//   matches
//     [i][0]: whole hypertext
//     [i][1]: hyperlink
//     [i][2]: text
func HypertextBatch(html []byte) (matches [][][]byte) {
	re := regexp.MustCompile(PatternHypertext)
	matches = re.FindAllSubmatch(html, -1)
	return matches
}

// param
//   html: html content
//   pattern: regex pattern
// return
//   matches
//     [i][0]: whole matched text
//     [i][1]: 1st captured text
//     [i][2]: 2nd captured text
//     [i][...]
func FilterBatch(html []byte, pattern *string) (matches [][][]byte) {
	re, err := regexp.Compile(*pattern)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	
	matches = re.FindAllSubmatch(html, -1)
	return matches
}