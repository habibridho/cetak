package cetak

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Title             string
	Content           string
	List              []string
	UndefinedPipeline string
}

type TestCase struct {
	name         string
	templatePath string
	expectedPath string
	data         TestData
}

func TestGenerate(t *testing.T) {
	cases := []TestCase{
		{
			name:         "test with simple template",
			templatePath: "sample_templates/simple-template.docx",
			expectedPath: "test_resources/simple-result.docx",
			data: TestData{
				Title:   "Some Title",
				Content: "Some Content",
			},
		},
		{
			name:         "test with advanced template",
			templatePath: "sample_templates/advanced-template.docx",
			expectedPath: "test_resources/advanced-result.docx",
			data: TestData{
				Title:   "Some Title",
				Content: "Some Content",
				List:    []string{"Apple", "Banana", "Cherry", "Durian"},
			},
		},
	}

	for _, c := range cases {
		log.Println(c.name)
		d, err := New(c.templatePath)
		if err != nil {
			t.Fatalf("could not create docx object: %s", err.Error())
		}

		id, err := uuid.NewRandom()
		if err != nil {
			t.Fatalf("could not generate uuid: %s", err.Error())
		}
		destination := fmt.Sprintf("%s.docx", id.String())
		err = d.Generate(c.data, destination)
		assert.NoError(t, err)
		defer os.Remove(destination)

		expected, err := getDocxContentAsString(c.expectedPath)
		if err != nil {
			t.Fatalf("could not read test resource file: %s", err.Error())
		}
		actual, err := getDocxContentAsString(destination)
		if err != nil {
			t.Fatalf("could not read test resource file: %s", err.Error())
		}
		assert.Equal(t, expected, actual)
	}
}

func TestTemplate(t *testing.T) {
	tpl := template.New("")
	d := docx{
		tpl: tpl,
	}

	result := d.Template()
	assert.Equal(t, tpl, result)
}

func TestSetTemplate(t *testing.T) {
	tpl := template.New("")
	d := docx{
		tpl: tpl,
	}

	newTpl, err := tpl.Clone()
	if err != nil {
		t.Fatalf("could not clone template: %s", err.Error())
	}

	d.SetTemplate(newTpl)
	result := d.Template()
	assert.Equal(t, newTpl, result)
}
