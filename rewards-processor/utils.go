package main

import "regexp"

func isAlphaNumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isInteger(floatString string) bool {
	re := regexp.MustCompile(`^\d+(\.0+)?$`)
	return re.MatchString(floatString)
}
