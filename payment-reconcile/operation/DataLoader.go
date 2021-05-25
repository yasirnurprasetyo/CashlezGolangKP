// This package covers all reconcile operations. The codes must be generic and
// supports all acquirers/issuers.
package operation

import (
	"cashlez.com/common-util/crud"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// DSN for loading trx data from internal source to SQL table
var IntSourceDsn string

func LoadData(cfg *viper.Viper, sourceOwner string, sourceChannel string,
	sourceFilePath string, startDate string, endDate string, force bool) {

	IntSourceDsn = crud.GetFormattedDsn(
		cfg.Get("database.int_source.host_ip").(string),
		cfg.Get("database.int_source.port").(int),
		cfg.Get("database.int_source.schema").(string),
		cfg.Get("database.int_source.username").(string),
		cfg.Get("database.int_source.password").(string))

	log.Debug("Load internal data")
}
