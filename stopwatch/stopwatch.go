package stopwatch

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type stopwatch struct {
	snapshots []snapshot
	logger    *slog.Logger
}

func New() stopwatch {
	initialSnapshot := newSnapshot("")

	return stopwatch{
		snapshots: []snapshot{initialSnapshot},
		logger:    slog.Default(),
	}
}

func (s *stopwatch) SetLogger(logger *slog.Logger) {
	if logger == nil {
		s.logger = slog.Default()

		return
	}

	s.logger = logger
}

func (s *stopwatch) Toc(messages ...string) {
	var joinedMessages string

	if len(messages) > 0 {
		joinedMessages = strings.Join(messages, ", ")
	}

	s.snapshots = append(s.snapshots, newSnapshot(joinedMessages))
}

func (s *stopwatch) PrintFromPrevious() {
	if len(s.snapshots) == 1 {
		s.logger.Info("No data")

		return
	}

	intervals := getTimeIntervals(s)

	for i, interval := range intervals {
		message := fmt.Sprintf(
			"[%d] %d-%d: %v %s",
			i+1,
			s.snapshots[i].lineNumber,
			s.snapshots[i+1].lineNumber,
			interval.Round(time.Millisecond),
			s.snapshots[i].joinedMessages,
		)

		s.logger.Info(message)
	}
}

func getTimeIntervals(s *stopwatch) []time.Duration {
	intervals := make([]time.Duration, 0, len(s.snapshots)-1)

	for i := range s.snapshots {
		if i == 0 {
			continue
		}

		timeElapsed := s.snapshots[i].timestamp.Sub(s.snapshots[i-1].timestamp)

		intervals = append(intervals, timeElapsed)
	}

	return intervals
}
