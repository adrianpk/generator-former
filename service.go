package main

import (
	"fmt"
	"log"
	"path/filepath"
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
	mg.writeServer()
	err := mg.writeTransport()
	//mg.writeService()
	//mg.writeHelper()
	//mg.writeTest()
	return err
}

func (mg *serviceGenerator) writeServer() error {
	md := mg.Meta
	n := fmt.Sprintf("%sserver.go", md.SingularLowercase)
	f := filepath.Join(md.PackageDir, "pkg", md.ServicePkgPath, n)
	log.Printf("Service file: %s\n", f)

	w, err := getFileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := mg.serverTemplate()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (sg *serviceGenerator) serverTemplate() (*template.Template, error) {
	res, err := Asset("assets/templates/server.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}

func (mg *serviceGenerator) writeTransport() error {
	md := mg.Meta
	n := fmt.Sprintf("%stransport.go", md.SingularLowercase)
	f := filepath.Join(md.PackageDir, "pkg", md.ServicePkgPath, n)
	log.Printf("Service file: %s\n", f)

	w, err := getFileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := mg.transportTemplate()
	if err != nil {
		return err
	}

	return t.Execute(w, md)
}

func (sg *serviceGenerator) transportTemplate() (*template.Template, error) {
	res, err := Asset("assets/templates/transport.tmpl")
	if err != nil {
		return nil, err
	}
	t := template.New("template")
	return t.Parse(string(res))
}
