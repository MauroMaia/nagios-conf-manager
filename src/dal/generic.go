package dal

import (
	"regexp"
	"strings"
)

var reComment = regexp.MustCompile(`^( )*#`)
var reStartHost = regexp.MustCompile(`^define host.*`)
var reStartHostGroup = regexp.MustCompile(`^define hostgroup.*`)
var reStartCommand = regexp.MustCompile(`^define command.*`)
var reEndDefineStatment = regexp.MustCompile(`^}`)
var inLineComment = regexp.MustCompile(` *;.*`)

func RemoveComments(str string) string {
	return inLineComment.ReplaceAllString(str, "")
}

func IgnoreConfigLine(str string) bool {
	// Ignore comments and empty lines
	if strings.Compare(str, "") == 0 || reComment.MatchString(str) {
		return true
	}
	return false
}
