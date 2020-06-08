package main

import (
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {

}
