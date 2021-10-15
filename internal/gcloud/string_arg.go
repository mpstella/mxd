package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
)

type StringArg struct {
	Name string
}

func NewStringArg(name string) *StringArg {
	return &StringArg{name}
}

func (o StringArg) ViperGet() []string {
	return []string{fmt.Sprintf("--%s=%s", o.Name, viper.GetString(o.Name))}
}
