package css

import (
	"context"
	"embed"
	"github.com/a-h/templ"
	"io"
)

var (
	//go:embed *.css
	fs embed.FS
)

// Include dirty include css file for tailwind cli
func Include(filename string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		content, err := fs.ReadFile(filename)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<style type='text/tailwindcss'>"+string(content)+"</style>")
		return err
	})
}
