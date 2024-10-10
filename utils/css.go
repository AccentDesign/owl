package utils

import (
	"context"
	"github.com/a-h/templ"
	"io"
	"os"
)

// InlineCss dirty include css file for tailwind cli
func InlineCss(filename string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		content, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<style type='text/tailwindcss'>"+string(content)+"</style>")
		return err
	})
}
