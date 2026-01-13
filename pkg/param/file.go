package param

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"os"
	"path/filepath"
	"strings"
)

type ParamFile struct {
	ParamReader
	filePath string
}

func NewFile(name string, file *os.File, conf *config.Config) *ParamFile {
	paramReader := NewReader(name, file, filepath.Base(file.Name()), conf)
	return &ParamFile{*paramReader, file.Name()}
}

func NewPath(name string, path string, conf *config.Config) Parameter {
	file, err := os.Open(path)
	if err != nil {
		return NewError(name, err)
	}
	return NewFile(name, file, conf)
}

func (pf *ParamFile) Prepare() error {
	file, err := os.Open(pf.filePath)
	if err != nil {
		return err
	}
	pf.reader = file
	return pf.ParamReader.Prepare()
}

func (pf *ParamFile) Values() ([]string, error) {
	err := pf.Prepare()
	return pf.values, err
}

func (pf *ParamFile) String() string {
	return fmt.Sprintf("%s: %s -> %s", pf.name, pf.filePath, strings.Join(pf.values, " "))
}
