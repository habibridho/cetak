package cetak

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"text/template"
)

// Docx is the interface used to interact with cetak's doc generator
type Docx interface {
	// Generate the document using given data.
	// The result will be written to the given destination path.
	Generate(data interface{}, destination string) error
}

type docx struct {
	tpl          *template.Template
	templatePath string
}

// New creates a new Docx object. It will return error if template file could not be opened.
func New(templatePath string) (Docx, error) {
	tplString, err := getDocxContentAsString(templatePath)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("template").Parse(tplString)
	if err != nil {
		return nil, err
	}

	return &docx{
		tpl:          tpl,
		templatePath: templatePath,
	}, nil
}

func (d *docx) Generate(data interface{}, destination string) error {
	var resultBuf bytes.Buffer
	err := d.tpl.Execute(&resultBuf, data)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}

	docxReader, err := zip.OpenReader(d.templatePath)
	if err != nil {
		return err
	}
	defer docxReader.Close()

	docxWriter := zip.NewWriter(destFile)
	for _, tplFile := range docxReader.File {
		f, err := docxWriter.Create(tplFile.Name)
		if err != nil {
			return err
		}

		if tplFile.Name == "word/document.xml" {
			if _, err = f.Write([]byte(resultBuf.String())); err != nil {
				return err
			}
		} else {
			tplFileReader, err := tplFile.Open()
			if err != nil {
				return err
			}
			if _, err = io.Copy(f, tplFileReader); err != nil {
				return err
			}
		}
	}
	return docxWriter.Close()
}
