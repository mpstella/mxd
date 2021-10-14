package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

type GcloudCommand struct {
	App       []string
	Arguments []*GcloudArgument
	mapping   map[string]GcloudArgument
}

type GcloudArgument interface {
	ViperGet() []string
}

type GcloudMapOption struct {
	Name string
}

type GcloudStringOption struct {
	Name string
}

type GcloudListOption struct {
	Name string
}

type GcloudMapListOption struct {
	Name string
}

func NewMapOption(name string) *GcloudMapOption {
	return &GcloudMapOption{name}
}

func NewStringOption(name string) *GcloudStringOption {
	return &GcloudStringOption{name}
}

func NewListOption(name string) *GcloudListOption {
	return &GcloudListOption{name}
}

func NewMapListOption(name string) *GcloudMapListOption {
	return &GcloudMapListOption{name}
}

func NewGcloudCommand(app ...string) *GcloudCommand {
	return &GcloudCommand{app, make([]*GcloudArgument, 0, 10), make(map[string]GcloudArgument)}
}

func (g *GcloudCommand) AddStringMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewStringOption(arg)
	}
}

func (g *GcloudCommand) AddListMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewListOption(arg)
	}
}

func (g *GcloudCommand) AddMapListMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewMapListOption(arg)
	}
}

func (g *GcloudCommand) AddMapMapping(args ...string) {
	for _, arg := range args {
		g.mapping[arg] = NewMapOption(arg)
	}
}

func (g *GcloudCommand) AddArg(arg *GcloudArgument) {
	g.Arguments = append(g.Arguments, arg)
}

func (g *GcloudCommand) ViperBuild() {

	for k, _ := range viper.AllSettings() {

		if _, ok := g.mapping[k]; !ok {
			g.AddStringMapping(k)
		}
		gco := g.mapping[k]
		g.AddArg(&gco)
	}
}

func (o GcloudListOption) ViperGet() []string {
	var opts []string
	for _, opt := range viper.GetStringSlice(o.Name) {
		opts = append(opts, fmt.Sprintf("--%s", opt))
	}
	return opts
}

func (o GcloudStringOption) ViperGet() []string {
	return []string{fmt.Sprintf("--%s=%s", o.Name, viper.GetString(o.Name))}
}

func (o GcloudMapOption) ViperGet() []string {
	var l []string
	for k, v := range viper.GetStringMapString(o.Name) {
		l = append(l, fmt.Sprintf("%s=%s", k, v))
	}
	return []string{
		fmt.Sprintf("--%s", o.Name),
		strings.Join(l, ","),
	}
}

func (o GcloudMapListOption) ViperGet() []string {
	return []string{
		fmt.Sprintf("--%s=[%s]", o.Name, strings.Join(viper.GetStringSlice(o.Name), ",")),
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

	cmd = append(cmd, g.App...)
	cmd = append(cmd, args...)

	for _, arg := range g.Arguments {
		cmd = append(cmd, (*arg).ViperGet()...)
	}

	if Verbose {
		fmt.Printf("%s\n", cmd)
	}

	shellcmd := exec.Command("gcloud", cmd...)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	shellcmd.Stdout = &stdOut
	shellcmd.Stderr = &stdErr

	if err := shellcmd.Run(); err != nil {
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
