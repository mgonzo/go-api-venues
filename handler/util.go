package handler

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/spf13/viper"
)

func config() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func connect() string {
	config()
	connect := fmt.Sprintf(
		"%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s",
		viper.GetString("sqluser"),
		viper.GetString("sqlpass"),
		viper.GetString("sqlhost"),
		viper.GetString("sqlport"),
		viper.GetString("sqlname"),
	)

	return connect
}
