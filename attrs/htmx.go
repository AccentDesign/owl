package attrs

import (
	"github.com/a-h/templ"
)

type HX struct {
	Confirm     string
	Get         string
	OnAfterLoad string
	Post        string
	PushUrl     bool
	Select      string
	Swap        string
	Target      string
}

func (hx HX) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "hx-confirm", hx.Confirm)
	addIfNotEmpty(attrs, "hx-get", hx.Get)
	addIfNotEmpty(attrs, "hx-on:htmx-after-on-load", hx.OnAfterLoad)
	addIfNotEmpty(attrs, "hx-post", hx.Post)
	addIfNotEmpty(attrs, "hx-select", hx.Select)
	addIfNotEmpty(attrs, "hx-swap", hx.Swap)
	addIfNotEmpty(attrs, "hx-target", hx.Target)
	if hx.PushUrl {
		attrs["hx-push-url"] = "true"
	}
	return
}
