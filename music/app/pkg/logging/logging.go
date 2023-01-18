package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

const separator = "\n\n"

func New(consoleLevel, path string) (*zap.Logger, error) {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("create/open log file (%s): %w", path, err)
	}

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("get file stats: %w", err)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("%s is directory", info.Name())
	}

	if info.Size() > 0 {
		_, err = file.WriteString(separator)
		if err != nil {
			return nil, fmt.Errorf("write separator to file: %w", err)
		}
	}

	return NewLogger(consoleLevel, os.Stdout, file), nil
}
