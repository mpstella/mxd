package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"mxd/internal"
	"path/filepath"
	"strings"
)

type Argument interface {
	viperGet(v *viper.Viper) []string
}

var Verbose = false

const gcloudRunCommand = "gcloud"

type Command struct {
	app       []string
	arguments []*Argument
	mapping   map[string]Argument
	viper     *viper.Viper
}

var component = ""

func SetAlpha() {
	component = "alpha"
}

func SetBeta() {
	component = "beta"
}

func NewCommand(app ...string) *Command {
	return &Command{app,
		make([]*Argument, 0, 10),
		make(map[string]Argument),
		viper.New(),
	}
}

func (g *Command) getViper(configPath string) error {

	abs, err := filepath.Abs(configPath)

	if err != nil {
		return err
	}

	base := filepath.Base(abs)
	path := filepath.Dir(abs)

	g.viper.SetConfigName(strings.Split(base, ".")[0])
	g.viper.AddConfigPath(path)

	if err = g.viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (g *Command) ReadConfig(configPath string) error {

	err := g.getViper(configPath)

	if err != nil {
		return err
	}

	for k := range g.viper.AllSettings() {

		if _, ok := g.mapping[k]; !ok {
			g.AddStringMapping(k)
		}
		gco, _ := g.mapping[k]
		g.AddArgument(&gco)
	}
	return nil
}

func (g *Command) AddStringMapping(args ...string) *Command {
	for _, arg := range args {
		g.mapping[arg] = NewStringArg(arg)
	}
	return g
}

func (g *Command) AddListMapping(args ...string) *Command {
	for _, arg := range args {
		g.mapping[arg] = NewListArg(arg)
	}
	return g
}

func (g *Command) AddMapListMapping(args ...string) *Command {
	for _, arg := range args {
		g.mapping[arg] = NewMapListArg(arg)
	}
	return g
}

func (g *Command) AddMapMapping(args ...string) *Command {
	for _, arg := range args {
		g.mapping[arg] = NewMapArg(arg)
	}
	return g
}

func (g *Command) AddArgument(arg *Argument) *Command {
	g.arguments = append(g.arguments, arg)
	return g
}

func (g *Command) Debug() {

	fmt.Println("========== GcloudCommand ==========")
	fmt.Println("MAP: {")
	for k, v := range g.mapping {
		fmt.Printf("  %s : %#v\n", k, v)
	}
	fmt.Println("}")
	fmt.Printf("CMD: %s {\n", g.app)
	for _, arg := range g.arguments {
		fmt.Printf("  %s\n", strings.Join((*arg).viperGet(g.viper), " "))
	}
	fmt.Println("}")
	fmt.Println("===================================")
}

func (g *Command) Run(args ...string) error {

	runCommand := make([]string, 0, 10)

	if component != "" {
		runCommand = append(runCommand, component)
	}

	runCommand = append(runCommand, g.app...)
	runCommand = append(runCommand, args...)

	for _, arg := range g.arguments {
		runCommand = append(runCommand, (*arg).viperGet(g.viper)...)
	}

	if Verbose {
		fmt.Println(runCommand)
	}

	_, _, err := internal.RunShellCommand(gcloudRunCommand, runCommand...)

	return err
}
