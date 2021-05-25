// This package contains both extractor and formatter for trx data provided by
// acquirers/issuers, covers all trx channels by acquirer/issuer
package importer

import (
	"database/sql"
	// "strings"
	"strconv"
	"time"

	"cashlez.com/payment-reconcile/util"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Struct type of CreditDebit trx details provided by Bank Mandiri
// The properties also represents columns in related SQL table
type BankMandiriCreditDebitDetail struct {
	MID                     string
	MerchantOfficial        string         `db:"merchant_official"`
	TradingName             string         `db:"trading_name"`
	BankMandiriAccount      string         `db:"bank_mandiri_acc"`
	OtherBankAccount        sql.NullString `db:"other_bank_acc"`
	MerchantAccounnt        sql.NullString `db:"merchacc"`
	TransactionDate         time.Time      `db:"trxdate"`
	SettlementDate          time.Time      `db:"settledate"`
	TransactionCode         sql.NullString `db:"trxcode"`
	Description             sql.NullString
	CardPan                 string `db:"card"`
	CardType                string `db:"crdtype"`
	TID                     string
	AuthCode                string
	PaymentBatch            sql.NullString
	TidBatch                sql.NullString
	BatchSequence           sql.NullString `db:"batchseq"`
	TransactionAmount       int64          `db:"amount"`
	NonMdrTransactionAmount int64          `db:"nonmdramount"`
}

// Struct of CreditDebit trx summary provided by Bank Mandiri
// The properties also represents columns in related table
type BankMandiriCreditDebitSummary struct {
	MID string
}

// Extracts raw data from XLSX file provided by Bank Mandiri to CSV format prior to
// table inserts
func Extract_BankMandiri_CreditDebit(sourceFilePath string, sourceFileType string,
	config *viper.Viper) (*util.RawData, error) {
	log.Debug("Extract Bank Mandiri Credit Debit data...")

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
				XlsxStartColumn: "A",
				XlsxHasTitle:    true,
				XlsxHasClosing:  true,
				XlsxOpeningWord: "DETAIL",
				XlsxClosingWord: "TOTAL",
				DateRegex: config.
					Get("ext_source.date_column.regex." +
						"bankmandiri_creditdebit_details").(string),
				DatePattern: config.
					Get("ext_source.date_column.pattern." +
						"bankmandiri_creditdebit_details").(string),
				DatePosition: config.
					Get("ext_source.date_column.position." +
						"bankmandiri_creditdebit_details").([]interface{}),
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
				Get("ext_source.raw_table.bankmandiri_creditdebit_details").(string)
			detailsTable.HeaderCsv = xHdr
			detailsTable.ContentCsv = xContent
		}

		summaryTable := &util.RawTable{
			SourceConfig: util.SourceConfiguration{
				XlsxStartColumn: "A",
				XlsxHasTitle:    true,
				XlsxHasClosing:  true,
				XlsxOpeningWord: "SUMMARY GROUP",
				XlsxClosingWord: "TOTAL",
				DateRegex: config.
					Get("ext_source.date_column.regex." +
						"bankmandiri_creditdebit_summary").(string),
				DatePattern: config.
					Get("ext_source.date_column.pattern." +
						"bankmandiri_creditdebit_summary").(string),
				DatePosition: config.
					Get("ext_source.date_column.position." +
						"bankmandiri_creditdebit_summary").([]interface{}),
			},
		}
		y, err := util.GetRawTableFromXlsx(sourceFilePath, summaryTable.SourceConfig)
		if err != nil {
			log.Error(err.Error())
		}
		if len(y) > 1 {
			log.Debug("Summary data length: " + strconv.Itoa(len(y)))
			yHdr, yContent, err := util.ParseRawTable(y, summaryTable.SourceConfig)
			if err != nil {
				log.Error(err.Error())
			}
			summaryTable.TableName = config.
				Get("ext_source.raw_table.bankmandiri_creditdebit_summary").(string)
			summaryTable.HeaderCsv = yHdr
			summaryTable.ContentCsv = yContent
		}

		rawData.RawDetails = *detailsTable
		rawData.RawSummary = *summaryTable
	}
	return rawData, err
}

// Merge trx details with its summary to obtain complete data of each trx, because
// some of information in trx details was normalized into the summary
func Denomarlize_BankMandiri_CreditDebit(rawDetailsTable string,
	rawSummaryTable string, config *viper.Viper) ([]string, error) {

	/* Function draft. Requires complete codes.
	 */

	var err error
	var denomarlized []string

	return denomarlized, err
}
