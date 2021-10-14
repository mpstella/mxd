package gcloudutils

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type GcloudMapArg struct {
	Name string
}

func NewMapArg(name string) *GcloudMapArg {
	return &GcloudMapArg{name}
}

func (o GcloudMapArg) ViperGet() []string {
	var l []string
	for k, v := range viper.GetStringMapString(o.Name) {
		l = append(l, fmt.Sprintf("%s=%s", k, v))
	}
	return []string{
		fmt.Sprintf("--%s", o.Name),
		strings.Join(l, ","),
	}
}
