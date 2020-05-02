# Cetak
[![GoDoc](https://godoc.org/github.com/habibridho/cetak?status.svg)](https://godoc.org/github.com/habibridho/cetak)
[![Build Status](https://travis-ci.org/habibridho/cetak.svg?branch=master)](https://travis-ci.org/habibridho/cetak)

Golang package that helps you generate Word document (.docx) from a template. It utilises go's `text/template` package behind the scene so you can use the same annotation in your template.

## Usage
To use this package, simply create new Docx object and call Generate with your data dan destination path.
```go
...
d, err := cetak.New("/path/to/template.docx")
if err != nil {
    // handle error
}

type TemplateData struct {
    Title   string
	Content string
}

data := TemplateData{
    Title:   "This is The Title",
    Content: "Hello World!",
}

err = d.Generate(data, "/path/to/result.docx")
if err != nil {
    // handle error
}
...
```
In the above sample, we assume to have a template ready in `/path/to/template.docx`. We will inject `data` which has an arbitrary `TemplateData` type. Note that the `data` type can be anything where the fields correspond to annotations inside the template. In this sample, `template.docx` has `{{.Title}}` and `{{.Content}}` inside of it. The result will be stored in the given result path, ie `/path/to/result.docx`.

## License
This package is licensed under [MIT License](LICENSE).
