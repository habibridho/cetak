package cetak

import (
	"archive/zip"
	"bytes"
	"io"
	"text/template"
)

// Docx is the interface used to interact with cetak's doc generator
type Docx interface {
	// Generate the document using given data. This will write the result to internal buffer.
	Generate(data interface{}) error
}

type docx struct {
	tpl    *template.Template
	buffer bytes.Buffer
}

// New creates a new Docx object. It will return error if template file could not be opened.
func New(templatePath string) (Docx, error) {
	reader, err := zip.OpenReader(templatePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var tplBuffer bytes.Buffer
	for _, f := range reader.File {
		if f.Name == "word/document.xml" {
			contentReader, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer contentReader.Close()

			_, err = io.Copy(&tplBuffer, contentReader)
			if err != nil {
				return nil, err
			}
		}
	}

	tpl, err := template.New("template").Parse(tplBuffer.String())
	if err != nil {
		return nil, err
	}

	return &docx{
		tpl: tpl,
	}, nil
}

func (d *docx) Generate(data interface{}) error {
	err := d.tpl.Execute(&d.buffer, data)
	return err
}
