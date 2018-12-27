package main

import (
	"flag"
	"os"

	"github.com/farhaanbukhsh/file-indexer/lib/conf"
	"github.com/farhaanbukhsh/file-indexer/lib/indexer"
	"github.com/farhaanbukhsh/file-indexer/lib/server"
)

var configFlag = flag.String("c", "", "Use alternative config file.")
var verboseFlag = flag.Bool("v", false, "Verbosely print the log output to Standard Output.")
var helpFlag = flag.Bool("h", false, "Print this help information.")

func main() {
	flag.Parse()

	if *helpFlag == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	config := conf.NewConfig(*configFlag, *verboseFlag)
	go indexer.NewIndex(config)
	server.NewServer(config).Start()
}
