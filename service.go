package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"
	"time"
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

func (sg *serviceGenerator) updateMetadata() {
	sg.genTestMaps()
	sg.genTestStructs()
}

func (sg *serviceGenerator) write() error {
	sg.writeFile(".", "srv", "server")
	sg.writeFile("transport", "tp", "transport")
	sg.writeFile("jsonrest", "ep", "endpoint")
	sg.writeFile("conv", "cnv", "conv")
	sg.writeFile("service", "svc", "service")
	err := sg.writeFile(".", "_test", "test")
	return err
}

func (mg *serviceGenerator) writeFile(dir, sufix, template string) error {
	md := mg.Meta
	n := fmt.Sprintf("%s%s.go", md.SingularLowercase, sufix)
	f := filepath.Join(md.PackageDir, "pkg", md.ServicePkgPath, dir, n)

	log.Printf("%s file: %s\n", strings.Title(strings.ToLower(n)), f)

	w, err := getFileWriter(f, mg.force)
	if err != nil {
		return err
	}
	defer w.Close()

	t, err := mg.template(template)
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

// ----------------------------------------------------------------------------
func (sg *serviceGenerator) genTestMaps() {
	md := sg.Meta
	props := md.ClientPropDefs
	end := ",\n"

	// Maps
	var crt bytes.Buffer
	var upd bytes.Buffer
	var smpl1 bytes.Buffer
	var smpl2 bytes.Buffer

	//pc := md.SingularPascalCase
	cc := md.SingularCamelCase

	// Create
	crt.WriteString(fmt.Sprintf("\tcreate%sValid = map[string]interface{}{\n", cc))

	for i, prop := range props {
		if prop.IsEmbedded || prop.IsBackendOnly {
			continue
		}

		prop := props[i]
		var line string
		varName := prop.SingularCamelCase
		// Create
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		crt.WriteString(fmt.Sprintf("%s", line))

		// Update
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		upd.WriteString(fmt.Sprintf("%s", line))

		// Sample 1
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		smpl1.WriteString(fmt.Sprintf("%s", line))

		// Sample 2
		line = fmt.Sprintf("\t\t%s : %s%s", varName, sampleVal(prop), end)
		smpl2.WriteString(fmt.Sprintf("%s", line))
	}

	// Create
	crt.WriteString("\t}")

	md.SvcTestCreateMap = crt.String()
	md.SvcTestUpdateMap = upd.String()
	md.SvcTestSample1Map = smpl1.String()
	md.SvcTestSample2Map = smpl2.String()
}

func sampleFormattedVal(prop propDef) string {
	return formatVal(sampleVal(prop), prop)
}

func sampleVal(prop propDef) string {
	if prop.isKeyType() {
		u := generateUUIDString()
		return fmt.Sprintf("\"%s\"", u)
	} else if prop.isUUIDType() {
		u := generateUUIDString()
		return fmt.Sprintf("\"%s\"", u)
	} else if prop.isTextType() {
		return fmt.Sprintf("\"%s\"", sampleString(prop.Name, 4, 8))
	} else if prop.isIntType() {
		return fmt.Sprintf("%d", sampleInt(2))
	} else if prop.isDecimalType() {
		return fmt.Sprintf("%f", sampleDecimal(4))
	} else if prop.isBooleanType() {
		return fmt.Sprintf("%t", bg.SampleBool())
	} else if prop.isTimeType() {
		return fmt.Sprintf("\"%s\"", time.Now().Format(time.RFC3339))
	}
	return fmt.Sprintf("\"%s\"", "")
}

func formatVal(val interface{}, prop propDef) string {
	if prop.isKeyType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if prop.isUUIDType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if prop.isTextType() {
		return fmt.Sprintf("\"%s\"", val)
	} else if prop.isIntType() {
		return fmt.Sprintf("%d", val)
	} else if prop.isDecimalType() {
		return fmt.Sprintf("%f", val)
	} else if prop.isBooleanType() {
		return fmt.Sprintf("%t", val)
	} else if prop.isTimeType() {
		return fmt.Sprintf("\"%s\"", val)
	}
	return fmt.Sprintf("\"%s\"", val)
}

func (sg *serviceGenerator) genTestStructs() {
	md := sg.Meta
	props := md.ClientPropDefs

	//l := len(props) - 1
	end := ",\n"

	// Maps
	var crt bytes.Buffer
	var upd bytes.Buffer
	var smpl1 bytes.Buffer
	var smpl2 bytes.Buffer

	pc := md.SingularPascalCase
	cc := md.SingularCamelCase

	// Create
	crt.WriteString(fmt.Sprintf("\treq := Create%sReq{\n", pc))
	crt.WriteString(fmt.Sprintf("\t\t%s{\n", pc))

	// Update
	upd.WriteString(fmt.Sprintf("\treq := Update%sReq{\n", pc))
	upd.WriteString(fmt.Sprintf("\t\t%s.Identifier{\n", "mod"))
	upd.WriteString(fmt.Sprintf("\t\t\tSlug: %s.Slug.String,\n", cc))
	upd.WriteString(fmt.Sprintf("\t\t},\n"))
	upd.WriteString(fmt.Sprintf("\t\t%s{\n", pc))

	// Sample1
	smpl1.WriteString(fmt.Sprintf("\t%s1 := &model.%s{\n", cc, pc))

	// Sample2
	smpl2.WriteString(fmt.Sprintf("\t%s2 := &model.%s{\n", cc, pc))

	for i, prop := range props {
		if prop.IsBackendOnly {
			fmt.Printf("BEO %+v\n\n", prop)
			continue
		}

		prop := props[i]
		var line string
		field := prop.SingularPascalCase
		key := prop.Name
		stm := prop.SafeTypeMaker

		// Create
		line = fmt.Sprintf("\t\t\t%s:\tcreate%sValid[\"%s\"]%s", field, pc, key, end)
		crt.WriteString(fmt.Sprintf("%s", line))

		// Update
		line = fmt.Sprintf("\t\t\t%s:\tupdate%sValid[\"%s\"]%s", field, pc, key, end)
		upd.WriteString(fmt.Sprintf("%s", line))

		// Sample 1
		line = fmt.Sprintf("\t\t\t%s:\t%s(sample%s1[\"%s\"]) %s", field, stm, pc, key, end)
		smpl1.WriteString(fmt.Sprintf("%s", line))

		// Sample 2
		line = fmt.Sprintf("\t\t\t%s:\t%s(sample%s2[\"%s\"]) %s", field, stm, pc, key, end)
		smpl2.WriteString(fmt.Sprintf("%s", line))
	}

	// Create
	crt.WriteString("\t\t},\n")
	crt.WriteString("\t}")

	// Update
	upd.WriteString("\t\t},\n")
	upd.WriteString("\t}\n")

	// Sample1
	smpl1.WriteString("\t}\n")

	// Sample2
	smpl2.WriteString("\t}\n")

	md.SvcTestCreateStruct = crt.String()
	md.SvcTestUpdateStruct = upd.String()
	md.SvcTestSample1Struct = smpl1.String()
	md.SvcTestSample2Struct = smpl2.String()
}
