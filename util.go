package cetak

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
)

func getDocxContentAsString(templatePath string) (string, error) {
	reader, err := zip.OpenReader(templatePath)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var tplBuffer bytes.Buffer
	for _, f := range reader.File {
		if f.Name == "word/document.xml" {
			contentReader, err := f.Open()
			if err != nil {
				return "", err
			}
			defer contentReader.Close()

			_, err = io.Copy(&tplBuffer, contentReader)
			if err != nil {
				return "", err
			}

			return tplBuffer.String(), nil
		}
	}

	return "", errors.New("docx content not found")
}
