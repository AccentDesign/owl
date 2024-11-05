// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package examples

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/accentdesign/owl/css"

func Base() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@2.0.3\" integrity=\"sha384-0895/pl2MU10Hqc6jd4RvrthNlDiE9U1tWmX7WRESftEDRosgxNsQG/Ze9YMRzHq\" crossorigin=\"anonymous\"></script><script src=\"https://cdn.tailwindcss.com\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = css.Include("typography.css").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"container flex gap-4 mx-auto my-10 antialiased\"><aside class=\"w-64\"><h4 class=\"owl-h4 mb-4\">components</h4><ul class=\"list-disc list-inside mb-4 text-gray-500\"><li><a href=\"accordion.html\">accordion</a></li><li><a href=\"alert.html\">alert</a></li><li><a href=\"alert-dialog.html\">alert dialog</a></li><li><a href=\"avatar.html\">avatar</a></li><li><a href=\"badge.html\">badge</a></li><li><a href=\"breadcrumb.html\">breadcrumb</a></li><li><a href=\"button.html\">button</a></li><li><a href=\"card.html\">card</a></li><li><a href=\"code.html\">code</a></li><li><a href=\"dialog.html\">dialog</a></li><li><a href=\"dropdown-menu.html\">dropdown menu</a></li><li><a href=\"form.html\">form</a></li><li><a href=\"switch.html\">switch</a></li><li><a href=\"table.html\">table</a></li><li><a href=\"tabs.html\">tabs</a></li></ul><h4 class=\"owl-h4 mb-4\">links</h4><ul class=\"list-disc list-inside text-gray-500\"><li><a href=\"index.html\">home</a></li><li><a target=\"_blank\" href=\"https://github.com/AccentDesign/owl\">github</a></li><li><a target=\"_blank\" href=\"https://go.dev\">golang</a></li><li><a target=\"_blank\" href=\"https://templ.guide\">templ</a></li><li><a target=\"_blank\" href=\"https://github.com/air-verse/air\">air</a></li><li><a target=\"_blank\" href=\"https://icones.js.org/collection/lucide\">icônes</a></li><li><a target=\"_blank\" href=\"https://tailwindcss.com\">tailwind</a></li><li><a target=\"_blank\" href=\"https://ui.shadcn.com/docs\">shadcn</a></li><li><a target=\"_blank\" href=\"https://htmx.org/\">htmx</a></li></ul></aside><main class=\"w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
