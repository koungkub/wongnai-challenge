package main

import (
	"strings"

	"github.com/koungkub/wongnai/internal/connection"
	"github.com/koungkub/wongnai/internal/model"
	"github.com/koungkub/wongnai/internal/route"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {

	svcName := viper.GetString("WONGNAI.NAME")
	log := connection.NewLog(svcName)

	host, pass := viper.GetString("APP.CACHE.HOST"), viper.GetString("APP.CACHE.PASS")
	cache := connection.NewCache(host, pass)

	driver, dsn := viper.GetString("APP.DB.DRIVER"), viper.GetString("APP.DB.DSN")
	db, err := connection.NewDB(driver, dsn)
	if err != nil {
		panic(err)
	}

	conf := model.NewConf(db, log, cache)

	r := route.New(conf)

	r.Listen(3000)
}
