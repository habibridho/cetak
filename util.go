package cetak

import (
	"archive/zip"
	"bytes"
	"io"
)

type actions map[string]func(f *zip.File) error

func executeOnDocx(path string, acts actions, defaultAct func(f *zip.File) error) error {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if act, ok := acts[f.Name]; ok {
			if err := act(f); err != nil {
				return err
			}
		} else if defaultAct != nil {
			if err := defaultAct(f); err != nil {
				return err
			}
		}
	}

	return nil
}

func getDocxContentAsString(path string) (string, error) {
	var tplBuffer bytes.Buffer
	var content string
	acts := actions{
		"word/document.xml": func(f *zip.File) error {
			contentReader, err := f.Open()
			if err != nil {
				return err
			}
			defer contentReader.Close()

			_, err = io.Copy(&tplBuffer, contentReader)
			if err != nil {
				return err
			}

			content = tplBuffer.String()
			return nil
		},
	}
	err := executeOnDocx(path, acts, nil)
	if err != nil {
		return "", err
	}

	return content, nil
}
