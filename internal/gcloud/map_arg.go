package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"sort"
	"strings"
)

type MapArg struct {
	Name string
}

func NewMapArg(name string) *MapArg {
	return &MapArg{name}
}

func (o MapArg) viperGet(v *viper.Viper) []string {

	var args []string
	for k, v := range v.GetStringMapString(o.Name) {
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	// create deterministic ordering - o
	sort.Strings(args)

	return []string{
		fmt.Sprintf("--%s", o.Name),
		strings.Join(args, ","),
	}
}
