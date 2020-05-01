package cetak

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	TestContent string
}

func TestGenerate(t *testing.T) {
	mockTemplate, err := template.New("template").Parse("Hello {{.TestContent}}")
	if err != nil {
		t.Fatalf("could not parse mock tempalte")
	}
	d := &docx{
		tpl: mockTemplate,
	}
	data := TestData{
		TestContent: "World!",
	}

	err = d.Generate(data)
	assert.NoError(t, err)

	expected := "Hello World!"
	actual := d.buffer.String()
	assert.Equal(t, expected, actual)
}
