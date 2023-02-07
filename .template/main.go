package main

import (
	"fmt"
	"os"

	"github.com/ghostsecurity/PROJECTNAME/pkg/logger"
)

var (
	debugEnabled = true
)

func main() {
	// configure the logger with output and debug status
	logger.Configure(os.Stdout, debugEnabled)

	logger.Debug("we're starting up the engines")
	logger.Info("let's do the impossible")
	fmt.Print("\nHello, World!\n\n")
	logger.Error("something went wrong")
}
