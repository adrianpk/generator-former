package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type (
	gen struct {
		cmd    string
		target string

		using string
		force bool

		data     []byte
		metadata *metadata
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
	tartgets = []string{handlerTgt, migrationTgt, modelTgt, repoTgt, restcltTgt, testTgt, allTgt}
)

func main() {
	var g gen

	args := os.Args[1:]

	flag.StringVar(&g.using, "using", "", "Input file.")
	flag.BoolVar(&g.force, "force", false, "Overwrite output files.")
	flag.Parse()

	err := g.setCmdAndTarget(args)
	if err != nil {
		log.Println(err.Error())
	}

	err = g.genMeta()
	if err != nil {
		log.Println(err.Error())
	}

	if g.targetIs(handlerTgt) {
		panic("Not implemented yet")
	}

	if g.targetIs(migrationTgt) {
		panic("Not implemented yet.")
	}

	if g.targetIs(modelTgt) {
		panic("Not implemented yet.")
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

func (g *gen) setCmdAndTarget(args []string) error {
	err := g.setCmd(args)
	if err != nil {
		return err
	}

	return g.setTarget(args)
}

func (g *gen) setCmd(args []string) error {
	if !g.isValidCmd(args) {
		return errors.New("not a valid command")
	}

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

	g.target = args[1]
	return nil
}

func (g *gen) isValidTarget(args []string) (valid bool) {
	if len(args) < 2 {
		return false
	}

	for _, v := range tartgets {
		if v == args[1] {
			return true
		}
	}

	return false
}

func (g *gen) cmdIs(cmd string) bool {
	return g.cmd == cmd
}

func (g *gen) targetIs(target string) bool {
	return g.target == target || g.target == allTgt
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
	log.Printf("Reading input file: '%s'\n", g.using)

	data, err := ioutil.ReadFile(g.using)
	if err != nil {
		return fmt.Errorf("Cannot read input file: %s", g.using)
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

	log.Println(spew.Sdump(md))

	g.metadata = &md
	return nil
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}
