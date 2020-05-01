package cetak

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Title   string
	Content string
}

func TestGenerate(t *testing.T) {
	d, err := New("sample_templates/simple-template.docx")
	if err != nil {
		t.Fatalf("could not create docx object: %s", err.Error())
	}

	data := TestData{
		Title:   "Some Title",
		Content: "Some Content",
	}

	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("could not generate uuid: %s", err.Error())
	}
	destination := fmt.Sprintf("%s.docx", id.String())
	err = d.Generate(data, destination)
	assert.NoError(t, err)
	defer os.Remove(destination)

	expected, err := getDocxContentAsString("test_resources/simple-result.docx")
	if err != nil {
		t.Fatalf("could not read test resource file: %s", err.Error())
	}
	actual, err := getDocxContentAsString(destination)
	if err != nil {
		t.Fatalf("could not read test resource file: %s", err.Error())
	}
	assert.Equal(t, expected, actual)
}
