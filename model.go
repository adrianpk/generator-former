package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"
)

type (
	modelGenerator struct {
		Meta  *metadata
		force bool
	}
)

func (g *gen) genModel() {
	md := g.Meta
	mg := modelGenerator{
		Meta:  md,
		force: g.Force,
	}

	mg.updateMetadata()

	err := mg.write()
	if err != nil {
		log.Printf("Not done: %s\n", err.Error())
		return
	}
	log.Println("Done!")
}

func (mg *modelGenerator) updateMetadata() {
	mg.genMatchCondition()
}

func (mg *modelGenerator) genMatchCondition() {
	md := mg.Meta
	props := md.ClientPropDefs
	l := len(props) - 1
	var mcond bytes.Buffer
	mcond.WriteString(fmt.Sprintf("%s.Identification.Match(tc.Identification) &&\n", md.SingularCamelCase))
	for i, prop := range props {
		if prop.Name != "ID" {
			prop := props[i]
			var line string
			if l == i {
				line = fmt.Sprintf("\t\t%s.%s == tc.%s\n", md.SingularCamelCase, prop.Name, prop.Name)
			} else {
				line = fmt.Sprintf("\t\t%s.%s == tc.%s &&\n", md.SingularCamelCase, prop.Name, prop.Name)
			}
			mcond.WriteString(line)
		}
	}
	md.ModelMatchCond = mcond.String()
}

func (mg *modelGenerator) write() error {
	md := mg.Meta
	n := fmt.Sprintf("%s.go", md.SingularLowercase)
	f := filepath.Join(md.PackageDir, "internal", "model", n)
	log.Printf("Model file: %s\n", f)

	w, err := getFileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t := template.Must(template.New("template").Parse(modelTempl))
	return t.Execute(w, md)
}

var modelTempl = `
package model

import (
	"database/sql"

	"github.com/lib/pq"
	"gitlab.com/mikrowezel/backend/db"
	m "gitlab.com/mikrowezel/backend/model"
	"golang.org/x/crypto/bcrypt"
)

type (
	// {{.SingularPascalCase}} model
	{{.SingularPascalCase}} struct {
		m.Identification
		{{- range $key, $prop := .PropDefs}}
		{{- if not $prop.IsEmbedded}}
		{{$prop.Name}} {{$prop.SafeType}} ` + "`" + `db:"{{$prop.SQLColumn}}" json:"{{$prop.SingularCamelCase}},omitempty"` + "`" + `
		{{- end}}
		{{- end}}
    m.Audit
	}
)

// SetCreateValues for model.
func ({{.SingularCamelCase}} *{{.SingularPascalCase}}) SetCreateValues() error {
	{{.SingularCamelCase}}Name := {{.SingularCamelCase}}.Name.String
	{{.SingularCamelCase}}.Identification.SetCreateValues(accountName)
	{{.SingularCamelCase}}.Audit.SetCreateValues()
	return nil
}

// SetUpdateValues for model.
func ({{.SingularCamelCase}} *{{.SingularPascalCase}}) SetUpdateValues() error {
	{{.SingularCamelCase}}.Audit.SetUpdateValues()
	return nil
}

// Match condition for model.
func ({{.SingularCamelCase}} *{{.SingularPascalCase}}) Match(tc *{{.SingularPascalCase}}) bool {
	r := {{.ModelMatchCond}}
	return r
}`
