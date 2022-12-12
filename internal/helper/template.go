package helper

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

func ExecuteTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(tmplStr)
	if err != nil {
		return "", errors.Wrapf(err, "failed to parse template for %T", data)
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, data); err != nil {
		return "", errors.Wrapf(err, "failed to execute template for %T", data)
	}

	return buf.String(), nil
}
