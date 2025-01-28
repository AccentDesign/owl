// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package examples

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

var tableStyles = []string{"typography.css", "table.css"}

func Table() templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"max-w-4xl\"><table class=\"owl-table\"><caption class=\"owl-table-caption\">A list of your recent invoices.</caption> <thead class=\"owl-table-header\"><tr class=\"owl-table-row\"><th class=\"owl-table-head\">Invoice</th><th class=\"owl-table-head\">Status</th><th class=\"owl-table-head\">Method</th><th class=\"owl-table-head text-right\">Amount</th></tr></thead> <tbody class=\"owl-table-body\"><tr class=\"owl-table-row\"><td class=\"owl-table-cell\">INV001</td><td class=\"owl-table-cell\">Paid</td><td class=\"owl-table-cell\">Credit Card</td><td class=\"owl-table-cell text-right\">£250.00</td></tr><tr class=\"owl-table-row\"><td class=\"owl-table-cell\">INV002</td><td class=\"owl-table-cell\">Paid</td><td class=\"owl-table-cell\">Credit Card</td><td class=\"owl-table-cell text-right\">£500.00</td></tr><tr class=\"owl-table-row\"><td class=\"owl-table-cell\">INV003</td><td class=\"owl-table-cell\">Paid</td><td class=\"owl-table-cell\">Credit Card</td><td class=\"owl-table-cell text-right\">£250.00</td></tr></tbody><tfoot class=\"owl-table-footer\"><tr class=\"owl-table-row\"><td class=\"owl-table-cell\">Total</td><td class=\"owl-table-cell\"></td><td class=\"owl-table-cell\"></td><td class=\"owl-table-cell text-right\">£1,000.00</td></tr></tfoot></table></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Base(tableStyles, "table").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
