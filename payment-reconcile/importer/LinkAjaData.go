package importer

import (
	"strconv"
	"time"

	"cashlez.com/payment-reconcile/util"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Struct type of CreditDebit trx details provided by LinkAjaDetail
// The properties also represents columns in related SQL table
type LinkAjaDetail struct {
	TopOrganization      string    `db:"top_organization"`
	Organization         string    `db:"organization"`
	TransactionID        string    `db:"transaction_id"`
	InvoiceID            string    `db:"invoice_id"`
	FinalizedDate        time.Time `db:"finalized_date"`
	FinalizedTime        time.Time `db:"finalized_time"`
	InitiatedDate        time.Time `db:"initiated_date"`
	InitiatedTime        time.Time `db:"initiated_time"`
	TransactionType      string    `db:"transaction_type"`
	TransactionScenario  string    `db:"transaction_scenario"`
	TransactionStatus    string    `db:"transaction_status"`
	TransactionStatement string    `db:"transaction_statement"`
	Account              string    `db:"account"`
	CounterParty         string    `db:"counter_party"`
	Debit                int64     `db:"debit"`
	Credit               int64     `db:"credit"`
	Balance              int64     `db:"balance"`
}

// Extracts raw data from XLSX file provided by LinkAja to CSV format prior to
// table inserts
func Extract_LinkAja_Details(sourceFilePath string, sourceFileType string,
	config *viper.Viper) (*util.RawData, error) {
	log.Debug("Extract LinkAja detail data...")

	var err error
	rawData := &util.RawData{
		RequiresDenormalization: true,
	}

	// Ensure the data source is an XLSX type
	if sourceFileType != "XLSX" {
		log.Fatal("Only XLSX data is supported")
	} else {
		detailsTable := &util.RawTable{
			SourceConfig: util.SourceConfiguration{
				XlsxStartColumn: "B",
				XlsxHasTitle:    true,
				XlsxHasClosing:  true,
				XlsxOpeningWord: "This report provides transaction information as per selected criteria",
				XlsxClosingWord: "Total",
				DateRegex: config.
					Get("ext_source.date_column.regex." +
						"linkaja_details").(string),
				DatePattern: config.
					Get("ext_source.date_column.pattern." +
						"linkaja_details").(string),
				DatePosition: config.
					Get("ext_source.date_column.position." +
						"linkaja_details").([]interface{}),
			},
		}
		x, err := util.GetRawTableFromXlsx(sourceFilePath, detailsTable.SourceConfig)
		if err != nil {
			log.Error(err.Error())
		}
		if len(x) > 1 {
			log.Debug("Details data length: " + strconv.Itoa(len(x)))
			xHdr, xContent, err := util.ParseRawTable(x, detailsTable.SourceConfig)
			if err != nil {
				log.Error(err.Error())
			}
			detailsTable.TableName = config.
				Get("ext_source.raw_table.linkaja_details").(string)
			detailsTable.HeaderCsv = xHdr
			detailsTable.ContentCsv = xContent
		}
		rawData.RawDetails = *detailsTable
	}
	return rawData, err
}
