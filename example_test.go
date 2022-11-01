package extemplate_test

import (
	"html/template"
	"os"
	"strings"

	"github.com/DHBosworth/extemplate"
)

func ExampleExtemplate_ParseDir() {
	xt := extemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})
	_ = xt.ParseDir("examples", ".tmpl")
	_ = xt.ExecuteTemplate(os.Stdout, "child.tmpl", nil)
	/* Output: Hello from child.tmpl
	Hello from partials/question.tmpl */
}
