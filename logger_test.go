package leveledlogger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/BorisPlus/leveledlogger"
)

func TestLogger(t *testing.T) {
	testCases := []struct {
		level          leveledlogger.LogLevel
		debugMessage   string
		infoMessage    string
		warningMessage string
		errorMessage   string
		expectedStack  []string
	}{
		{
			leveledlogger.DEBUG,
			"DEBUG message",
			"INFO message",
			"WARNING message",
			"ERROR message",
			[]string{"DEBUG message", "INFO message", "WARNING message", "ERROR message"},
		},
		{
			leveledlogger.INFO,
			"DEBUG message",
			"INFO message",
			"WARNING message",
			"ERROR message",
			[]string{"INFO message", "WARNING message", "ERROR message"},
		},
		{
			leveledlogger.WARNING,
			"DEBUG message",
			"INFO message",
			"WARNING message",
			"ERROR message",
			[]string{"WARNING message", "ERROR message"},
		},
		{
			leveledlogger.ERROR,
			"DEBUG message",
			"INFO message",
			"WARNING message",
			"ERROR message",
			[]string{"ERROR message"},
		},
	}
	for _, testCase := range testCases {
		outputInto := &bytes.Buffer{}
		logger := leveledlogger.NewLogger(testCase.level, outputInto)
		logger.Debug(testCase.debugMessage)
		logger.Info(testCase.infoMessage)
		logger.Warning(testCase.warningMessage)
		logger.Error(testCase.errorMessage)
		output := outputInto.String()
		for _, expected := range testCase.expectedStack {
			if !strings.Contains(output, expected) {
				t.Errorf("Error logger %s output: %q not contains %s\n", logger.LogLevel(), output, expected)
			}
		}
	}
}

func TestFormatSuffix(t *testing.T) {
	expected := "I am suffix from formatting trmplate +100500\n"
	outputInto := &bytes.Buffer{}
	logger := leveledlogger.NewLogger("INFO", outputInto)
	logger.Info("%s +%d", "I am suffix from formatting trmplate", 100500)
	output := outputInto.String()
	if !strings.HasSuffix(outputInto.String(), expected) {
		t.Errorf("Logger %s output: %q not contains %q\n", logger.LogLevel(), output, expected)
	}
}
