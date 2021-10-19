package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"sort"
	"strings"
)

type MapListArg struct {
	Name string
}

func NewMapListArg(name string) *MapListArg {
	return &MapListArg{name}
}

func (o MapListArg) viperGet(v *viper.Viper) []string {
	args := v.GetStringSlice(o.Name)
	sort.Strings(args)
	return []string{
		fmt.Sprintf("--%s=[%s]", o.Name, strings.Join(args, ",")),
	}
}
