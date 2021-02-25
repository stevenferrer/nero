package template

import (
	"text/template"
)

// ParseTemplate parses the repository template
func ParseTemplate(tmpl string) (*template.Template, error) {
	tmplt, err := template.New("impl.tmpl").
		Funcs(template.FuncMap{
			"type":       typeFunc,
			"zero":       zeroFunc,
			"plural":     pluralFunc,
			"lowerCamel": lowerCamelFunc,
		}).
		Parse(tmpl)

	return tmplt, err
}
