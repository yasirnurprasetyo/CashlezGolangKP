// This is the main package
package main

import (
	"flag"
	"os"
	"path"
	"strings"
	"time"

	"cashlez.com/payment-reconcile/operation"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// This to obtain path of installed GO-lang in current environment
func GetGoPath() string {
	gopath := os.Getenv("GOPATH")
	if strings.Index(gopath, ":") > -1 {
		gp := strings.Split(gopath, ":")
		gopath = gp[len(gp)-1]
	}
	return gopath
}

// This to obtain path of project application relative to installed GO-lang path
func GetProjectSourcePath() string {
	return path.Join(GetGoPath(), "src/cashlez.com/payment-reconcile")
}

// This to obtain path of current log file
func getLogFilePath(operation string, sourceOwner string, sourceChannel string) string {
	gopath := os.Getenv("GOPATH")
	if strings.Index(gopath, ":") > -1 {
		gp := strings.Split(gopath, ":")
		gopath = gp[len(gp)-1]
	}
	now := time.Now()
	return path.Join(gopath, "log/reconcile_"+operation+"__"+sourceOwner+"__"+sourceChannel+"__"+now.Format("2006-01-02")+".log")
}

// Main function to received request parameters and routes the request to appropriate
// operation
func main() {

	op := flag.String("op", "operation", "operation to run")
	startDate := flag.String("startDate", "2016-01-01", "date to start")
	endDate := flag.String("endDate", "2016-01-01", "date to end")
	// targetTable := flag.String("target", "dst_table_name", "destination table")
	sourceDateColumn := flag.String("sourceDateColumn", "created_date", "field of the date")
	sourceFilePath := flag.String("sourcePath", "csv", "path of source file")
	// sourceFileType := flag.String("sourceType", "csv", "type of source file")
	owner := flag.String("sourceOwner", "source_owner", "owner who provides the source")
	channel := flag.String("sourceChannel", "source_channel", "channel which data was recorded")
	force := flag.Bool("force", false, "force to execute although it's already success")

	flag.Parse()

	sourceOwner := strings.Replace(*owner, " ", "", -1)
	sourceChannel := strings.Replace(*channel, " ", "", -1)

	log.SetLevel(log.DebugLevel)
	file, err := os.OpenFile(getLogFilePath(strings.ToLower(*op), sourceOwner, sourceChannel),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(file)
	}
	defer file.Close()

	cfg := viper.New()
	cfg.AddConfigPath(GetProjectSourcePath())
	cfg.SetConfigName("settings")
	cfg.SetConfigType("yaml")
	err = cfg.ReadInConfig()
	if err != nil {
		log.Error(err.Error())
	}

	*op = strings.ToUpper(*op)
	switch *op {
	case "LOAD":
		operation.ImportData(cfg, sourceOwner, sourceChannel, *sourceFilePath,
			*sourceDateColumn, *startDate, *endDate, *force)
	default:
		log.Fatal("No operations was defined")
	}

}
