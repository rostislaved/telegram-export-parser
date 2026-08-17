package main

import (
	"log/slog"
	"os"
	"time"

	"telegram-export-parser/internal/app"
)

var (
	// Skip messages older than this date.
	dateFrom = time.Date(2005, time.January, 1, 1, 1, 1, 1, time.UTC)

	// Print the messages table to stdout as well.
	printToStdout = false

	// Number of lines to write to the file. Negative values and 0 mean no limit.
	linesToWrite = 0

	outputFilename = "parser_output.txt"
	inputFilename  = "result.json"

	outputType = "json"
)

func main() {
	config := app.Config{
		InputFilename:  inputFilename,
		OutputFilename: outputFilename,
		LinesToWrite:   linesToWrite,
		DateFrom:       dateFrom,
		PrintToStdout:  printToStdout,
		OutputType:     outputType,
	}

	err := app.Start(config)
	if err != nil {
		slog.Info(err.Error())

		os.Exit(1)
	}
}
