package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"

	"telegram-export-parser/domain/history"
	message_pkg "telegram-export-parser/domain/message"
	"telegram-export-parser/domain/telegram_time"
	"telegram-export-parser/stopwatch"
)

var (
	// Skip messages older than this date.
	dateFrom = time.Date(2005, time.January, 1, 1, 1, 1, 1, time.UTC)

	// Print the messages table to stdout as well.
	printToStdout = false

	// Number of lines to write to the file. Negative values and 0 mean no limit.
	linesToWrite = 0

	inputFilename  = "result.json"
	outputFilename = "output.txt"

	outputType = "json"
)

func main() {
	err := Start()
	if err != nil {
		panic(err)
	}

}

func Start() (err error) {
	s := stopwatch.New()

	inputFileInfo, err := os.Stat(inputFilename)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf(" input file size: %s", formatFileSize(inputFileInfo.Size())))

	resultBytes, err := os.ReadFile(inputFilename)
	if err != nil {
		return err
	}

	s.Toc("Open input file")

	var history history.History

	err = json.Unmarshal(resultBytes, &history)
	if err != nil {
		return err
	}

	s.Toc("Unmarshal")

	outputFile, err := os.OpenFile(
		outputFilename,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0o600,
	)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := outputFile.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	currentLine := 1

	pms := make([]ParsedMessage, 0, 1000)

	for _, message := range history.Messages {
		if message.Date.Before(dateFrom) {
			continue
		}

		if message.Type == message_pkg.Service {
			continue
		}

		if linesToWrite > 0 {
			if currentLine > linesToWrite {
				break
			}
		}

		if printToStdout {
			printTable(currentLine, message)
		}

		pm := ParsedMessage{
			ID:        message.ID,
			ReplyToID: message.ReplyToMessageID,
			Timestamp: message.Date,
			Text:      message.Text.String(),
		}

		pms = append(pms, pm)

		currentLine++
	}

	switch outputType {
	case "json":
		pmsBytes, err := json.MarshalIndent(pms, " ", " ")
		if err != nil {
			return err
		}

		_, err = outputFile.Write(pmsBytes)
		if err != nil {
			panic(err)
		}
	case "text":
		for _, pm := range pms {
			_, err = outputFile.WriteString(pm.Text + "\n")
			if err != nil {
				panic(err)
			}
		}
	}

	s.Toc("Write to output file")

	outputFileInfo, err := os.Stat(outputFilename)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("output file size: %s\n", formatFileSize(outputFileInfo.Size())))

	s.PrintFromPrevious()

	return nil
}

type ParsedMessage struct {
	ID        int                        `json:"id"`
	ReplyToID int                        `json:"reply_to_id,omitzero"`
	Timestamp telegram_time.TelegramTime `json:"date"`
	Text      string                     `json:"text"`
}

func printTable(lines int, message message_pkg.Message) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendRows([]table.Row{
		{lines, message.From, message.Date.Format(time.DateTime), message.Text.String()},
	})
	t.Render()
	t.ResetRows()
}

func formatFileSize(size int64) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)

	if size < kb {
		return fmt.Sprintf("%d байт", size)
	}

	if size < mb {
		return fmt.Sprintf("%.2f КБ", float64(size)/kb)
	}

	if size < gb {
		return fmt.Sprintf("%.2f МБ", float64(size)/mb)
	}

	return fmt.Sprintf("%.2f ГБ", float64(size)/gb)
}
