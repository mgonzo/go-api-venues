package route

import (
	"fmt"
	"github.com/mgonzo/venues/handler"
)

var model = "venues"

func base(model string) string {
	return fmt.Sprintf("/%[1]s", model)
}

func qualifier(model string) string {
	return fmt.Sprintf("/%[1]s/{qualifier}", model)
}

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
