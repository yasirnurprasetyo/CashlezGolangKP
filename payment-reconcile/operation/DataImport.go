// This package covers all reconcile operations. The codes must be generic and
// supports all acquirers/issuers.
package operation

import (
	// "fmt"
	"math"
	"strconv"
	"strings"

	// "strconv"
	// "os"
	"errors"
	"path/filepath"

	"cashlez.com/common-util/crud"
	"cashlez.com/payment-reconcile/importer"
	"cashlez.com/payment-reconcile/util"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

var ExtSourceDsn string

// This inserts the imported data (in CSV format) to appropriate SQL table.
// Must be generic for any acquirer/issuer.
func InsertToRawTable(rawTable util.RawTable) bool {

	// Create connection to SQL for inserting data
	targetDB := &crud.DB{}
	targetDB.Connect(ExtSourceDsn)
	defer targetDB.Disconnect()

	completed := true

	var data []map[string]interface{}
	columns := strings.Split(rawTable.HeaderCsv, util.DataColumnSeparator)

	// Create basic form of insert placeholder
	p := []string{}
	for h := 0; h < len(columns); h++ {
		p = append(p, "?")
		log.Debug("header: " + columns[h])
	}
	p0 := "(" + strings.Join(p, ",") + ")"

	// Convert string array of raw data line to string map of interface,
	// since prepared statement expect interface form to insert the data
	for i, sourceLine := range rawTable.ContentCsv {
		log.Debug(sourceLine)
		if i > 0 {
			m := make(map[string]interface{})
			c := strings.Split(sourceLine, util.DataColumnSeparator)
			for j, s := range c {
				if j < len(columns) {
					var v interface{}
					if s != "" {
						v = s
					} else {
						v = nil
					}
					m[columns[j]] = v
				}
			}
			data = append(data, m)
		}
	}

	// To prevent a haavy process of inserting thousands of data at once, we split
	// the data insert into small chunks.
	chunkSize := 100
	l := float64(len(data)) / float64(chunkSize)
	loops := int(math.Ceil(l))

	log.Info("There are " + strconv.Itoa(len(data)) +
		" rows to copy, splitted into " +
		strconv.Itoa(loops) + " chunk(s), each with size of " +
		strconv.Itoa(chunkSize) + " rows")

	log.Info("Truncate the table first")
	_, err := targetDB.TryExecute("TRUNCATE TABLE " + rawTable.TableName)
	if err != nil {
		log.Error(err.Error())
		completed = false
	} else {
		for x := 0; x < loops; x++ {
			bulkdata := []interface{}{}

			y0 := x * chunkSize
			y9 := x*chunkSize + chunkSize
			if y9 > len(data) {
				y9 = len(data)
			}

			// Multiple the placeholder by size of a chunk
			placeholder := []string{}
			for y := y0; y < y9; y++ {
				placeholder = append(placeholder, p0)
				for _, cn := range columns {
					bulkdata = append(bulkdata, data[y][cn])
				}
			}

			// Form the prepared insert statement with prepared placeholder
			insert := "INSERT INTO " + rawTable.TableName + " (" + strings.Join(columns, ",") +
				") VALUES " + strings.Join(placeholder, ",")
			log.Println("Statement: " + insert)
			_, err := targetDB.TryExecutePrepared(insert, bulkdata)
			if err != nil {
				completed = false
			}
		}
	}
	return completed
}

// Generic data importer from XLSX files provided by all acquirers/issuers. This
// callable from the main function.
func ImportData(cfg *viper.Viper, sourceOwner string, sourceChannel string,
	sourceFilePath string, sourceDateColumn string, startDate string, endDate string,
	force bool) {

	/* This aims to run complete process of importing data from external resource (i.e.
	   i.e. XLSX file) to reformatting the data in SQL table.
	   Representative importer scripts of each acquirer/issuer must be located in
	   separated package.
	*/

	ExtSourceDsn = crud.GetFormattedDsn(
		cfg.Get("database.ext_source.host_ip").(string),
		cfg.Get("database.ext_source.port").(int),
		cfg.Get("database.ext_source.schema").(string),
		cfg.Get("database.ext_source.username").(string),
		cfg.Get("database.ext_source.password").(string))

	var err error
	log.Debug("Load source data")
	ext := filepath.Ext(sourceFilePath)
	log.Debug("ext: " + ext)
	if ext != "" {

		// Ensure we support the resource type
		ext := strings.ToLower(ext)
		if ext == ".xls" { //".xls"
			log.Fatal("XLSX format is required instead of XLS") //log.Fatal("XLSX format is required instead of XLS")
		} else if ext == ".xlsx" || ext == ".csv" {
			log.Debug("Passed")
			// ok, skipped
		} else {
			err = errors.New("Unsupported file type")
		}

		// Calling appropriate extractor based on acquirer/issuer
		var rawData *util.RawData
		sourceFileType := strings.TrimLeft(strings.ToUpper(ext), ".")
		if sourceOwner == "BankMandiri" {
			if sourceChannel == "CreditDebit" || sourceChannel == "DebitCredit" {
				rawData, err = importer.Extract_BankMandiri_CreditDebit(
					sourceFilePath, sourceFileType, cfg)
			}
		} else if sourceOwner == "BRI" || sourceOwner == "BankBRI" {
			if sourceChannel == "Credit" {
				rawData, err = importer.Extract_BRI_Details(sourceFilePath,
					sourceFileType, cfg)
			}
		} else if sourceOwner == "BNI46" || sourceOwner == "BankBNI46" {
			if sourceChannel == "Credit" {
				rawData, err = importer.Extract_BNI46_Credit(sourceFilePath,
					sourceFileType, cfg)
			}
		} else if sourceOwner == "LinkAja" {
			if sourceChannel == "Credit" {
				rawData, err = importer.Extract_LinkAja_Details(sourceFilePath,
					sourceFileType, cfg)
			}
		} else {
			err = errors.New("Unknown source owner")
		}

		// Insert extracted raw data to SQL table
		if rawData.RawDetails.TableName != "" &&
			rawData.RawDetails.ContentCsv != nil &&
			len(rawData.RawDetails.ContentCsv) > 0 {
			InsertToRawTable(rawData.RawDetails)
		}
		if rawData.RawSummary.TableName != "" &&
			rawData.RawSummary.ContentCsv != nil &&
			len(rawData.RawSummary.ContentCsv) > 0 {
			InsertToRawTable(rawData.RawSummary)
		}

	} else {
		err = errors.New("No file extention to recognize")
	}
	if err != nil {
		log.Error(err.Error())
	}
}
