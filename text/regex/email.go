package regex

import (
	"regexp"
)

const PatternEmail = `([a-zA-Z0-9._-]+)@([a-zA-Z0-9_-]+)(\.[a-zA-Z0-9.]*[a-zA-Z0-9]+)`


func Email(in string) (match string) {
	re := regexp.MustCompile(PatternEmail)
	match = re.FindString(in)
	return match
}

func EmailBatch(in string) (matches []string) {
	re := regexp.MustCompile(PatternEmail)
	matches = re.FindAllString(in, -1)
	return matches
}

func NameAndComOfEmail(in string) (match []string) {
	re := regexp.MustCompile(PatternEmail)
	match = re.FindStringSubmatch(in)
	return match
}

func NameAndComOfEmailBatch(in string) (matches [][]string) {
	re := regexp.MustCompile(PatternEmail)
	matches = re.FindAllStringSubmatch(in, -1)
	return matches
}
