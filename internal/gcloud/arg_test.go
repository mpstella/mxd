package gcloud

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var config = []byte(`
{
	"arg1": "val1",
	"arg2": {
		"arg2.1": "val2.1",
		"arg2.2": "val2.2"
	},
	"arg3": [
		"val3.2",
		"val3.1"
	]
}
`)

var v *viper.Viper

func init() {
	v = viper.New()
	v.SetConfigType("json")
	_ = v.ReadConfig(bytes.NewReader(config))
}

func TestStringArg_ViperGet(t *testing.T) {

	t.Run("Test StringArg", func(t *testing.T) {

		expected := "--arg1=val1"
		actual := NewStringArg("arg1").viperGet(v)
		assert.Equal(t, expected, actual[0])
	})
}

func TestListArg_ViperGet(t *testing.T) {
	t.Run("Test ListArg", func(t *testing.T) {

		expected := "--val3.1|--val3.2"
		val := NewListArg("arg3").viperGet(v)
		actual := strings.Join(val, "|")
		assert.Equal(t, expected, actual)
	})
}

func TestMapArg_ViperGet(t *testing.T) {
	t.Run("Test MapArg", func(t *testing.T) {

		expected := "--arg2|arg2.1=val2.1,arg2.2=val2.2"
		val := NewMapArg("arg2").viperGet(v)
		actual := strings.Join(val, "|")
		assert.Equal(t, expected, actual)
	})
}

func TestMapListArg_ViperGet(t *testing.T) {
	t.Run("Test MapList", func(t *testing.T) {

		expected := "--arg3=[val3.1,val3.2]"
		actual := NewMapListArg("arg3").viperGet(v)
		assert.Equal(t, expected, actual[0])
	})
}
