package controllers

import (
    "regexp"
)

// validDate is regular expression for valid date
const validDate = `(?:2[0-9]{3})\-(?:(?:0[0-9])|(?:1[0-2]))\-(?:(?:[0-2][0-9])|(?:3[01]))`

// IsValidDate checks date format
func IsValidDate(date string) bool {
    var re = regexp.MustCompile(validDate)
    return re.Match([]byte(date))
}
