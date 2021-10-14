package gcloudutils

import (
	"fmt"
	"github.com/spf13/viper"
)

type GcloudListArg struct {
	Name string
}

func NewListArg(name string) *GcloudListArg {
	return &GcloudListArg{name}
}

func (o GcloudListArg) ViperGet() []string {
	var opts []string
	for _, opt := range viper.GetStringSlice(o.Name) {
		opts = append(opts, fmt.Sprintf("--%s", opt))
	}
	return opts
}
