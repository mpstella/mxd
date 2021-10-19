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

func (o StringArg) viperGet(v *viper.Viper) []string {
	return []string{fmt.Sprintf("--%s=%s", o.Name, v.GetString(o.Name))}
}
