package attrs

import (
	"github.com/a-h/templ"
	"maps"
)

type Attributes interface {
	Attrs() templ.Attributes
}

type Global struct {
	ID     string
	Class  templ.CSSClasses
	Role   string
	Title  string
	Hidden bool
	Actions
	Aria
	HX
}

func (g Global) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "id", g.ID)
	addIfNotEmpty(attrs, "role", g.Role)
	addIfNotEmpty(attrs, "title", g.Title)
	if len(g.Class) > 0 {
		attrs["class"] = g.Class.String()
	}
	if g.Hidden {
		attrs["hidden"] = true
	}
	maps.Copy(attrs, g.Actions.Attrs())
	maps.Copy(attrs, g.Aria.Attrs())
	maps.Copy(attrs, g.HX.Attrs())
	return
}

var Empty = Global{}
