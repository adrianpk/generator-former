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

	t, err := mg.template()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (mg *modelGenerator) template() (*template.Template, error) {
	res, err := Asset("assets/templates/model.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
