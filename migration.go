package main

import (
	"os"
	"text/template"
)

func migration(g *gen) {
	var d metadata
	t := template.Must(template.New("template").Parse(migrationTempl))
	t.Execute(os.Stdout, d)
}

var migrationTempl = `
package migration

import (
)
`
