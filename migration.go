package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

func (g *gen) genMigration() {
	md := g.meta
	md.genMigStatements()
	err := g.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return
	}
	log.Println("Done")
}

func (md *metadata) genMigStatements() {
	md.genCreateStatement()
	md.genDropStatement()
	md.genFKAlterStatement()
}

func (md *metadata) genCreateStatement() {
	props := md.PropDefs
	var createSQL bytes.Buffer
	createSQL.WriteString(fmt.Sprintf("CREATE TABLE %s\n(", md.PluralSnakeCase))
	last := len(props) - 1
	for i := range props {
		prop := props[i]
		var ending string
		if i < last {
			if prop.SQLModifier != "" {
				ending = fmt.Sprintf("%s %s,\n", prop.SQLType, prop.SQLModifier)
			} else {
				ending = fmt.Sprintf("%s,\n", prop.SQLType)
			}
		} else {
			if prop.SQLModifier != "" {
				ending = fmt.Sprintf("%s %s\n", prop.SQLType, prop.SQLModifier)
			} else {
				ending = fmt.Sprintf("%s\n", prop.SQLType)
			}
		}
		createSQL.WriteString(fmt.Sprintf("\t %s %s", prop.SQLColumn, ending))
	}
	createSQL.WriteString(");")
	md.CreateStatement = createSQL.String()
}

func (md *metadata) genDropStatement() {
	var dropSQL bytes.Buffer
	dropSQL.WriteString(fmt.Sprintf("DROP TABLE %s CASCADE;", md.PluralSnakeCase))
	md.DropStatement = dropSQL.String()
}

func (md *metadata) genFKAlterStatement() {
	props := md.NonVirtualPropDefs
	for i := range props {
		prop := props[i]
		var alterSQL bytes.Buffer
		if prop.Ref.FKName != "" {
			var a bytes.Buffer
			a.WriteString(fmt.Sprintf("ALTER TABLE %s\n", md.PluralSnakeCase))
			a.WriteString(fmt.Sprintf("\tADD CONSTRAINT %s\n", prop.Ref.FKName))
			a.WriteString(fmt.Sprintf("\tFOREIGN KEY (%s)\n", prop.SQLColumn))
			a.WriteString(fmt.Sprintf("\tREFERENCES %s\n", prop.Ref.TrgTable))
			a.WriteString("\tON DELETE CASCADE;\n")
			alterSQL.WriteString(a.String())
			md.AlterStatement = append(md.AlterStatement, alterSQL.String())
		}
	}
}

func (g *gen) write() error {
	md := g.meta
	//n := fmt.Sprintf("%screatetable%s.go", "00000", md.PluralLowercase)
	n := fmt.Sprintf("%screatetable%s.go", newMigrationPrefix(), md.PluralLowercase)
	f := filepath.Join(md.PackageDir, "internal", "migration", n)
	log.Printf("Migration file: %s\n", f)

	w, err := getFileWriter(f, g.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t := template.Must(template.New("template").Parse(migrationTempl))
	return t.Execute(w, md)
}

var migrationTempl = `
package migration

import "log"

// Create{{- .PluralPascalCase -}}Table migration
func (m *mig) Create{{- .PluralPascalCase -}}Table() error {
	tx := m.GetTx()

	st := ` + "`" + `{{- .CreateStatement -}}` + "`" + `

	{{- "\n" -}}

	_, err = tx.Exec(st)
	if err != nil {
		return err
	}

	{{- "\n" -}}

	{{ range $key2, $sqlString := .AlterStatement }}
  st = ` + "`" + `{{- $sqlString -}}` + "`" + `

	{{- "\n" -}}

	_, err = tx.Exec(st)
	if err != nil {
		return err
	}
	{{ end }}

	return nil
}

// Drop{{- .PluralPascalCase -}}Table rollback
func (m *mig) Drop{{- .PluralPascalCase -}}Table() error {
	tx := m.GetTx()

	st := ` + "`" + `{{- .DropStatement -}}` + "`" + `

	{{- "\n" -}}
	{{- "\n" -}}

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}
`
