package attrs

import (
	"github.com/a-h/templ"
	"maps"
)

type Img struct {
	Src    string
	Alt    string
	Width  string
	Height string
	Global
}

func (i Img) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "src", i.Src)
	addIfNotEmpty(attrs, "alt", i.Alt)
	addIfNotEmpty(attrs, "width", i.Width)
	addIfNotEmpty(attrs, "height", i.Height)
	maps.Copy(attrs, i.Global.Attrs())
	return
}
