package main

import (
	"context"
	"github.com/a-h/templ"
	"github.com/accentdesign/owl/examples"
	"log"
	"net/http"
	"os"
)

func main() {
	build()
	serve()
}

func build() {
	createHtml("docs/index.html", examples.Index())
	createHtml("docs/accordion.html", examples.Accordion())
	createHtml("docs/alert.html", examples.Alert())
	createHtml("docs/alert-dialog.html", examples.AlertDialog())
	createHtml("docs/alert-dialog-open.html", examples.AlertDialogOpen())
	createHtml("docs/alert-dialog-closed.html", examples.AlertDialogClosed())
	createHtml("docs/avatar.html", examples.Avatar())
	createHtml("docs/badge.html", examples.Badge())
	createHtml("docs/breadcrumb.html", examples.Breadcrumb())
	createHtml("docs/button.html", examples.Button())
	createHtml("docs/card.html", examples.Card())
	createHtml("docs/code.html", examples.Code())
	createHtml("docs/datepicker.html", examples.Datepicker())
	createHtml("docs/dialog.html", examples.Dialog())
	createHtml("docs/dialog-open.html", examples.DialogOpen())
	createHtml("docs/dialog-closed.html", examples.DialogClosed())
	createHtml("docs/dropdown-menu.html", examples.DropdownMenu())
	createHtml("docs/dropdown-menu-open.html", examples.DropdownMenuComponent(true))
	createHtml("docs/dropdown-menu-closed.html", examples.DropdownMenuComponent(false))
	createHtml("docs/form.html", examples.Form())
	createHtml("docs/pagination.html", examples.Pagination())
	createHtml("docs/separator.html", examples.Separator())
	createHtml("docs/switch.html", examples.Switch())
	createHtml("docs/table.html", examples.Table())
	createHtml("docs/tabs.html", examples.Tabs())
	createHtml("docs/tooltip.html", examples.Tooltip())
	createHtml("docs/tabs-account.html", examples.TabContentAccount())
	createHtml("docs/tabs-password.html", examples.TabContentPassword())
	createHtml("docs/toast.html", examples.Toast())
	createHtml("docs/toaster.html", examples.Toaster())
	createHtml("docs/toast-message.html", examples.ToastMessage(true))
	createHtml("docs/toast-message-confirm.html", examples.ToastMessage(false))
}

func serve() {
	http.Handle("/", http.FileServer(http.Dir("docs")))
	log.Println("Listening on :80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func createHtml(filePath string, component templ.Component) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = component.Render(context.Background(), f)
	if err != nil {
		log.Fatal(err)
	}
}
