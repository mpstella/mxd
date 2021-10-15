package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type MapArg struct {
	Name string
}

func NewMapArg(name string) *MapArg {
	return &MapArg{name}
}

func (o MapArg) ViperGet() []string {
	var l []string
	for k, v := range viper.GetStringMapString(o.Name) {
		l = append(l, fmt.Sprintf("%s=%s", k, v))
	}
	return []string{
		fmt.Sprintf("--%s", o.Name),
		strings.Join(l, ","),
	}
}
