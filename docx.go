// Package cetak helps you generate Word document (.docx) from a template.
package cetak

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"text/template"
)

// Docx is the interface used to interact with cetak's doc generator
//
// Generate generates the document using given data.
// Generate receive data which can be any type that is accepted by text/template package.
// Generate also receive destination path which is the intended path to write the result.
type Docx interface {
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

// Generate generates the document content based on the template set in d object
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
	docxWriter := zip.NewWriter(destFile)

	acts := actions{
		"word/document.xml": func(f *zip.File) error {
			docFile, err := docxWriter.Create(f.Name)
			if err != nil {
				return err
			}
			if _, err = docFile.Write([]byte(resultBuf.String())); err != nil {
				return err
			}
			return nil
		},
	}
	defaultAct := func(f *zip.File) error {
		docFile, err := docxWriter.Create(f.Name)
		if err != nil {
			return err
		}
		tplFileReader, err := f.Open()
		if err != nil {
			return err
		}
		if _, err = io.Copy(docFile, tplFileReader); err != nil {
			return err
		}
		return nil
	}

	if err := executeOnDocx(d.templatePath, acts, defaultAct); err != nil {
		return err
	}
	return docxWriter.Close()
}
