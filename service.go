package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

type (
	serviceGenerator struct {
		Meta  *metadata
		force bool
	}
)

func (g *gen) genService() {
	md := g.Meta
	mg := serviceGenerator{
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

func (mg *serviceGenerator) updateMetadata() {
	mg.write()
}

func (mg *serviceGenerator) write() error {
	mg.writeFile("service")
	mg.writeFile("transport")
	err := mg.writeFile("endpoint")
	//mg.writeFile("helper")
	//mg.writeiFile("test")
	return err
}

func (mg *serviceGenerator) writeFile(name string) error {
	md := mg.Meta
	n := fmt.Sprintf("%s%s.go", md.SingularLowercase, name)
	f := filepath.Join(md.PackageDir, "pkg", md.ServicePkgPath, n)
	log.Printf("%s file: %s\n", strings.Title(strings.ToLower(name)), f)

	w, err := getFileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := mg.template(name)
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (sg *serviceGenerator) template(name string) (*template.Template, error) {
	path := fmt.Sprintf("assets/templates/%s.tmpl", name)
	res, err := Asset(path)
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
