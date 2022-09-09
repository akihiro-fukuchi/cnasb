package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akihiro-fukuchi/cnasb/pkg/fileprocessor"
)

type LogParser interface {
	Parse() error
}

type NginxTSVLogParser struct {
	fp          *fileprocessor.FileProcessor
	logFilePath string
}

func NewNginxTSVLogParser(logFilePath string) *NginxTSVLogParser {
	return &NginxTSVLogParser{
		fp: fileprocessor.New(logFilePath, os.Stdout),
	}
}

func (p *NginxTSVLogParser) Parse() error {
	return p.fp.Process(func(line string) []byte {
		cs := strings.Split(line, "\t")
		return []byte(fmt.Sprintf("path: %s, status: %s\n", cs[3], cs[5]))
	})
}

func main() {
	logParser := NewNginxTSVLogParser(os.Args[1])

	if err := logParser.Parse(); err != nil {
		fmt.Printf("failed to parse %+v", err)
		os.Exit(1)
	}
}
