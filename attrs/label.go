package attrs

import (
	"github.com/a-h/templ"
	"maps"
)

type Label struct {
	For string
	Global
}

func (i Label) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "for", i.For)
	maps.Copy(attrs, i.Global.Attrs())
	return
}
