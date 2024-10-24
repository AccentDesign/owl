package main

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/owl/examples"
	"log"
	"net/http"
)

var (
	listen = ":80"
)

func main() {
	serve()
}

func serve() {
	http.Handle("/", templ.Handler(examples.Index()))
	http.Handle("/accordion", templ.Handler(examples.Accordion()))
	http.Handle("/alert", templ.Handler(examples.Alert()))
	http.Handle("/avatar", templ.Handler(examples.Avatar()))
	http.Handle("/badge", templ.Handler(examples.Badge()))
	http.Handle("/breadcrumb", templ.Handler(examples.Breadcrumb()))
	http.Handle("/button", templ.Handler(examples.Button()))
	http.Handle("/card", templ.Handler(examples.Card()))
	http.Handle("/code", templ.Handler(examples.Code()))
	http.Handle("/dropdown-menu", templ.Handler(examples.DropdownMenu()))
	http.Handle("/form", templ.Handler(examples.Form()))
	http.Handle("/switch", templ.Handler(examples.Switch()))
	http.Handle("/table", templ.Handler(examples.Table()))
	http.Handle("/tabs", templ.Handler(examples.Tabs()))
	http.Handle("/tabs/account", templ.Handler(examples.TabContentAccount()))
	http.Handle("/tabs/password", templ.Handler(examples.TabContentPassword()))
	fmt.Printf("listening on %s\n", listen)
	log.Panic(http.ListenAndServe(listen, nil))
}
