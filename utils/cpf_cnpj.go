package utils

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

func IsCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, size, pos)
}

func IsCNPJ(doc string) bool {
	const (
		size = 12
		pos  = 5
	)

	return isCPFOrCNPJ(doc, CNPJRegexp, size, pos)
}

func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := calculateDigit(d, position)

	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

func cleanNonDigits(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

func calculateDigit(doc string, position int) string {
	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

func toInt(r rune) int {
	return int(r - '0')
}
