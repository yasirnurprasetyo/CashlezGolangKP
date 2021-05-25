package importer

import (
	"strconv"
	"time"

	"cashlez.com/payment-reconcile/util"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type BriDetails struct {
	NO                int64     `db:"no"`
	NamaMerchant      string    `db:"nama_merchant"`
	RkDate            time.Time `db:"rk_date"`
	ProcDate          time.Time `db:"proc_date"`
	MID               string    `db:"mid"`
	Cardtype          string    `db:"cardtype"`
	TrxDate           time.Time `db:"trx_date"`
	Auth              string    `db:"auth"`
	CardNo            string    `db:"cardno"`
	JenisTrx          string    `db:"jenis_trx"`
	Amount            int64     `db:"amount"`
	Nonfare           int64     `db:"nonfare"`
	Rate              int64     `db:"rate"`
	DiscAmt           int64     `db:"disc_amt"`
	Airfare           int64     `db:"airfare"`
	Flag              int64     `db:"flag"`
	NetAmt            int64     `db:"net_amt"`
	MerchantRefNumber string    `db:"merchant_ref_number"`
}

// Extracts raw data from XLSX file provided by BNI46 to CSV format prior to
// table inserts
func Extract_BRI_Details(sourceFilePath string, sourceFileType string,
	config *viper.Viper) (*util.RawData, error) {
	log.Debug("Extract BRI Credit data")

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
				// XlsxHasTitle:    true,
				// XlsxHasClosing: true,
				// XlsxOpeningWord: "BNI46",
				// XlsxClosingWord: "STOP",
				DateRegex: config.
					Get("ext_source.date_column.regex." +
						"bri_details").(string),
				DatePattern: config.
					Get("ext_source.date_column.pattern." +
						"bri_details").(string),
				DatePosition: config.
					Get("ext_source.date_column.position." +
						"bri_details").([]interface{}),
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
				Get("ext_source.raw_table.bri_details").(string)
			detailsTable.HeaderCsv = xHdr
			detailsTable.ContentCsv = xContent
		}
		rawData.RawDetails = *detailsTable
	}

	return rawData, err
}
