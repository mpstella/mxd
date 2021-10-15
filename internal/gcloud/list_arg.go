package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
)

type ListArg struct {
	Name string
}

func NewListArg(name string) *ListArg {
	return &ListArg{name}
}

func (o ListArg) ViperGet() []string {
	var opts []string
	for _, opt := range viper.GetStringSlice(o.Name) {
		opts = append(opts, fmt.Sprintf("--%s", opt))
	}
	return opts
}
