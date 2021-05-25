// This package covers common utilities.
package util

// Column seoarator for data CSV.
const DataColumnSeparator = ";"

// Configuration for each data source.
type SourceConfiguration struct {
	// FilePath         string
	XlsxStartRow    int
	XlsxStartColumn string
	XlsxHasTitle    bool
	XlsxHasClosing  bool
	XlsxOpeningWord string
	XlsxClosingWord string
	DateRegex       string
	DatePattern     string
	DatePosition    []interface{}
}

// Structure of data table for each details and summary.
type RawTable struct {
	TableName    string
	SourceConfig SourceConfiguration
	HeaderCsv    string
	ContentCsv   []string
}

// Structure of raw data, both details and summary.
type RawData struct {
	RawDetails              RawTable
	RawSummary              RawTable
	RequiresDenormalization bool
}
