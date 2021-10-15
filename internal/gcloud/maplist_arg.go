package gcloud

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type MapListArg struct {
	Name string
}

func NewMapListArg(name string) *MapListArg {
	return &MapListArg{name}
}

func (o MapListArg) ViperGet() []string {
	return []string{
		fmt.Sprintf("--%s=[%s]", o.Name, strings.Join(viper.GetStringSlice(o.Name), ",")),
	}
}
