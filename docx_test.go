package cetak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	d := &docx{}
	err := d.Generate(nil)
	assert.NoError(t, err)
}
