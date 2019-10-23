package main

//go:generate rm -rf resources.go
//go:generate go-bindata -pkg main -o resources.go assets/templates/...

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	//"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type (
	gen struct {
		Cmd    string
		Target string
		Using  string

		Force bool

		data []byte
		Meta *metadata
	}
)

var (
	// Commands
	generateCmd = "generate"
	helpCmd     = "help"
	// Target
	handlerTgt   = "handler"
	migrationTgt = "migration"
	modelTgt     = "model"
	repoTgt      = "repo"
	restcltTgt   = "restctl"
	testTgt      = "test"
	allTgt       = "all"

	commands = []string{generateCmd, helpCmd}
	targets  = []string{handlerTgt, migrationTgt, modelTgt, repoTgt, restcltTgt, testTgt, allTgt}
)

func main() {
	g := &gen{}

	args := os.Args[1:]

	flag.BoolVar(&g.Force, "force", true, "Overwrite output files.")
	flag.Parse()

	err := g.setup(args)
	if err != nil {
		log.Println(errMsg(err))
	}

	err = g.genMeta()
	if err != nil {
		log.Println(errMsg(err))
	}

	if g.targetIs(handlerTgt) {
		//g.genHandler()
		//return
	}

	if g.targetIs(migrationTgt) {
		//g.genMigration()
	}

	if g.targetIs(modelTgt) {
		g.genModel()
		return
	}

	if g.targetIs(repoTgt) {
		panic("Not implemented yet.")
	}

	if g.targetIs(restcltTgt) {
		panic("Not implemented yet.")
	}

	if g.targetIs(testTgt) {
		panic("Not implemented yet.")
	}

}

func (g *gen) setup(args []string) error {
	err := g.setCmd(args)
	if err != nil {
		return err
	}

	err = g.setTarget(args)
	if err != nil {
		return err
	}

	err = g.setUsing(args)
	if err != nil {
		return err
	}

	return nil
}

func (g *gen) setCmd(args []string) error {
	if !g.isValidCmd(args) {
		return errors.New("not a valid command")
	}

	g.Cmd = args[0]
	return nil
}

func (g *gen) isValidCmd(args []string) (valid bool) {
	if len(args) < 1 {
		return false
	}

	for _, v := range commands {
		if v == args[0] {
			return true
		}
	}

	return false
}

func (g *gen) setTarget(args []string) error {
	if (g.cmdIs(generateCmd) || g.cmdIs(helpCmd)) && !g.isValidTarget(args) {
		return errors.New("no valid target specified")
	}

	g.Target = args[1]
	return nil
}

func (g *gen) isValidTarget(args []string) (valid bool) {
	if len(args) < 2 {
		return false
	}

	for _, v := range targets {
		if v == args[1] {
			return true
		}
	}

	return false
}

func (g *gen) setUsing(args []string) error {
	if len(args) < 3 || args[2] == "" {
		return errors.New("no input file provided")
	}

	g.Using = args[2]
	return nil
}

func (g *gen) cmdIs(cmd string) bool {
	return g.Cmd == cmd
}

func (g *gen) targetIs(target string) bool {
	return g.Target == target || g.Target == allTgt
}

func (g *gen) genMeta() error {
	err := g.readFile()
	if err != nil {
		return err
	}

	err = g.parseData()
	if err != nil {
		return err
	}

	err = g.procMetadata()
	if err != nil {
		return err
	}

	return nil
}

func (g *gen) readFile() error {
	log.Printf("Reading input file: '%s'\n", g.Using)

	data, err := ioutil.ReadFile(g.Using)
	if err != nil {
		return fmt.Errorf("cannot read input file: %s", g.Using)
	}

	g.data = data

	return nil
}

func (g *gen) parseData() error {
	log.Println("Generating metadata")

	md := metadata{}
	err := yaml.Unmarshal(g.data, &md)
	if err != nil {
		return err
	}

	//log.Println(spew.Sdump(md))

	g.Meta = &md
	return nil
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}

func errMsg(err error) string {
	return strings.Title(strings.ToLower(err.Error()))
}
