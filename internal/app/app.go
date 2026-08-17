package app

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"

	"telegram-export-parser/internal/app/domain/datetime"
	"telegram-export-parser/internal/app/domain/export_dto"
	message_pkg "telegram-export-parser/internal/app/domain/message"
	"telegram-export-parser/internal/app/domain/text_entity"
	"telegram-export-parser/internal/pkg/stopwatch"
)

type Config struct {
	InputFilename  string
	OutputFilename string
	LinesToWrite   int
	DateFrom       time.Time
	PrintToStdout  bool
	OutputType     string
}

// nolint:revive,cyclop //cyclomatic after Start refactoring
func Start(config Config) (err error) {
	s := stopwatch.New()

	inputFileInfo, err := os.Stat(config.InputFilename)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf(" input file size: %s", FormatFileSize(inputFileInfo.Size())))

	resultBytes, err := os.ReadFile(config.InputFilename)
	if err != nil {
		return err
	}

	s.Toc("Open input file")

	var exportDTO export_dto.ExportDTO

	err = json.Unmarshal(resultBytes, &exportDTO)
	if err != nil {
		return err
	}

	s.Toc("Unmarshal")

	outputFile, err := os.OpenFile(
		config.OutputFilename,
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

	for _, message := range exportDTO.Messages {
		if message.Date.Before(config.DateFrom) {
			continue
		}

		if message.Type == message_pkg.Service {
			continue
		}

		if config.LinesToWrite > 0 {
			if currentLine > config.LinesToWrite {
				break
			}
		}

		if config.PrintToStdout {
			PrintTable(currentLine, message)
		}

		text := lo.Reduce(
			message.TextEntities,
			func(agg string, item text_entity.TextEntity, _ int) string {
				return agg + item.String()
			},
			"",
		)

		pm := ParsedMessage{
			ID:        message.ID,
			ReplyToID: message.ReplyToMessageID,
			Timestamp: message.Date,
			Text:      text,
		}

		pms = append(pms, pm)

		currentLine++
	}

	switch config.OutputType {
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
	default:
		panic("unreachable")
	}

	s.Toc("Write to output file")

	outputFileInfo, err := os.Stat(config.OutputFilename)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("output file size: %s\n", FormatFileSize(outputFileInfo.Size())))

	s.PrintFromPrevious()

	return nil
}

func FormatFileSize(size int64) string {
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

func PrintTable(lines int, message message_pkg.Message) {
	text := lo.Reduce(
		message.TextEntities,
		func(agg string, item text_entity.TextEntity, _ int) string {
			return agg + item.String()
		},
		"",
	)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendRows([]table.Row{
		{lines, message.From, message.Date.Format(time.DateTime), text},
	})
	t.Render()
	t.ResetRows()
}

type ParsedMessage struct {
	ID        int               `json:"id"`
	ReplyToID int               `json:"reply_to_id,omitzero"`
	Timestamp datetime.DateTime `json:"date"`
	Text      string            `json:"text"`
}
