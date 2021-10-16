package gcloud

import (
	"bytes"
	"github.com/spf13/viper"
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
		"val3.1",
		"val3.2"
	]
}
`)

func init() {
	viper.SetConfigType("json")
	_ = viper.ReadConfig(bytes.NewReader(config))
}

func TestStringArg_ViperGet(t *testing.T) {

	t.Run("Test StringArg", func(t *testing.T) {

		expected := "--arg1=val1"

		val := NewStringArg("arg1").ViperGet()

		if val[0] != expected {
			t.Fatalf("%s != %s", val, expected)
		}
	})
}

func TestListArg_ViperGet(t *testing.T) {
	t.Run("Test ListArg", func(t *testing.T) {

		expected := "--val3.1|--val3.2"

		val := NewListArg("arg3").ViperGet()
		actual := strings.Join(val, "|")

		if actual != expected {
			t.Fatalf("%s != %s", actual, expected)
		}
	})
}

func TestMapArg_ViperGet(t *testing.T) {
	t.Run("Test MapArg", func(t *testing.T) {

		expected := "--arg2|arg2.1=val2.1,arg2.2=val2.2"
		val := NewMapArg("arg2").ViperGet()
		actual := strings.Join(val, "|")

		if actual != expected {
			t.Fatalf("%s != %s", actual, expected)
		}
	})
}

func TestMapListArg_ViperGet(t *testing.T) {
	t.Run("Test MapList", func(t *testing.T) {

		expected := "--arg3=[val3.1,val3.2]"
		val := NewMapListArg("arg3").ViperGet()

		if val[0] != expected {
			t.Fatalf("%s != %s", val, expected)
		}
	})
}
