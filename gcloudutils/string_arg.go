package gcloudutils

import (
	"fmt"
	"github.com/spf13/viper"
)

type GcloudStringArg struct {
	Name string
}

func NewStringArg(name string) *GcloudStringArg {
	return &GcloudStringArg{name}
}

func (o GcloudStringArg) ViperGet() []string {
	return []string{fmt.Sprintf("--%s=%s", o.Name, viper.GetString(o.Name))}
}
