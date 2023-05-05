package main

import (
	"flag"
	"log"
	"os"

	"github.com/EzekielKemboi/moneyspend/bolt"
	"github.com/EzekielKemboi/moneyspend/logic"
)

const (
	version = "1.0.3"
)

var isVersion = flag.Bool("v", false, "show the version.")
var reClass = flag.Bool("r", false, "reclass all the items.")
var classify = flag.Bool("c", false, "do the classify procedure.")
var moneyCsvFileName = flag.String("n", "data.csv", "the name of money csv file.")

func main() {
	preProcess()

	spendData := logic.ProcessMoneyCsv(*moneyCsvFileName)

	logic.GenerateOutput(spendData)

	logic.ClassifyItems(spendData, *classify)
}

func preProcess() {
	flag.Parse()
	if *isVersion {
		log.Printf("moneyspend version: %v", version)
		os.Exit(0)
	}

	bolt.MustInitBolt(*reClass)
}
