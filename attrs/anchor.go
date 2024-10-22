package attrs

import (
	"github.com/a-h/templ"
	"maps"
)

type Anchor struct {
	Href string
	Global
}

func (i Anchor) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "href", i.Href)
	maps.Copy(attrs, i.Global.Attrs())
	return
}
