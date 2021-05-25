// This package covers common utilities
package util

import (
	// "fmt"
	"regexp"
	"strings"

	// "strconv"
	log "github.com/Sirupsen/logrus"
)

// dd-mm-yyyy
// `/^([0-9]{1,2})+\\-+([0-9]{1,2})+\\-+([0-9]{1,4})$`
// dd/mm/yyyy
// `/^([0-9]{1,2})+\\/+([0-9]{1,2})+\\/+([0-9]{1,4})$`
// yyyy-mm-dd
// `/^([0-9]{1,4})+\\-+([0-9]{1,2})+\\-+([0-9]{1,2})$`
// yyyy/mm/dd
// `/^([0-9]{1,4})+\\/+([0-9]{1,2})+\\/+([0-9]{1,2})$`
// dd/mm/yy
// `/^([0-9]{1,2})+\\/+([0-9]{1,2})+\\/+([0-9]{1,2})$`

// Convert 2-digits year to 4-digits
func getFullYear(s string) string {
	if len(s) == 2 {
		return "20" + s
	}
	return s
}

// Convert date from its origin format to Sql date format, based on given data regex
func ToSqlDateFormat(s string, regex string, pattern string) (string, error) {
	var err error
	var d string

	date := make(map[string]int)
	a := strings.Split(strings.Trim(pattern, " "), "")
	for i, s := range a {
		if s == "D" {
			date["day"] = i + 1
		} else if s == "M" {
			date["month"] = i + 1
		} else if s == "Y" {
			date["year"] = i + 1
		}
	}
	if re, err := regexp.Compile(regex); err == nil {
		sm := re.FindStringSubmatch(s)
		if len(sm) == 4 {
			d = getFullYear(sm[date["year"]]) + "-" + sm[date["month"]] + "-" +
				sm[date["day"]]
		}
	} else {
		log.Error(err.Error())
	}
	return d, err
}
