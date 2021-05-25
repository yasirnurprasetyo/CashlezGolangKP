// This package contains both extractor and formatter for trx data provided by
// acquirers/issuers, covers all trx channels by acquirer/issuer.
package importer

import (
	// "strings"

	"strconv"
	"time"

	"cashlez.com/payment-reconcile/util"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type BankBNICreditDebit struct {
	Procdate         time.Time `db:"proc_date"`
	MID              string    `db:"mid"`
	OB               string    `db:"ob"`
	GB               string    `db:"gb"`
	SEQ              string    `db:"seq"`
	Type             string    `db:"type"`
	TransactionDate  time.Time `db:"trx_date"`
	Auth             string    `db:"auth"`
	CardNo           string    `db:"card_no"`
	Amount           int64     `db:"amount"`
	TID              string    `db:"tid"`
	JenisTrx         string    `db:"jenis_trx"`
	PTR              string    `db:"ptr"`
	Rate             int64     `db:"rate"`
	DiscAmount       int64     `db:"disc_amount"`
	AirFare          string    `db:"air_fare"`
	Plan             string    `db:"plan"`
	SsAmount         int64     `db:"ss_amount"`
	SsFeeType        int64     `db:"ss_fee_type"`
	Flag             string    `db:"flag"`
	NettAmount       int64     `db:"nett_amount"`
	MerchantAccounnt string    `db:"merchant_account"`
	MerchantName     string    `db:"merchant_name"`
}

// Extracts raw data from XLSX file provided by BNI46 to CSV format prior to
// table inserts
func Extract_BNI46_Credit(sourceFilePath string, sourceFileType string,
	config *viper.Viper) (*util.RawData, error) {
	log.Debug("Extract BNI46 Credit data")

	var err error
	rawData := &util.RawData{
		RequiresDenormalization: true,
	}

	if sourceFileType != "XLSX" {
		log.Fatal("Only XLSX data is supported")
	} else {
		// Details
		detailsTable := &util.RawTable{
			SourceConfig: util.SourceConfiguration{
				XlsxStartColumn: "A",
				XlsxHasClosing:  true,
				XlsxClosingWord: "STOP",
				DateRegex: config.
					Get("ext_source.date_column.regex." +
						"bni46_credit").(string),
				DatePattern: config.
					Get("ext_source.date_column.pattern." +
						"bni46_credit").(string),
				DatePosition: config.
					Get("ext_source.date_column.position." +
						"bni46_credit").([]interface{}),
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
				Get("ext_source.raw_table.bni46_credit").(string)
			detailsTable.HeaderCsv = xHdr
			detailsTable.ContentCsv = xContent
		}
		rawData.RawDetails = *detailsTable
	}

	return rawData, err
}
