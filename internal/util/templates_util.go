package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
)

const templatePath = "../../assets/templates/"

func GetContentFromTemplate(fileName, extension string) ([]byte, error) {
	content := make([]byte, 0)
	absoluteTemplatePath, err := filepath.Abs(templatePath)
	if err != nil {
		return nil, err
	}
	templateFilePath := absoluteTemplatePath + fileName + "_template" + extension
	file, err := os.OpenFile(templateFilePath, syscall.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	_, err = file.Read(content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
