package css

import (
	"bytes"
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
func Include(filenames ...string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var buffer bytes.Buffer
		for _, filename := range filenames {
			content, err := fs.ReadFile(filename)
			if err != nil {
				return err
			}
			buffer.Write(content)
			buffer.WriteString("\n")
		}
		_, err := io.WriteString(w, "<style type='text/tailwindcss'>"+buffer.String()+"</style>")
		return err
	})
}
