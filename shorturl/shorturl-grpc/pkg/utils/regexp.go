package utils

import "regexp"

func IsUrl(url string) bool {
	pattern := `^(http|https)://[a-zA-z0-9\-\.]+\.[a-zA-Z]{2,}(?:/[^/]*)*$`
	regExp := regexp.MustCompile(pattern)
	return regExp.MatchString(url)
}
