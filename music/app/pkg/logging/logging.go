package logging

import (
	"log"
	"os"
)

const separator = "\n\n"

func Init(consoleLevel, path string) {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("create/open log file (%s): %w", path, err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal("get file stats: %w", err)
	}

	if info.IsDir() {
		log.Fatal("%s is directory", info.Name())
	}

	if info.Size() > 0 {
		_, err = file.WriteString(separator)
		if err != nil {
			log.Fatal("write separator to file: %w", err)
		}
	}

	initLogger(consoleLevel, os.Stdout, file)

}
