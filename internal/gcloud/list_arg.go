package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"sort"
)

type ListArg struct {
	Name string
}

func NewListArg(name string) *ListArg {
	return &ListArg{name}
}

func (o ListArg) viperGet(v *viper.Viper) []string {
	var args []string
	for _, arg := range v.GetStringSlice(o.Name) {
		args = append(args, fmt.Sprintf("--%s", arg))
	}
	sort.Strings(args)
	return args
}
