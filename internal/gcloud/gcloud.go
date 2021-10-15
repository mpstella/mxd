package gcloud

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
	"path/filepath"
	"strings"
)

type Argument interface {
	ViperGet() []string
}

var Verbose = false

type Command struct {
	app       []string
	arguments []*Argument
	mapping   map[string]Argument
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
	}
}

func (g *Command) ReadConfig(configPath string) error {

	abs, err := filepath.Abs(configPath)

	if err != nil {
		panic(err)
	}

	base := filepath.Base(abs)
	path := filepath.Dir(abs)

	viper.SetConfigName(strings.Split(base, ".")[0])
	viper.AddConfigPath(path)

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	g.viperBuild()
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

func (g *Command) viperBuild() {

	for k := range viper.AllSettings() {

		if _, ok := g.mapping[k]; !ok {
			g.AddStringMapping(k)
		}
		gco, _ := g.mapping[k]
		g.AddArgument(&gco)
	}
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
		fmt.Printf("  %s\n", strings.Join((*arg).ViperGet(), " "))
	}
	fmt.Println("}")
	fmt.Println("===================================")
}

func (g *Command) Run(args ...string) error {

	cmd := make([]string, 0, 10)

	if component != "" {
		cmd = append(cmd, component)
	}

	cmd = append(cmd, g.app...)
	cmd = append(cmd, args...)

	for _, arg := range g.arguments {
		cmd = append(cmd, (*arg).ViperGet()...)
	}

	if Verbose {
		fmt.Println(cmd)
	}

	shellCmd := exec.Command("gcloud", cmd...)
	var stdOut, stdErr bytes.Buffer
	shellCmd.Stdout = &stdOut
	shellCmd.Stderr = &stdErr

	if err := shellCmd.Run(); err != nil {
		fmt.Printf("An error has occurred with %s\n", err)

		for _, s := range strings.Split(stdErr.String(), "\n") {
			fmt.Println(s)
		}
		return err
	}

	for _, s := range strings.Split(stdOut.String(), "\n") {
		fmt.Println(s)
	}

	return nil
}
