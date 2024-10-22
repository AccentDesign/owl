package attrs

import (
	"github.com/a-h/templ"
)

type Aria struct {
	// Provides an accessible name for an element when the native HTML labeling is insufficient.
	Label string
	// Points to another element that serves as the label for this element.
	LabelledBy string
	// Points to another element that provides a description for the element.
	DescribedBy string
	// Hides elements from screen readers, often used for decorative content or elements not meant for interaction.
	Hidden bool
	// Indicates whether a collapsible element is expanded or collapsed.
	Expanded bool
	// Points to an element that the current element controls, such as a menu or tab content.
	Controls string
	// Marks an element as disabled, signaling that it is inactive and not focusable.
	Disabled bool
	// Indicates the selection state for elements within a selectable group, like tabs or list items.
	Selected bool
	// Defines the role of an element (e.g., alert, button, dialog), especially for custom components not native to HTML.
	Role string
	// Indicates that an element is a modal window, trapping focus within it until dismissed.
	Modal bool
	// Indicates the current item within a set of related items, such as the current page in a pagination list.
	Current string
}

func (a Aria) Attrs() (attrs templ.Attributes) {
	attrs = make(templ.Attributes)
	addIfNotEmpty(attrs, "aria-label", a.Label)
	addIfNotEmpty(attrs, "aria-labelledby", a.LabelledBy)
	addIfNotEmpty(attrs, "aria-describedby", a.DescribedBy)
	addIfNotEmpty(attrs, "aria-controls", a.Controls)
	addIfNotEmpty(attrs, "aria-role", a.Role)
	addIfNotEmpty(attrs, "aria-current", a.Current)
	if a.Hidden {
		attrs["aria-hidden"] = "true"
	}
	if a.Expanded {
		attrs["aria-expanded"] = "true"
	}
	if a.Disabled {
		attrs["aria-disabled"] = "true"
	}
	if a.Selected {
		attrs["aria-selected"] = "true"
	}
	if a.Modal {
		attrs["aria-modal"] = "true"
	}
	return
}
