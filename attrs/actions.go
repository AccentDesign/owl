package attrs

import (
	"github.com/a-h/templ"
)

type Actions struct {
	OnClick string
	OnError string
}

func (a Actions) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "onclick", a.OnClick)
	addIfNotEmpty(attrs, "onerror", a.OnError)
	return
}
