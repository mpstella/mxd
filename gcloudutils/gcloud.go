package gcloudutils

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

type GcloudArgument interface {
	ViperGet() []string
}

type GcloudCommand struct {
	App       []string
	Arguments []*GcloudArgument
	mapping   map[string]GcloudArgument
	Verbose   bool
	Component string
}

func NewGcloudCommand(app ...string) *GcloudCommand {
	return &GcloudCommand{app,
		make([]*GcloudArgument, 0, 10),
		make(map[string]GcloudArgument),
		false,
		"",
	}
}

func (g *GcloudCommand) UseAlpha() {
	g.Component = "alpha"
}

func (g *GcloudCommand) UseBeta() {
	g.Component = "beta"
}

func (g *GcloudCommand) AddStringMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewStringArg(arg)
	}
}

func (g *GcloudCommand) AddListMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewListArg(arg)
	}
}

func (g *GcloudCommand) AddMapListMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewMapListArg(arg)
	}
}

func (g *GcloudCommand) AddMapMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewMapArg(arg)
	}
}

func (g *GcloudCommand) AddArg(arg *GcloudArgument) {
	g.Arguments = append(g.Arguments, arg)
}

func (g *GcloudCommand) ViperBuild() {

	for k := range viper.AllSettings() {

		if _, ok := g.mapping[k]; !ok {
			g.AddStringMapping(k)
		}
		gco, _ := g.mapping[k]
		g.AddArg(&gco)
	}
}

func (g *GcloudCommand) Debug() {

	fmt.Println("========== GcloudCommand ==========")
	fmt.Println("MAP: {")
	for k, v := range g.mapping {
		fmt.Printf("  %s : %#v\n", k, v)
	}
	fmt.Println("}")
	fmt.Printf("CMD: %s {\n", g.App)
	for _, arg := range g.Arguments {
		fmt.Printf("  %s\n", strings.Join((*arg).ViperGet(), " "))
	}
	fmt.Println("}")
	fmt.Println("===================================")
}

func (g *GcloudCommand) Run(args ...string) error {

	cmd := make([]string, 0, 10)

	if g.Component != "" {
		cmd = append(cmd, g.Component)
	}

	cmd = append(cmd, g.App...)
	cmd = append(cmd, args...)

	for _, arg := range g.Arguments {
		cmd = append(cmd, (*arg).ViperGet()...)
	}

	if g.Verbose {
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
