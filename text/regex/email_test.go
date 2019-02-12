package regex

import (
	"testing"
)


var email       = "My email is rongjin.zh@gmail.com"
var verifyEmail = "rongjin.zh@gmail.com"
func TestEmail(t *testing.T) {
	match := Email(email)
	if match == verifyEmail {
		t.Logf("OK, email: %s\n", match)
	} else {
		t.Errorf("WRONG, expected: %s, got: %s\n", verifyEmail, match)
	}
}


var verifyNameAndComOfEmail = []string{
	"rongjin.zh@gmail.com", "rongjin.zh", "gmail", ".com",
}
func TestNameAndComOfEmail(t *testing.T) {
	match := NameAndComOfEmail(email)
	for i, ma := range match {
		if ma == verifyNameAndComOfEmail[i] {
			t.Logf("OK, got: %s\n", ma)
		} else {
			t.Errorf("WRONG, expected: %s, got: %s\n", verifyNameAndComOfEmail[i], ma)
		}
	}
}


var emailBatch          = `
My email is rongjin.zh@gmail.com
Are your email zhang_career@github-hehe.com
He get a @hotmail that is rich-billion@hotmail.com.cn
`
var verifyEmailBatch    = []string{
	"rongjin.zh@gmail.com",
	"zhang_career@github-hehe.com",
	"rich-billion@hotmail.com.cn",
}
func TestEmailBatch(t *testing.T) {
	matches := EmailBatch(emailBatch)
	for i, ma := range matches {
		if ma == verifyEmailBatch[i] {
			t.Logf("OK, email: %s\n", ma)
		} else {
			t.Errorf("WRONG, expected: %s, got: %s\n", verifyEmailBatch[i], ma)
		}
	}
}

var verifyNameAndComOfEmailBatch = [][]string{
	{"rongjin.zh@gmail.com", "rongjin.zh", "gmail", ".com"},
	{"zhang_career@github-hehe.com", "zhang_career", "github-hehe", ".com"},
	{"rich-billion@hotmail.com.cn", "rich-billion", "hotmail", ".com.cn"},
}
func TestNameAndComOfEmailBatch(t *testing.T) {
	matches := NameAndComOfEmailBatch(emailBatch)
	for i, match := range matches {
		t.Logf("test %s...\n", verifyNameAndComOfEmailBatch[i][0])
		for j, ma := range match {
			if ma == verifyNameAndComOfEmailBatch[i][j] {
				t.Logf("OK, got: %s\n", ma)
			} else {
				t.Errorf("WRONG, expected: %s, got: %s\n", verifyNameAndComOfEmailBatch[i][j], ma)
			}
		}
	}
}