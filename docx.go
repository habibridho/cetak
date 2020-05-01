package cetak

import "errors"

// Docx is the interface used to interact with cetak's doc generator
type Docx interface {
	Generate(data map[string]interface{}) error
}

type docx struct{}

func (d *docx) Generate(data map[string]interface{}) error {
	return errors.New("no implementation")
}
