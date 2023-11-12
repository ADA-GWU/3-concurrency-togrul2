package arg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RunMode string

const (
	SingleThreaded RunMode = "S"
	MultiThreaded  RunMode = "M"
)

func stringToRunMode(value string) (RunMode, error) {
	switch value = strings.ToUpper(value); value {
	case "S", "M":
		return RunMode(value), nil
	default:
		return "", fmt.Errorf("invalid option for mode argument")
	}
}

// Parses command line arguments and returns file name, size of a
// single square and mode in which app should proceed data.
func ParseArguments() (string, int, RunMode, error) {
	if len(os.Args) != 4 {
		return "", 0, "", fmt.Errorf("invalid number of arguments")
	}

	fileName := os.Args[1]
	// Parse square size argument.
	squareSize, sizeParseErr := strconv.Atoi(os.Args[2])
	if sizeParseErr != nil {
		return "", 0, "", fmt.Errorf("provide a valid integer for square size parameter")
	}

	// Parse mode argument.
	mode, runModeParseErr := stringToRunMode(os.Args[3])
	if runModeParseErr != nil {
		return "", 0, "", runModeParseErr
	}

	return fileName, squareSize, mode, nil
}
