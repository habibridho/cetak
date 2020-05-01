package cetak

import (
	"bytes"
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

func (d *docx) Generate(data interface{}) error {
	err := d.tpl.Execute(&d.buffer, data)
	return err
}
