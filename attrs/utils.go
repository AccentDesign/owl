package attrs

import "github.com/a-h/templ"

func addIfNotEmpty(attrs templ.Attributes, key, value string) {
	if value != "" {
		attrs[key] = value
	}
}
