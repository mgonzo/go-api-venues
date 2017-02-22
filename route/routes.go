package route

import (
	"bytes"
	"fmt"
	"github.com/mgonzo/venues/handler"
	"github.com/spf13/viper"
)

func config() string {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return viper.GetString("model")
}

func base(model string) string {
	var buffer bytes.Buffer
	buffer.WriteString("/")
	buffer.WriteString(model)
	return buffer.String()
}

func qualifier(model string) string {
	var pattern = base(model)
	var buffer bytes.Buffer
	buffer.WriteString(pattern)
	buffer.WriteString("/{qualifier}")
	return buffer.String()
}

var model = config()

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Error,
	},
	Route{
		"Collection",
		"GET",
		base(model),
		handler.Collection,
	},
	Route{
		"Show",
		"GET",
		qualifier(model),
		handler.Show,
	},
	Route{
		"Create",
		"POST",
		base(model),
		handler.Create,
	},
	Route{
		"Replace",
		"PUT",
		qualifier(model),
		handler.Replace,
	},
	Route{
		"Update",
		"PATCH",
		qualifier(model),
		handler.Update,
	},
	Route{
		"Remove",
		"DELETE",
		qualifier(model),
		handler.Remove,
	},
}
