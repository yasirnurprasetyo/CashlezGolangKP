// This package covers common utilities within data reconcile.
package util

import (
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	log "github.com/Sirupsen/logrus"
)

// To obtain raw data header passed all string validations. This ensure each header
// is clean, well-formatted and does not contain characters that raise error during
// the sql insert.
func GetValidatedHeader(a []string) string {
	for i, _ := range a {
		a[i] = strings.Replace(a[i], "\"", "", -1)
		a[i] = strings.Replace(a[i], "\\/", "", -1)
		a[i] = strings.Replace(a[i], ",", ".", -1)
		a[i] = strings.Replace(a[i], ";", ".", -1)
		a[i] = strings.Replace(strings.ToLower(strings.Trim(a[i], " ")), " ", "_", -1)
	}
	return strings.Join(a, DataColumnSeparator)
}

// To obtain raw data rows passed all string validations. This ensure each data row
// is clean, well-formatted and does not contain characters that raise error during
// the sql insert.
func GetValidatedContent(a []string, datePosition []int,
	config SourceConfiguration) string {
	var err error
	for i, _ := range a {
		a[i] = strings.Trim(a[i], " ")

		// We need to check whether the current column match datePosition
		var isDate bool
		if len(datePosition) > 0 {
			for _, x := range datePosition {
				if x == i {
					isDate = true
				}
			}
		}
		if isDate {
			a[i], err = ToSqlDateFormat(strings.Trim(a[i], " "), config.DateRegex,
				config.DatePattern)
			if err != nil {
				panic(err)
			}
		} else {
			a[i] = strings.Replace(a[i], "\"", "", -1)
			a[i] = strings.Replace(a[i], "\\/", "", -1)
			a[i] = strings.Replace(a[i], ",", ".", -1)
			a[i] = strings.Replace(a[i], ";", ".", -1)

			a[i] = strings.Replace(a[i], "\xBB", "", -1)
			a[i] = strings.Replace(a[i], "\xBD", "", -1)
			a[i] = strings.Replace(a[i], "\xBF", "", -1)
			a[i] = strings.Replace(a[i], "\xEF", "", -1)
			a[i] = strings.Replace(a[i], "\xFE", "", -1)
			a[i] = strings.Replace(a[i], "\xFF", "", -1)
		}
	}
	return strings.Join(a, DataColumnSeparator)
}

// Parse received raw trx data, to separate header from data row and validate both
// header and data row.
func ParseRawTable(data [][]string, config SourceConfiguration) (string, []string,
	error) {
	var err error
	var header string
	var content []string
	if len(data) > 1 {
		// Get date columns position in integer array from configuration settings
		dp := GetDatePositionInt(config.DatePosition)
		// Iterate all rows and separate the header
		for i, r := range data {
			if i == 0 {
				header = GetValidatedHeader(r)
			} else {
				content = append(content, GetValidatedContent(r, dp, config))
			}
		}
	}
	return header, content, err
}

// Read raw trx data from XLSX file.
func GetRawTableFromXlsx(filePath string, config SourceConfiguration) ([][]string,
	error) {

	xls, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	sheets := xls.GetSheetList()
	startColIx := IntOfPointerChars(config.XlsxStartColumn)
	log.Debug("startColIx: " + strconv.Itoa(startColIx))
	var data [][]string

	read := false
	rows, err := xls.GetRows(sheets[0])
	for i, row := range rows {
		if len(row) > 0 {
			var cols []string
			for j, cell := range row {
				if j >= startColIx {
					// When data has closing word and match the current row
					if config.XlsxHasClosing &&
						strings.Index(cell, config.XlsxClosingWord) > -1 {
						read = false
						log.Debug("Stop reading by this line")
					}
					// When data has no titles before its header and starting point
					// match the current row
					if !read && !config.XlsxHasTitle && i == config.XlsxStartRow {
						read = true
						log.Debug("Start reading by this line")
					}

					// When read-flag was true, read the line as raw data, stop
					// reading when it's false
					if read {
						cols = append(cols, cell)
					}

					// When data has titles before its header and match the current
					// row
					if !read && config.XlsxHasTitle &&
						strings.Index(cell, config.XlsxOpeningWord) > -1 {
						read = true
						log.Debug("Start reading by next line")
					}
				}
			}
			if len(cols) > 0 {
				data = append(data, cols)
			}
		}
	}
	log.Debug("data rows: " + strconv.Itoa(len(data)))
	return data, err
}

// Get integer array of date columns position in raw trx data from given pointer
// characters defined in source configuration settings
func GetDatePositionInt(a []interface{}) []int {
	var i []int
	if len(a) > 0 {
		for _, x := range a {
			if x != nil {
				s := x.(string)
				i = append(i, IntOfPointerChars(s))
			}
		}
	}
	return i
}
