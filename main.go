package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"

	"telegram_history_parse/domain"
)

var (
	// Дата, до которой сообщения пропускаются
	dateFrom = time.Date(2025, time.January, 1, 1, 1, 1, 1, time.UTC)

	// Напечатать ли дополнительно в stdout таблицу с сообщениями
	printToStdout = true

	// Сколько строк записать в файл
	linesToWrite = 30
)

func main() {
	resultBytes, err := os.ReadFile("result.json")
	if err != nil {
		panic(err)
	}

	var history domain.HistoryDTO

	err = json.Unmarshal(resultBytes, &history)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(
		"converted.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	currentLine := 1

	for _, message := range history.Messages {
		if message.Date.Before(dateFrom) {
			continue
		}

		if currentLine > linesToWrite {
			break
		}

		if printToStdout {
			printTable(currentLine, message)
		}

		_, err = file.WriteString(message.Text.String() + "\n")
		if err != nil {
			panic(err)
		}

		currentLine++
	}
}

func printTable(lines int, message domain.MessageDTO) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendRows([]table.Row{
		{lines, message.Date.Format(time.DateTime), message.Text.String()},
	})
	t.Render()
	t.ResetRows()
}
