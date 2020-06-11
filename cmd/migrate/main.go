package main

import (
	"flag"
	"strings"

	"github.com/koungkub/wongnai/internal/connection"
	"github.com/koungkub/wongnai/internal/worker"
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
	food := flag.String("food", "datasets/food_dictionary.txt", "relative path to food_dictionary.txt")
	review := flag.String("review", "datasets/test_file.csv", "relative path to test_file.csv")
	flag.Parse()

	svcName := viper.GetString("MIGRATE.NAME")
	log := connection.NewLog(svcName)

	driver, dsn := viper.GetString("APP.DB.DRIVER"), viper.GetString("APP.DB.DSN")
	db, err := connection.NewDB(driver, dsn)
	if err != nil {
		panic(err)
	}

	f, err := worker.OpenFile(*food)
	if err != nil {
		log.Fatal(err)
	}

	r, err := worker.OpenFile(*review)
	if err != nil {
		log.Fatal(err)
	}

	if err := worker.MigrateFoodDic(db, f); err != nil {
		log.Fatal(err)
	}

	if err := worker.MigrateReview(db, r); err != nil {
		log.Fatal(err)
	}
}
