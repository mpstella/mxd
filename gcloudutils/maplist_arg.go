package gcloudutils

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type GcloudMapListArg struct {
	Name string
}

func NewMapListArg(name string) *GcloudMapListArg {
	return &GcloudMapListArg{name}
}

func (o GcloudMapListArg) ViperGet() []string {
	return []string{
		fmt.Sprintf("--%s=[%s]", o.Name, strings.Join(viper.GetStringSlice(o.Name), ",")),
	}
}
