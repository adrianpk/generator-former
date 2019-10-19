package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	gen struct {
		command string

		handler   bool
		migration bool
		model     bool
		repo      bool
		restcl    bool
		test      bool
		all       bool

		input string
		force bool

		data     []byte
		metadata *metadata
	}
)

var (
	commands = []string{"generate", "help"}
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Run 'mw help' to see a list of valid commands.")
	}

	var g gen
	g.command = args[0]

	flag.BoolVar(&g.handler, "handler", false, "Generate handler and associated files.")
	flag.BoolVar(&g.migration, "migration", false, "Generate migration file.")
	flag.BoolVar(&g.model, "model", false, "Generate model file.")
	flag.BoolVar(&g.repo, "repo", false, "Generate repo file.")
	flag.BoolVar(&g.restcl, "restcl", false, "Generate REST cURL invocation shell scripts.")
	flag.BoolVar(&g.test, "test", false, "Generate handler integration test suite.")
	flag.BoolVar(&g.all, "all", true, "Generate all resource files.")
	flag.StringVar(&g.input, "input", "", "HCL input file.")
	flag.BoolVar(&g.force, "force", false, "Overwrite output files.")
	flag.Parse()

	noCmd := true

	if !g.validCommand(g.command) {
		noCmd = noCmd && false
		panic("Run 'mw help' to see a list of valid commands")
	}

	err := g.genMetada()
	if err != nil {
		log.Println(err.Error())
	}

	if g.handler || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet")
	}

	if g.migration || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet.")
	}

	if g.model || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet.")
	}

	if g.repo || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet.")
	}

	if g.restcl || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet.")
	}

	if g.test || g.all {
		noCmd = noCmd && false
		panic("Not implemented yet.")
	}

	if noCmd {
		// Show help
		panic("Not implemented yet.")
	}

}

func (g *gen) validCommand(cmd string) (ok bool) {
	for _, v := range commands {
		if v == cmd {
			return true
		}
	}
	return false
}

func (g *gen) genMetada() error {
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
	log.Printf("Reading input file: '%s'\n", g.input)

	data, err := ioutil.ReadFile(g.input)
	if err != nil {
		return fmt.Errorf("Cannot read input file: %s", g.input)
	}

	g.data = data
	return nil
}

func (g *gen) parseData() error {
	md := metadata{}
	err := yaml.Unmarshal(g.data, &md)
	if err != nil {
		return err
	}

	g.metadata = &md
	return nil
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}
