package fileprocessor

import (
	"bufio"
	"os"
)

type FileProcessor struct {
	in  string
	out *os.File
}

func New(inputFilePath string, output *os.File) *FileProcessor {
	return &FileProcessor{
		in:  inputFilePath,
		out: output,
	}
}

func (fp *FileProcessor) Process(processFunc func(string) []byte) error {
	f, err := os.Open(fp.in)
	if err != nil {
		return err
	}
	defer f.Close()

	fr := bufio.NewScanner(f)
	fw := bufio.NewWriter(fp.out)

	for fr.Scan() {
		fw.Write(processFunc(fr.Text()))
	}

	return fw.Flush()
}
